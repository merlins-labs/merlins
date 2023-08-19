package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/osmosis-labs/osmosis/osmoutils/osmocli"
	clmodel "github.com/merlins-labs/merlins/v17/x/concentrated-liquidity/model"
	"github.com/merlins-labs/merlins/v17/x/concentrated-liquidity/types"
)

func NewTxCmd() *cobra.Command {
	txCmd := osmocli.TxIndexCmd(types.ModuleName)
	osmocli.AddTxCmd(txCmd, NewCreatePositionCmd)
	osmocli.AddTxCmd(txCmd, NewAddToPositionCmd)
	osmocli.AddTxCmd(txCmd, NewWithdrawPositionCmd)
	osmocli.AddTxCmd(txCmd, NewCreateConcentratedPoolCmd)
	osmocli.AddTxCmd(txCmd, NewCollectSpreadRewardsCmd)
	osmocli.AddTxCmd(txCmd, NewCollectIncentivesCmd)
	osmocli.AddTxCmd(txCmd, NewFungifyChargedPositionsCmd)
	return txCmd
}

var poolIdFlagOverride = map[string]string{
	"poolid": FlagPoolId,
}

func NewCreateConcentratedPoolCmd() (*osmocli.TxCliDesc, *clmodel.MsgCreateConcentratedPool) {
	return &osmocli.TxCliDesc{
		Use:     "create-pool [denom-0] [denom-1] [tick-spacing] [spread-factor]",
		Short:   "create a concentrated liquidity pool with the given denom pair, tick spacing, and spread factor",
		Long:    "denom-1 (the quote denom), tick spacing, and spread factors must all be authorized by the concentrated liquidity module",
		Example: "merlins tx concentratedliquidity create-pool uion ufury 100 0.01 --from val --chain-id merlins-1 -b block --keyring-backend test --fees 1000ufury",
	}, &clmodel.MsgCreateConcentratedPool{}
}

func NewCreatePositionCmd() (*osmocli.TxCliDesc, *types.MsgCreatePosition) {
	return &osmocli.TxCliDesc{
		Use:     "create-position [pool-id] [lower-tick] [upper-tick] [tokensProvided] [token-0-min-amount] [token-1-min-amount]",
		Short:   "create or add to existing concentrated liquidity position",
		Example: "merlins tx concentratedliquidity create-position 1 \"[-69082]\" 69082 10000ufury,10000uion 0 0 --from val --chain-id merlins-1 -b block --keyring-backend test --fees 1000ufury",
	}, &types.MsgCreatePosition{}
}

func NewAddToPositionCmd() (*osmocli.TxCliDesc, *types.MsgAddToPosition) {
	return &osmocli.TxCliDesc{
		Use:     "add-to-position [position-id] [token-0] [token-1]",
		Short:   "add to an existing concentrated liquidity position",
		Example: "merlins tx concentratedliquidity add-to-position 10 1000000000ufury 10000000uion --from val --chain-id localmerlins -b block --keyring-backend test --fees 1000000ufury",
	}, &types.MsgAddToPosition{}
}

func NewWithdrawPositionCmd() (*osmocli.TxCliDesc, *types.MsgWithdrawPosition) {
	return &osmocli.TxCliDesc{
		Use:     "withdraw-position [position-id] [liquidity]",
		Short:   "withdraw from an existing concentrated liquidity position",
		Example: "merlins tx concentratedliquidity withdraw-position 1 1000 --from val --chain-id localmerlins --keyring-backend=test --fees=1000ufury",
	}, &types.MsgWithdrawPosition{}
}

func NewCollectSpreadRewardsCmd() (*osmocli.TxCliDesc, *types.MsgCollectSpreadRewards) {
	return &osmocli.TxCliDesc{
		Use:     "collect-spread-rewards [position-ids]",
		Short:   "collect spread rewards from liquidity position(s)",
		Example: "merlins tx concentratedliquidity collect-spread-rewards 998 --from val --chain-id localmerlins -b block --keyring-backend test --fees 1000000ufury",
	}, &types.MsgCollectSpreadRewards{}
}

func NewCollectIncentivesCmd() (*osmocli.TxCliDesc, *types.MsgCollectIncentives) {
	return &osmocli.TxCliDesc{
		Use:     "collect-incentives [position-ids]",
		Short:   "collect incentives from liquidity position(s)",
		Example: "merlins tx concentratedliquidity collect-incentives 1 --from val --chain-id localmerlins -b block --keyring-backend test --fees 10000ufury",
	}, &types.MsgCollectIncentives{}
}

func NewFungifyChargedPositionsCmd() (*osmocli.TxCliDesc, *types.MsgFungifyChargedPositions) {
	return &osmocli.TxCliDesc{
		Use:     "fungify-positions [position-ids]",
		Short:   "Combine fully charged positions within the same range into a new single fully charged position",
		Example: "merlins tx concentratedliquidity fungify-positions 1,2 --from val --keyring-backend test -b=block --chain-id=localmerlins --gas=1000000 --fees 20000ufury",
	}, &types.MsgFungifyChargedPositions{}
}

// NewCmdCreateConcentratedLiquidityPoolsProposal implements a command handler for create concentrated liquidity pool proposal
func NewCmdCreateConcentratedLiquidityPoolsProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-concentratedliquidity-pool-proposal [flags]",
		Args:  cobra.ExactArgs(0),
		Short: "Submit a create concentrated liquidity pool proposal",
		Long: strings.TrimSpace(`Submit a create concentrated liquidity pool proposal.

Passing in FlagPoolRecords separated by commas would be parsed automatically to pairs of pool records.
Ex) --pool-records=uion,ufury,100,0.003,stake,ufury,1000,0.005 ->
[uion<>ufury, tickSpacing 100, spreadFactor 0.3%]
[stake<>ufury, tickSpacing 1000, spreadFactor 0.5%]

		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			content, err := parseCreateConcentratedLiquidityPoolArgsToContent(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")
	cmd.Flags().Bool(govcli.FlagIsExpedited, false, "If true, makes the proposal an expedited one")
	cmd.Flags().String(govcli.FlagProposal, "", "Proposal file path (if this path is given, other proposal flags are ignored)")
	cmd.Flags().String(FlagPoolRecords, "", "The pool records array")

	return cmd
}

func NewTickSpacingDecreaseProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tick-spacing-decrease-proposal [flags]",
		Args:  cobra.ExactArgs(0),
		Short: "Submit a tick spacing decrease proposal",
		Long: strings.TrimSpace(`Submit a tick spacing decrease proposal.

Passing in FlagPoolIdToTickSpacingRecords separated by commas would be parsed automatically to pairs of PoolIdToTickSpacing records.
Ex) --pool-tick-spacing-records=1,10,5,1 -> [(poolId 1, newTickSpacing 10), (poolId 5, newTickSpacing 1)]
Note: The new tick spacing value must be less than the current tick spacing value.

		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			content, err := parsePoolIdToTickSpacingRecordsArgsToContent(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")
	cmd.Flags().Bool(govcli.FlagIsExpedited, false, "If true, makes the proposal an expedited one")
	cmd.Flags().String(govcli.FlagProposal, "", "Proposal file path (if this path is given, other proposal flags are ignored)")
	cmd.Flags().String(FlagPoolIdToTickSpacingRecords, "", "The pool ID to new tick spacing records array")

	return cmd
}

func parseCreateConcentratedLiquidityPoolArgsToContent(cmd *cobra.Command) (govtypes.Content, error) {
	title, err := cmd.Flags().GetString(govcli.FlagTitle)
	if err != nil {
		return nil, err
	}

	description, err := cmd.Flags().GetString(govcli.FlagDescription)
	if err != nil {
		return nil, err
	}

	poolRecords, err := parsePoolRecords(cmd)
	if err != nil {
		return nil, err
	}

	content := &types.CreateConcentratedLiquidityPoolsProposal{
		Title:       title,
		Description: description,
		PoolRecords: poolRecords,
	}

	return content, nil
}

func parsePoolIdToTickSpacingRecordsArgsToContent(cmd *cobra.Command) (govtypes.Content, error) {
	title, err := cmd.Flags().GetString(govcli.FlagTitle)
	if err != nil {
		return nil, err
	}

	description, err := cmd.Flags().GetString(govcli.FlagDescription)
	if err != nil {
		return nil, err
	}

	poolIdToTickSpacingRecords, err := parsePoolIdToTickSpacingRecords(cmd)
	if err != nil {
		return nil, err
	}

	content := &types.TickSpacingDecreaseProposal{
		Title:                      title,
		Description:                description,
		PoolIdToTickSpacingRecords: poolIdToTickSpacingRecords,
	}
	return content, nil
}

func parsePoolIdToTickSpacingRecords(cmd *cobra.Command) ([]types.PoolIdToTickSpacingRecord, error) {
	assetsStr, err := cmd.Flags().GetString(FlagPoolIdToTickSpacingRecords)
	if err != nil {
		return nil, err
	}

	assets := strings.Split(assetsStr, ",")

	if len(assets)%2 != 0 {
		return nil, fmt.Errorf("poolIdToTickSpacingRecords must be a list of pairs of poolId and newTickSpacing")
	}

	poolIdToTickSpacingRecords := []types.PoolIdToTickSpacingRecord{}
	i := 0
	for i < len(assets) {
		poolId, err := strconv.Atoi(assets[i])
		if err != nil {
			return nil, err
		}
		newTickSpacing, err := strconv.Atoi(assets[i+1])
		if err != nil {
			return nil, err
		}

		poolIdToTickSpacingRecords = append(poolIdToTickSpacingRecords, types.PoolIdToTickSpacingRecord{
			PoolId:         uint64(poolId),
			NewTickSpacing: uint64(newTickSpacing),
		})

		// increase counter by the next 2
		i = i + 2
	}

	return poolIdToTickSpacingRecords, nil
}

func parsePoolRecords(cmd *cobra.Command) ([]types.PoolRecord, error) {
	poolRecordsStr, err := cmd.Flags().GetString(FlagPoolRecords)
	if err != nil {
		return nil, err
	}

	poolRecords := strings.Split(poolRecordsStr, ",")

	if len(poolRecords)%4 != 0 {
		return nil, fmt.Errorf("poolRecords must be a list of denom0, denom1, tickSpacing, and spreadFactor")
	}

	finalPoolRecords := []types.PoolRecord{}
	i := 0
	for i < len(poolRecords) {
		denom0 := poolRecords[i]
		denom1 := poolRecords[i+1]

		tickSpacing, err := strconv.Atoi(poolRecords[i+2])
		if err != nil {
			return nil, err
		}

		spreadFactorStr := poolRecords[i+3]
		spreadFactor, err := sdk.NewDecFromStr(spreadFactorStr)
		if err != nil {
			return nil, err
		}

		finalPoolRecords = append(finalPoolRecords, types.PoolRecord{
			Denom0:       denom0,
			Denom1:       denom1,
			TickSpacing:  uint64(tickSpacing),
			SpreadFactor: spreadFactor,
		})

		// increase counter by the next 4
		i = i + 4
	}

	return finalPoolRecords, nil
}
