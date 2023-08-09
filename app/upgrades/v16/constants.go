package v16

import (
	"github.com/merlins-labs/merlins/v17/app/upgrades"

	store "github.com/cosmos/cosmos-sdk/store/types"

	cltypes "github.com/merlins-labs/merlins/v17/x/concentrated-liquidity/types"
	cosmwasmpooltypes "github.com/osmosis-labs/osmosis/v17/x/cosmwasmpool/types"
)

// UpgradeName defines the on-chain upgrade name for the Merlins v16 upgrade.
const UpgradeName = "v16"

// new token factory parameters
//
// at the current gas price of 0.0025ufury, this corresponds to 0.1 FURY per
// denom creation.
const NewDenomCreationGasConsume uint64 = 40_000_000

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{cltypes.StoreKey, cosmwasmpooltypes.StoreKey},
		Deleted: []string{},
	},
}
