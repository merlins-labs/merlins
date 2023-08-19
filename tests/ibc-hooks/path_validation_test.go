package ibc_hooks_test

import (
	"fmt"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctesting "github.com/cosmos/ibc-go/v4/testing"
)

// This sets up PFM on chainB and tests that it works as expected. We assume ChainA is merlins
func (suite *HooksTestSuite) SetupAndTestPFM(chainBId Chain, chainBName string, registryAddr sdk.AccAddress) {
	targetChain := suite.GetChain(chainBId)
	sendFrom := targetChain.SenderAccount.GetAddress()
	direction := suite.GetDirection(ChainA, chainBId)
	reverseDirection := suite.GetDirection(chainBId, ChainA)
	sender, receiver := suite.GetEndpoints(suite.GetDirection(ChainA, chainBId))

	merlinsApp := suite.chainA.GetMerlinsApp()
	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(merlinsApp.WasmKeeper)

	pfm_msg := fmt.Sprintf(`{"has_packet_forwarding": {"chain": "%s"}}`, chainBName)
	forwarding := suite.chainA.QueryContractJson(&suite.Suite, registryAddr, []byte(pfm_msg))
	suite.Require().False(forwarding.Bool())

	transferMsg := NewMsgTransfer(sdk.NewCoin("token0", sdk.NewInt(2000)), targetChain.SenderAccount.GetAddress().String(), sendFrom.String(), suite.GetSenderChannel(chainBId, ChainA), "")
	suite.FullSend(transferMsg, reverseDirection)
	tokenBA := suite.GetIBCDenom(chainBId, ChainA, "token0")

	ctx := suite.chainA.GetContext()

	msg := fmt.Sprintf(`{"propose_pfm":{"chain": "%s"}}`, chainBName)
	_, err := contractKeeper.Execute(ctx, registryAddr, sendFrom, []byte(msg), sdk.NewCoins(sdk.NewCoin(tokenBA, sdk.NewInt(1))))
	suite.Require().NoError(err)

	forwarding = suite.chainA.QueryContractJson(&suite.Suite, registryAddr, []byte(pfm_msg))
	suite.Require().False(forwarding.Bool())

	// Move forward one block
	suite.chainA.NextBlock()
	suite.chainA.Coordinator.IncrementTime()

	// Update both clients
	err = receiver.UpdateClient()
	suite.Require().NoError(err)
	err = sender.UpdateClient()
	suite.Require().NoError(err)

	events := ctx.EventManager().Events()
	packet0, err := ibctesting.ParsePacketFromEvents(events)
	suite.Require().NoError(err)
	result := suite.RelayPacketNoAck(packet0, direction) // No ack because it's a forward

	forwarding = suite.chainA.QueryContractJson(&suite.Suite, registryAddr, []byte(pfm_msg))
	suite.Require().False(forwarding.Bool())

	packet1, err := ibctesting.ParsePacketFromEvents(result.GetEvents())
	suite.Require().NoError(err)
	receiveResult, _ := suite.RelayPacket(packet1, reverseDirection)

	forwarding = suite.chainA.QueryContractJson(&suite.Suite, registryAddr, []byte(pfm_msg))
	suite.Require().False(forwarding.Bool())

	err = sender.UpdateClient()
	suite.Require().NoError(err)
	err = receiver.UpdateClient()
	suite.Require().NoError(err)

	ack, err := ibctesting.ParseAckFromEvents(receiveResult.GetEvents())
	suite.Require().NoError(err)

	err = sender.AcknowledgePacket(packet0, ack)
	suite.Require().NoError(err)

	// After the ack fully travels back to the initial chain, we consider PFM to be properly set
	forwarding = suite.chainA.QueryContractJson(&suite.Suite, registryAddr, []byte(pfm_msg))
	suite.Require().True(forwarding.Bool())
}

func (suite *HooksTestSuite) TestPathValidation() {
	owner := suite.chainA.SenderAccount.GetAddress()
	registryAddr, _, _, _ := suite.SetupCrosschainRegistry(ChainA)
	suite.setChainChannelLinks(registryAddr, ChainA)

	merlinsApp := suite.chainA.GetMerlinsApp()
	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(merlinsApp.WasmKeeper)

	msg := fmt.Sprintf(`{
		"modify_bech32_prefixes": {
		  "operations": [
			{"operation": "set", "chain_name": "merlins", "prefix": "fury"},
			{"operation": "set", "chain_name": "chainA", "prefix": "fury"},
			{"operation": "set", "chain_name": "chainB", "prefix": "fury"},
			{"operation": "set", "chain_name": "chainC", "prefix": "fury"}
		  ]
		}
	  }
	  `)
	_, err := contractKeeper.Execute(suite.chainA.GetContext(), registryAddr, owner, []byte(msg), sdk.NewCoins())
	suite.Require().NoError(err)
	suite.SetupAndTestPFM(ChainB, "chainB", registryAddr)
}