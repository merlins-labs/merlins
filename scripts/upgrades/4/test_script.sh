# Download a genesis.json for testing. The node that you this on will be your "validator"
# It should be on version v4.x

merlins init --chain-id=testing testing --home=$HOME/.merlins
merlins keys add validator --keyring-backend=test --home=$HOME/.merlins
merlins add-genesis-account $(merlins keys show validator -a --keyring-backend=test --home=$HOME/.merlins) 1000000000ufury,1000000000valtoken --home=$HOME/.merlins
sed -i -e "s/stake/ufury/g" $HOME/.merlins/config/genesis.json
merlins gentx validator 500000000ufury --commission-rate="0.0" --keyring-backend=test --home=$HOME/.merlins --chain-id=testing
merlins collect-gentxs --home=$HOME/.merlins

cat $HOME/.merlins/config/genesis.json | jq '.initial_height="711800"' > $HOME/.merlins/config/tmp_genesis.json && mv $HOME/.merlins/config/tmp_genesis.json $HOME/.merlins/config/genesis.json
cat $HOME/.merlins/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"]["denom"]="valtoken"' > $HOME/.merlins/config/tmp_genesis.json && mv $HOME/.merlins/config/tmp_genesis.json $HOME/.merlins/config/genesis.json
cat $HOME/.merlins/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"]["amount"]="100"' > $HOME/.merlins/config/tmp_genesis.json && mv $HOME/.merlins/config/tmp_genesis.json $HOME/.merlins/config/genesis.json
cat $HOME/.merlins/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="120s"' > $HOME/.merlins/config/tmp_genesis.json && mv $HOME/.merlins/config/tmp_genesis.json $HOME/.merlins/config/genesis.json
cat $HOME/.merlins/config/genesis.json | jq '.app_state["staking"]["params"]["min_commission_rate"]="0.050000000000000000"' > $HOME/.merlins/config/tmp_genesis.json && mv $HOME/.merlins/config/tmp_genesis.json $HOME/.merlins/config/genesis.json

# Now setup a second full node, and peer it with this v3.0.0-rc0 node.

# start the chain on both machines
merlins start
# Create proposals

merlins tx gov submit-proposal --title="existing passing prop" --description="passing prop"  --from=validator --deposit=1000valtoken --chain-id=testing --keyring-backend=test --broadcast-mode=block  --type="Text"
merlins tx gov vote 1 yes --from=validator --keyring-backend=test --chain-id=testing --yes
merlins tx gov submit-proposal --title="prop with enough fury deposit" --description="prop w/ enough deposit"  --from=validator --deposit=500000000ufury --chain-id=testing --keyring-backend=test --broadcast-mode=block  --type="Text"
# Check that we have proposal 1 passed, and proposal 2 in deposit period
merlins q gov proposals
# CHeck that validator commission is under min_commission_rate
merlins q staking validators
# Wait for upgrade block.
# Upgrade happened
# your full node should have crashed with consensus failure

# Now we test post-upgrade behavior is as intended

# Everything in deposit stayed in deposit
merlins q gov proposals
# Check that commissions was bumped to min_commission_rate
merlins q staking validators
# pushes 2 into voting period
merlins tx gov deposit 2 1valtoken --from=validator --keyring-backend=test --chain-id=testing --yes