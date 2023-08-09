package keeper

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"

	"github.com/osmosis-labs/osmosis/osmoutils"
	lockuptypes "github.com/osmosis-labs/osmosis/v17/x/lockup/types"
	"github.com/osmosis-labs/osmosis/v17/x/superfluid/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// GetTotalSyntheticAssetsLocked returns the total amount of the given denom locked.
func (k Keeper) GetTotalSyntheticAssetsLocked(ctx sdk.Context, denom string) sdk.Int {
	return k.lk.GetPeriodLocksAccumulation(ctx, lockuptypes.QueryCondition{
		LockQueryType: lockuptypes.ByDuration,
		Denom:         denom,
		Duration:      k.sk.UnbondingTime(ctx),
	})
}

// GetExpectedDelegationAmount returns the total number of fury the intermediary account
// has delegated using the most recent fury equivalent multiplier.
// This is labeled as expected because the way it calculates the amount can
// lead rounding errors from the true delegated amount.
func (k Keeper) GetExpectedDelegationAmount(ctx sdk.Context, acc types.SuperfluidIntermediaryAccount) (sdk.Int, error) {
	// (1) Find how many tokens total T are locked for (denom, validator) pair
	totalSuperfluidDelegation := k.GetTotalSyntheticAssetsLocked(ctx, stakingSyntheticDenom(acc.Denom, acc.ValAddr))
	// (2) Multiply the T tokens, by the number of superfluid fury per token, to get the total amount
	// of fury we expect.
	refreshedAmount, err := k.GetSuperfluidFURYTokens(ctx, acc.Denom, totalSuperfluidDelegation)
	if err != nil {
		return sdk.Int{}, err
	}
	return refreshedAmount, nil
}

// RefreshIntermediaryDelegationAmounts refreshes the amount of delegation for all intermediary accounts.
// This method includes minting new fury if the refreshed delegation amount has increased, and
// instantly undelegating and burning if the refreshed delgation has decreased.
func (k Keeper) RefreshIntermediaryDelegationAmounts(ctx sdk.Context) {
	// iterate over all intermedairy accounts - every (denom, validator) pair
	accs := k.GetAllIntermediaryAccounts(ctx)
	for _, acc := range accs {
		mAddr := acc.GetAccAddress()

		valAddress, err := sdk.ValAddressFromBech32(acc.ValAddr)
		if err != nil {
			panic(err)
		}

		validator, found := k.sk.GetValidator(ctx, valAddress)
		if !found {
			k.Logger(ctx).Error(fmt.Sprintf("validator not found or %s", acc.ValAddr))
			continue
		}

		currentAmount := sdk.NewInt(0)
		delegation, found := k.sk.GetDelegation(ctx, mAddr, valAddress)
		if !found {
			// continue if current delegation is 0, in case its really a dust delegation
			// that becomes worth something after refresh.
			k.Logger(ctx).Info(fmt.Sprintf("Existing delegation not found for %s with %s during superfluid refresh."+
				" It may have been previously bonded, but now unbonded.", mAddr.String(), acc.ValAddr))
		} else {
			currentAmount = validator.TokensFromShares(delegation.Shares).RoundInt()
		}

		refreshedAmount, err := k.GetExpectedDelegationAmount(ctx, acc)
		if err != nil {
			ctx.Logger().Error("Error in GetExpectedDelegationAmount (likely that underlying LP share is no longer superfluid capable), state update reverted", err)
		}

		if refreshedAmount.GT(currentAmount) {
			adjustment := refreshedAmount.Sub(currentAmount)
			err = k.mintOsmoTokensAndDelegate(ctx, adjustment, acc)
			if err != nil {
				ctx.Logger().Error("Error in forceUndelegateAndBurnOsmoTokens, state update reverted", err)
			}
		} else if currentAmount.GT(refreshedAmount) {
			// In this case, we want to change the IA's delegated balance to be refreshed Amount
			// which is less than what it already has.
			// This means we need to "InstantUndelegate" some of its delegation (not going through the unbonding queue)
			// and then burn that excessly delegated bits.
			adjustment := currentAmount.Sub(refreshedAmount)

			err := k.forceUndelegateAndBurnOsmoTokens(ctx, adjustment, acc)
			if err != nil {
				ctx.Logger().Error("Error in forceUndelegateAndBurnOsmoTokens, state update reverted", err)
			}
		} else {
			ctx.Logger().Info("Intermediary account already has correct delegation amount?" +
				" This with high probability implies the exact same spot price as the last epoch," +
				"and no delegation changes.")
		}
	}
}

// IncreaseSuperfluidDelegation increases the amount of existing superfluid delegation.
// This method would return an error if the lock has not been superfluid delegated before.
func (k Keeper) IncreaseSuperfluidDelegation(ctx sdk.Context, lockID uint64, amount sdk.Coins) error {
	acc, found := k.GetIntermediaryAccountFromLockId(ctx, lockID)
	if !found {
		return nil
	}

	// mint FURY token based on the most recent fury equivalent multiplier
	// of locked denom to denom module account
	furyAmt, err := k.GetSuperfluidFURYTokens(ctx, acc.Denom, amount.AmountOf(acc.Denom))
	if err != nil {
		return err
	}
	if furyAmt.IsZero() {
		return nil
	}

	err = k.mintOsmoTokensAndDelegate(ctx, furyAmt, acc)
	if err != nil {
		return err
	}

	return nil
}

// basic validation for locks to be eligible for superfluid delegation. This includes checking
// - that the sender is the owner of the lock
// - that the lock is consisted of single coin
func (k Keeper) validateLockForSF(ctx sdk.Context, lock *lockuptypes.PeriodLock, sender string) error {
	if lock.Owner != sender {
		return lockuptypes.ErrNotLockOwner
	}
	if lock.Coins.Len() != 1 {
		return types.ErrMultipleCoinsLockupNotSupported
	}
	return nil
}

// validateLockForSFDelegate runs the following sanity checks on the lock:
// - the sender is the owner of the lock
// - the lock is consisted of a single coin
// - the asset is registered as a superfluid asset via governance
// - the lock is not unlocking
// - lock duration is greater or equal to the unbonding time
// - lock should not be already superfluid staked
func (k Keeper) validateLockForSFDelegate(ctx sdk.Context, lock *lockuptypes.PeriodLock, sender string) error {
	err := k.validateLockForSF(ctx, lock, sender)
	if err != nil {
		return err
	}

	denom := lock.Coins[0].Denom

	// ensure that the locks underlying denom is for an existing superfluid asset
	_, err = k.GetSuperfluidAsset(ctx, denom)
	if err != nil {
		return err
	}

	// prevent unbonding lockups to be not able to be used for superfluid staking
	if lock.IsUnlocking() {
		return errorsmod.Wrapf(types.ErrUnbondingLockupNotSupported, "lock id : %d", lock.ID)
	}

	// ensure that lock duration >= staking.UnbondingTime
	unbondingTime := k.sk.GetParams(ctx).UnbondingTime
	if lock.Duration < unbondingTime {
		return errorsmod.Wrapf(types.ErrNotEnoughLockupDuration, "lock duration (%d) must be greater than unbonding time (%d)", lock.Duration, unbondingTime)
	}

	// Thus when we stake now, this will be the only superfluid position for this lockID.
	if k.alreadySuperfluidStaking(ctx, lock.ID) {
		return errorsmod.Wrapf(types.ErrAlreadyUsedSuperfluidLockup, "lock id : %d", lock.ID)
	}

	return nil
}

// ensure the valAddr is correctly formatted & corresponds to a real validator on chain.
func (k Keeper) validateValAddrForDelegate(ctx sdk.Context, valAddr string) (stakingtypes.Validator, error) {
	valAddress, err := sdk.ValAddressFromBech32(valAddr)
	if err != nil {
		return stakingtypes.Validator{}, err
	}
	validator, found := k.sk.GetValidator(ctx, valAddress)
	if !found {
		return stakingtypes.Validator{}, stakingtypes.ErrNoValidatorFound
	}
	return validator, nil
}

// SuperfluidDelegate superfluid delegates fury equivalent amount the given lock holds.
// The actual delegation is done by using/creating an intermediary account for the (denom, validator) pair
// and having the intermediary account delegate to the designated validator, not by the sender themselves.
// A state entry of IntermediaryAccountConnection is stored to store the connection between the lock ID
// and the intermediary account, as an intermediary account does not serve for delegations from a single delegator.
// The actual amount of delegation is not equal to the equivalent amount of fury the lock has. That is,
// the actual amount of delegation is amount * fury equivalent multiplier * (1 - k.RiskFactor(asset)).
func (k Keeper) SuperfluidDelegate(ctx sdk.Context, sender string, lockID uint64, valAddr string) error {
	lock, err := k.lk.GetLockByID(ctx, lockID)
	if err != nil {
		return err
	}

	// This guarantees the lockID does not already have a superfluid stake position
	// associated with it, the lock is sufficiently long, the lock only locks one asset, etc.
	// Thus when we stake this lock, it will be the only superfluid position for this lockID.
	err = k.validateLockForSFDelegate(ctx, lock, sender)
	if err != nil {
		return err
	}
	lockedCoin := lock.Coins[0]

	// get the intermediate account for this (denom, validator) pair.
	// This account tracks the amount of fury being considered as staked.
	// If an intermediary account doesn't exist, then create it + a perpetual gauge.
	acc, err := k.GetOrCreateIntermediaryAccount(ctx, lockedCoin.Denom, valAddr)
	if err != nil {
		return err
	}
	// create connection record between lock id and intermediary account
	k.SetLockIdIntermediaryAccountConnection(ctx, lockID, acc)

	// Register a synthetic lockup for superfluid staking
	err = k.createSyntheticLockup(ctx, lockID, acc, bondedStatus)
	if err != nil {
		return err
	}

	// Find how many new fury tokens this delegation is worth at superfluids current risk adjustment
	// and twap of the denom.
	amount, err := k.GetSuperfluidFURYTokens(ctx, acc.Denom, lockedCoin.Amount)
	if err != nil {
		return err
	}
	if amount.IsZero() {
		return types.ErrOsmoEquivalentZeroNotAllowed
	}

	return k.mintOsmoTokensAndDelegate(ctx, amount, acc)
}

// undelegateCommon is a helper function for SuperfluidUndelegate and superfluidUndelegateToConcentratedPosition.
// It performs the following tasks:
// - checks that the lock is valid for superfluid staking
// - gets the intermediary account associated with the lock id
// - deletes the connection between the lock id and the intermediary account
// - deletes the synthetic lockup associated with the lock id
// - undelegates the superfluid staking position associated with the lock id and burns the underlying fury tokens
// - returns the intermediary account
func (k Keeper) undelegateCommon(ctx sdk.Context, sender string, lockID uint64) (types.SuperfluidIntermediaryAccount, error) {
	lock, err := k.lk.GetLockByID(ctx, lockID)
	if err != nil {
		return types.SuperfluidIntermediaryAccount{}, err
	}
	err = k.validateLockForSF(ctx, lock, sender)
	if err != nil {
		return types.SuperfluidIntermediaryAccount{}, err
	}
	lockedCoin := lock.Coins[0]

	// get the intermediate account associated with lock id, and delete the connection.
	intermediaryAcc, found := k.GetIntermediaryAccountFromLockId(ctx, lockID)
	if !found {
		return types.SuperfluidIntermediaryAccount{}, types.ErrNotSuperfluidUsedLockup
	}
	k.DeleteLockIdIntermediaryAccountConnection(ctx, lockID)

	// Delete the old synthetic lockup
	synthdenom := stakingSyntheticDenom(lockedCoin.Denom, intermediaryAcc.ValAddr)
	err = k.lk.DeleteSyntheticLockup(ctx, lockID, synthdenom)
	if err != nil {
		return types.SuperfluidIntermediaryAccount{}, err
	}

	// undelegate this lock's delegation amount, and burn the minted fury.
	amount, err := k.GetSuperfluidFURYTokens(ctx, intermediaryAcc.Denom, lockedCoin.Amount)
	if err != nil {
		return types.SuperfluidIntermediaryAccount{}, err
	}
	err = k.forceUndelegateAndBurnOsmoTokens(ctx, amount, intermediaryAcc)
	if err != nil {
		return types.SuperfluidIntermediaryAccount{}, err
	}
	return intermediaryAcc, nil
}

// SuperfluidUndelegate starts undelegating superfluid delegated position for the given lock.
// Undelegation is done instantly and the equivalent amount is sent to the module account
// where it is burnt. Note that this method does not include unbonding the lock
// itself.
func (k Keeper) SuperfluidUndelegate(ctx sdk.Context, sender string, lockID uint64) error {
	intermediaryAcc, err := k.undelegateCommon(ctx, sender, lockID)
	if err != nil {
		return err
	}
	// Create a new synthetic lockup representing the unstaking side.
	return k.createSyntheticLockup(ctx, lockID, intermediaryAcc, unlockingStatus)
}

// SuperfluidUndelegateToConcentratedPosition starts undelegating superfluid delegated position for the given lock. It behaves similarly to SuperfluidUndelegate,
// however it does not create a new synthetic lockup representing the unstaking side. This is because after the time this function is called, we might
// want to perform more operations prior to creating a lock. Once the actual lock is created, the synthetic lockup representing the unstaking side
// should eventually be created as well. Use this function with caution to avoid accidentally missing synthetic lock creation.
func (k Keeper) SuperfluidUndelegateToConcentratedPosition(ctx sdk.Context, sender string, gammLockID uint64) (types.SuperfluidIntermediaryAccount, error) {
	return k.undelegateCommon(ctx, sender, gammLockID)
}

// partialUndelegateCommon acts similarly to undelegateCommon, but undelegates a partial amount of the lock's delegation rather than the full amount. The amount
// that is undelegated is placed in a new lock. This function returns the intermediary account associated with the original lock ID as well as the new lock that was created.
// An error is returned if the amount to undelegate is greater than the locked amount.
func (k Keeper) partialUndelegateCommon(ctx sdk.Context, sender string, lockID uint64, amountToUndelegate sdk.Coin) (intermediaryAcc types.SuperfluidIntermediaryAccount, newlock *lockuptypes.PeriodLock, err error) {
	lock, err := k.lk.GetLockByID(ctx, lockID)
	if err != nil {
		return types.SuperfluidIntermediaryAccount{}, &lockuptypes.PeriodLock{}, err
	}
	err = k.validateLockForSF(ctx, lock, sender)
	if err != nil {
		return types.SuperfluidIntermediaryAccount{}, &lockuptypes.PeriodLock{}, err
	}

	if amountToUndelegate.Amount.GTE(lock.Coins[0].Amount) {
		return types.SuperfluidIntermediaryAccount{}, &lockuptypes.PeriodLock{}, fmt.Errorf("partial undelegate amount must be less than the locked amount")
	}

	// get the intermediate account associated with lock id, and delete the connection.
	intermediaryAcc, found := k.GetIntermediaryAccountFromLockId(ctx, lockID)
	if !found {
		return types.SuperfluidIntermediaryAccount{}, &lockuptypes.PeriodLock{}, types.ErrNotSuperfluidUsedLockup
	}

	// undelegate the desired lock amount, and burn the minted fury.
	amount, err := k.GetSuperfluidFURYTokens(ctx, intermediaryAcc.Denom, amountToUndelegate.Amount)
	if err != nil {
		return types.SuperfluidIntermediaryAccount{}, &lockuptypes.PeriodLock{}, err
	}
	err = k.forceUndelegateAndBurnOsmoTokens(ctx, amount, intermediaryAcc)
	if err != nil {
		return types.SuperfluidIntermediaryAccount{}, &lockuptypes.PeriodLock{}, err
	}

	// Move the funds from the old lock to a new lock with the remaining amount.
	newLock, err := k.lk.SplitLock(ctx, *lock, sdk.NewCoins(amountToUndelegate), true)
	if err != nil {
		return types.SuperfluidIntermediaryAccount{}, &lockuptypes.PeriodLock{}, err
	}

	return intermediaryAcc, &newLock, nil
}

// partialSuperfluidUndelegate starts undelegating a portion of a superfluid delegated position for the given lock.
// Undelegation is done instantly and the equivalent amount is sent to the module account
// where it is burnt. Note that this method does not include unbonding the lock
// itself.
// nolint: unused
func (k Keeper) partialSuperfluidUndelegate(ctx sdk.Context, sender string, lockID uint64, amountToUndelegate sdk.Coin) error {
	intermediaryAcc, newLock, err := k.partialUndelegateCommon(ctx, sender, lockID, amountToUndelegate)
	if err != nil {
		return err
	}
	// Create a new synthetic lockup representing the unstaking side.
	return k.createSyntheticLockup(ctx, newLock.ID, intermediaryAcc, unlockingStatus)
}

// SuperfluidUnbondLock unbonds the lock that has been used for superfluid staking.
// This method would return an error if the underlying lock is not superfluid undelegating.
func (k Keeper) SuperfluidUnbondLock(ctx sdk.Context, underlyingLockId uint64, sender string) error {
	_, err := k.unbondLock(ctx, underlyingLockId, sender, sdk.Coins{})
	return err
}

// SuperfluidUndelegateAndUnbondLock unbonds given amount from the
// underlying lock that has been used for superfluid staking.
// This method returns the lock id, same lock id if unlock amount is equal to the
// underlying lock amount. Otherwise it returns the newly created lock id.
// Note that we can either partially or fully undelegate and unbond lock using this method.
func (k Keeper) SuperfluidUndelegateAndUnbondLock(ctx sdk.Context, lockID uint64, sender string, amount sdk.Int) (uint64, error) {
	lock, err := k.lk.GetLockByID(ctx, lockID)
	if err != nil {
		return 0, err
	}

	coins := sdk.Coins{sdk.NewCoin(lock.Coins[0].Denom, amount)}
	if coins[0].IsZero() {
		return 0, fmt.Errorf("amount to unlock must be greater than 0")
	}
	if lock.Coins[0].IsLT(coins[0]) {
		return 0, fmt.Errorf("requested amount to unlock exceeds locked tokens")
	}

	// get intermediary account before connection is deleted in SuperfluidUndelegate
	intermediaryAcc, found := k.GetIntermediaryAccountFromLockId(ctx, lockID)
	if !found {
		return 0, types.ErrNotSuperfluidUsedLockup
	}

	// undelegate all
	err = k.SuperfluidUndelegate(ctx, sender, lockID)
	if err != nil {
		return 0, err
	}

	// unbond partial or full locked amount
	newLockID, err := k.unbondLock(ctx, lockID, sender, coins)
	if err != nil {
		return 0, err
	}

	// check new lock id
	// If unbond amount == locked amount, then the underlying lock was not split.
	// So we double check that newLockID == lockID, and return.
	// This has the same effect as calling SuperfluidUndelegate and then SuperfluidUnbondLock.
	// Otherwise unbond amount < locked amount, and the underlying lock was split.
	// lockID contains the amount still locked in the lockup module.
	// newLockID contains the amount unlocked.
	// We double check that newLockID != lockID and then proceed to re-delegate
	// the remainder (locked amount - unbond amount).
	if lock.Coins[0].IsEqual(coins[0]) {
		if newLockID != lockID {
			panic(fmt.Errorf("expected new lock id %v to = lock id %v", newLockID, lockID))
		}
		return lock.ID, nil
	} else {
		if newLockID == lockID {
			panic(fmt.Errorf("expected new lock id %v to != lock id %v", newLockID, lockID))
		}
	}

	// delete synthetic unlocking lock created in the last step of SuperfluidUndelegate
	synthdenom := unstakingSyntheticDenom(lock.Coins[0].Denom, intermediaryAcc.ValAddr)
	err = k.lk.DeleteSyntheticLockup(ctx, lockID, synthdenom)
	if err != nil {
		return 0, err
	}

	// re-delegate remainder
	err = k.SuperfluidDelegate(ctx, sender, lockID, intermediaryAcc.ValAddr)
	if err != nil {
		return 0, err
	}

	// create synthetic unlocking lock for newLockID
	err = k.createSyntheticLockup(ctx, newLockID, intermediaryAcc, unlockingStatus)
	if err != nil {
		return 0, err
	}
	return newLockID, nil
}

// unbondLock unlocks the underlying lock. Same lock id is returned if the amount to unlock
// is equal to the entire locked amount. Otherwise, the amount to unlock is less
// than the amount locked, it will return a new lock id which was created as an unlocking lock.
func (k Keeper) unbondLock(ctx sdk.Context, underlyingLockId uint64, sender string, coins sdk.Coins) (uint64, error) {
	lock, err := k.lk.GetLockByID(ctx, underlyingLockId)
	if err != nil {
		return 0, err
	}
	err = k.validateLockForSF(ctx, lock, sender)
	if err != nil {
		return 0, err
	}
	synthLock, _, err := k.lk.GetSyntheticLockupByUnderlyingLockId(ctx, underlyingLockId)
	if err != nil {
		return 0, err
	}
	// TODO: Use !found
	if synthLock == (lockuptypes.SyntheticLock{}) {
		return 0, types.ErrNotSuperfluidUsedLockup
	}
	if !synthLock.IsUnlocking() {
		return 0, types.ErrBondingLockupNotSupported
	}
	return k.lk.BeginForceUnlock(ctx, underlyingLockId, coins)
}

// alreadySuperfluidStaking returns true if underlying lock used in superfluid staking.
// This method would also return true for undelegating position for the lock.
func (k Keeper) alreadySuperfluidStaking(ctx sdk.Context, lockID uint64) bool {
	// We need to catch two cases:
	// (1) lockID has another superfluid bond
	// (2) lockID has a superfluid unbonding
	// we check (1) by looking for presence of an intermediary account lock ID connection
	// we check (2) (and re-check 1 for suredness) by looking for the existence of
	// synthetic locks for this.
	intermediaryAccAddr := k.GetLockIdIntermediaryAccountConnection(ctx, lockID)
	if !intermediaryAccAddr.Empty() {
		return true
	}

	synthLock, _, err := k.lk.GetSyntheticLockupByUnderlyingLockId(ctx, lockID)
	if err != nil {
		return false
	}
	// TODO: return found
	return synthLock != (lockuptypes.SyntheticLock{})
}

// mintOsmoTokensAndDelegate mints furyAmount of FURY tokens, and immediately delegate them to validator on behalf of intermediary account.
func (k Keeper) mintOsmoTokensAndDelegate(ctx sdk.Context, furyAmount sdk.Int, intermediaryAccount types.SuperfluidIntermediaryAccount) error {
	validator, err := k.validateValAddrForDelegate(ctx, intermediaryAccount.ValAddr)
	if err != nil {
		return err
	}

	err = osmoutils.ApplyFuncIfNoError(ctx, func(cacheCtx sdk.Context) error {
		bondDenom := k.sk.BondDenom(cacheCtx)
		coins := sdk.Coins{sdk.NewCoin(bondDenom, furyAmount)}
		err = k.bk.MintCoins(cacheCtx, types.ModuleName, coins)
		if err != nil {
			return err
		}
		k.bk.AddSupplyOffset(cacheCtx, bondDenom, furyAmount.Neg())
		err = k.bk.SendCoinsFromModuleToAccount(cacheCtx, types.ModuleName, intermediaryAccount.GetAccAddress(), coins)
		if err != nil {
			return err
		}

		// make delegation from module account to the validator
		// TODO: What happens here if validator is jailed, tombstoned, or unbonding
		// For now, we don't worry since worst case it errors, in which case we revert mint.
		_, err = k.sk.Delegate(cacheCtx,
			intermediaryAccount.GetAccAddress(),
			furyAmount, stakingtypes.Unbonded, validator, true)
		return err
	})
	return err
}

// forceUndelegateAndBurnOsmoTokens force undelegates furyAmount worth of delegation shares
// from delegations between intermediary account and valAddr.
// We take the returned tokens, and then immediately burn them.
func (k Keeper) forceUndelegateAndBurnOsmoTokens(ctx sdk.Context,
	furyAmount sdk.Int, intermediaryAcc types.SuperfluidIntermediaryAccount,
) error {
	valAddr, err := sdk.ValAddressFromBech32(intermediaryAcc.ValAddr)
	if err != nil {
		return err
	}
	// TODO: Better understand and decide between ValidateUnbondAmount and SharesFromTokens
	// briefly looked into it, did not understand whats correct.
	// TODO: ensure that intermediate account has at least furyAmount staked.
	shares, err := k.sk.ValidateUnbondAmount(
		ctx, intermediaryAcc.GetAccAddress(), valAddr, furyAmount,
	)
	if err == stakingtypes.ErrNoDelegation {
		return nil
	} else if err != nil {
		return err
	}
	err = osmoutils.ApplyFuncIfNoError(ctx, func(cacheCtx sdk.Context) error {
		undelegatedCoins, err := k.sk.InstantUndelegate(cacheCtx, intermediaryAcc.GetAccAddress(), valAddr, shares)
		if err != nil {
			return err
		}

		// TODO: Should we compare undelegatedCoins vs furyAmount?
		err = k.bk.SendCoinsFromAccountToModule(cacheCtx, intermediaryAcc.GetAccAddress(), types.ModuleName, undelegatedCoins)
		if err != nil {
			return err
		}
		err = k.bk.BurnCoins(cacheCtx, types.ModuleName, undelegatedCoins)
		if err != nil {
			return err
		}
		bondDenom := k.sk.BondDenom(cacheCtx)
		k.bk.AddSupplyOffset(cacheCtx, bondDenom, undelegatedCoins.AmountOf(bondDenom))

		return err
	})

	return err
}

// TODO: Need to (eventually) override the existing staking messages and queries, for undelegating, delegating, rewards, and redelegating, to all be going through all superfluid module.
// Want integrators to be able to use the same staking queries and messages
// Eugen’s point: Only rewards message needs to be updated. Rest of messages are fine
// Queries need to be updated
// We can do this at the very end though, since it just relates to queries.

// IterateBondedValidatorsByPower implements govtypes.StakingKeeper
func (k Keeper) IterateBondedValidatorsByPower(ctx sdk.Context, fn func(int64, stakingtypes.ValidatorI) bool) {
	k.sk.IterateBondedValidatorsByPower(ctx, fn)
}

// TotalBondedTokens implements govtypes.StakingKeeper
func (k Keeper) TotalBondedTokens(ctx sdk.Context) sdk.Int {
	return k.sk.TotalBondedTokens(ctx)
}

// IterateDelegations implements govtypes.StakingKeeper
// Iterates through staking keeper's delegations, and then all of the superfluid delegations.
func (k Keeper) IterateDelegations(ctx sdk.Context, delegator sdk.AccAddress, fn func(int64, stakingtypes.DelegationI) bool) {
	// call the callback with the non-superfluid delegations
	var index int64
	k.sk.IterateDelegations(ctx, delegator, func(i int64, delegation stakingtypes.DelegationI) (stop bool) {
		index = i
		return fn(i, delegation)
	})

	synthlocks := k.lk.GetAllSyntheticLockupsByAddr(ctx, delegator)
	for i, lock := range synthlocks {
		// get locked coin from the lock ID
		interim, ok := k.GetIntermediaryAccountFromLockId(ctx, lock.UnderlyingLockId)
		if !ok {
			continue
		}

		lock, err := k.lk.GetLockByID(ctx, lock.UnderlyingLockId)
		if err != nil {
			ctx.Logger().Error("lockup retrieval failed with underlying lock", "Lock", lock, "Error", err)
			continue
		}

		coin, err := lock.SingleCoin()
		if err != nil {
			ctx.Logger().Error("lock fails to meet expected invariant, it contains multiple coins", "Lock", lock, "Error", err)
			continue
		}

		// get fury-equivalent token amount
		amount, err := k.GetSuperfluidFURYTokens(ctx, interim.Denom, coin.Amount)
		if err != nil {
			ctx.Logger().Error("failed to get fury equivalent of token", "Denom", interim.Denom, "Amount", coin.Amount, "Error", err)
			continue
		}

		// get validator shares equivalent to the token amount
		valAddr, err := sdk.ValAddressFromBech32(interim.ValAddr)
		if err != nil {
			ctx.Logger().Error("failed to decode validator address", "Intermediary", interim.ValAddr, "LockID", lock.ID, "Error", err)
			continue
		}

		validator, found := k.sk.GetValidator(ctx, valAddr)
		if !found {
			ctx.Logger().Error("validator does not exist for lock", "Validator", valAddr, "LockID", lock.ID)
			continue
		}

		shares, err := validator.SharesFromTokens(amount)
		if err != nil {
			// tokens are not valid. continue.
			continue
		}

		// construct delegation and call callback
		delegation := stakingtypes.Delegation{
			DelegatorAddress: delegator.String(),
			ValidatorAddress: interim.ValAddr,
			Shares:           shares,
		}

		// if valid delegation has been found, increment delegation index
		fn(index+int64(i), delegation)
	}
}
