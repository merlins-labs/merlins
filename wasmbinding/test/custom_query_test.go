package wasmbinding

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/merlins-labs/merlins/v17/app"
	"github.com/merlins-labs/merlins/v17/app/apptesting"
	"github.com/merlins-labs/merlins/v17/wasmbinding/bindings"
)

func SetupCustomApp(t *testing.T, addr sdk.AccAddress) (*app.MerlinsApp, sdk.Context) {
	t.Helper()

	merlins, ctx := CreateTestInput()
	wasmKeeper := merlins.WasmKeeper

	storeReflectCode(t, ctx, merlins, addr)

	cInfo := wasmKeeper.GetCodeInfo(ctx, 1)
	require.NotNil(t, cInfo)

	return merlins, ctx
}

func TestQueryFullDenom(t *testing.T) {
	apptesting.SkipIfWSL(t)
	actor := RandomAccountAddress()
	merlins, ctx := SetupCustomApp(t, actor)

	reflect := instantiateReflectContract(t, ctx, merlins, actor)
	require.NotEmpty(t, reflect)

	// query full denom
	query := bindings.MerlinsQuery{
		FullDenom: &bindings.FullDenom{
			CreatorAddr: reflect.String(),
			Subdenom:    "ustart",
		},
	}
	resp := bindings.FullDenomResponse{}
	queryCustom(t, ctx, merlins, reflect, query, &resp)

	expected := fmt.Sprintf("factory/%s/ustart", reflect.String())
	require.EqualValues(t, expected, resp.Denom)
}

type ReflectQuery struct {
	Chain *ChainRequest `json:"chain,omitempty"`
}

type ChainRequest struct {
	Request wasmvmtypes.QueryRequest `json:"request"`
}

type ChainResponse struct {
	Data []byte `json:"data"`
}

func queryCustom(t *testing.T, ctx sdk.Context, merlins *app.MerlinsApp, contract sdk.AccAddress, request bindings.MerlinsQuery, response interface{}) {
	t.Helper()

	msgBz, err := json.Marshal(request)
	require.NoError(t, err)

	query := ReflectQuery{
		Chain: &ChainRequest{
			Request: wasmvmtypes.QueryRequest{Custom: msgBz},
		},
	}
	queryBz, err := json.Marshal(query)
	require.NoError(t, err)

	resBz, err := merlins.WasmKeeper.QuerySmart(ctx, contract, queryBz)
	require.NoError(t, err)
	var resp ChainResponse
	err = json.Unmarshal(resBz, &resp)
	require.NoError(t, err)
	err = json.Unmarshal(resp.Data, response)
	require.NoError(t, err)
}

func storeReflectCode(t *testing.T, ctx sdk.Context, merlins *app.MerlinsApp, addr sdk.AccAddress) {
	t.Helper()

	govKeeper := merlins.GovKeeper
	wasmCode, err := os.ReadFile("../testdata/fury_reflect.wasm")
	require.NoError(t, err)

	src := wasmtypes.StoreCodeProposalFixture(func(p *wasmtypes.StoreCodeProposal) {
		p.RunAs = addr.String()
		p.WASMByteCode = wasmCode
		checksum := sha256.Sum256(wasmCode)
		p.CodeHash = checksum[:]
	})

	// when stored
	storedProposal, err := govKeeper.SubmitProposal(ctx, src, false)
	require.NoError(t, err)

	// and proposal execute
	handler := govKeeper.Router().GetRoute(storedProposal.ProposalRoute())
	err = handler(ctx, storedProposal.GetContent())
	require.NoError(t, err)
}

func instantiateReflectContract(t *testing.T, ctx sdk.Context, merlins *app.MerlinsApp, funder sdk.AccAddress) sdk.AccAddress {
	t.Helper()

	initMsgBz := []byte("{}")
	contractKeeper := keeper.NewDefaultPermissionKeeper(merlins.WasmKeeper)
	codeID := uint64(1)
	addr, _, err := contractKeeper.Instantiate(ctx, codeID, funder, funder, initMsgBz, "demo contract", nil)
	require.NoError(t, err)

	return addr
}

func fundAccount(t *testing.T, ctx sdk.Context, merlins *app.MerlinsApp, addr sdk.AccAddress, coins sdk.Coins) {
	t.Helper()
	err := simapp.FundAccount(
		merlins.BankKeeper,
		ctx,
		addr,
		coins,
	)
	require.NoError(t, err)
}
