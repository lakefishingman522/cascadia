#!/bin/bash

KEYS[0]="dev0"
KEYS[1]="dev1"
KEYS[2]="dev2"
CHAINID="cascadia_6102-1"
MONIKER="localtestnet"
# MULTISIG_ADDRESS="cascadia1duc20j5qccrawl9n7url89lk5t23hur3w0rhem"
# Remember to change to other types of keyring like 'file' in-case exposing to outside world,
# otherwise your balance will be wiped quickly
# The keyring test does not require private key to steal tokens from you
KEYRING="test"
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
# Set dedicated home directory for the cascadiad instance
HOMEDIR="$HOME/.newtestcascadiad"
# to trace evm
#TRACE="--trace"
TRACE=""

# Path variables
CONFIG=$HOMEDIR/config/config.toml
APP_TOML=$HOMEDIR/config/app.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json

MNEMONIC_LOG="$HOME/utility/newtestcascadiad/log"

# validate dependencies are installed
command -v jq >/dev/null 2>&1 || {
    echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"
    exit 1
}

# used to exit on first error (any non-zero exit code)
set -e

# Reinstall daemon
# make install

# User prompt if an existing local node configuration is found.
if [ -d "$HOMEDIR" ]; then
	printf "\nAn existing folder at '%s' was found. You can choose to delete this folder and start a new local node with new keys from genesis. When declined, the existing local node is started. \n" "$HOMEDIR"
	echo "Overwrite the existing configuration and start a new local node? [y/n]"
	read -r overwrite
else
	overwrite="Y"
fi


# Setup local node if overwrite is set to Yes, otherwise skip setup
if [[ $overwrite == "y" || $overwrite == "Y" ]]; then
	# Remove the previous folder
	rm -rf "$HOMEDIR"

	# Set client config
	cascadiad config keyring-backend $KEYRING --home "$HOMEDIR"
	cascadiad config chain-id $CHAINID --home "$HOMEDIR"

	# If keys exist they should be deleted
	for KEY in "${KEYS[@]}"; do
		cascadiad keys add "$KEY" --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR"
	done

	# Set moniker and chain-id for Cascadia (Moniker can be anything, chain-id must be an integer)
	cascadiad init $MONIKER -o --chain-id $CHAINID --home "$HOMEDIR"

	jq '.app_state["staking"]["params"]["bond_denom"]="aCC"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["crisis"]["constant_fee"]["denom"]="aCC"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state.gov.deposit_params.min_deposit[0].denom="aCC"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["feemarket"]["block_gas"]="10000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# set gov proposing && voting period
	jq '.app_state.gov.params.max_deposit_period="120s"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state.gov.params.voting_period="120s"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# When upgrade to cosmos-sdk v0.47, use gov.params to edit the deposit params
	# check if the 'params' field exists in the genesis file
	if jq '.app_state.gov.params != null' "$GENESIS" | grep -q "true"; then
	jq '.app_state.gov.params.min_deposit[0].denom="aCC"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state.gov.params.max_deposit_period="120s"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state.gov.params.voting_period="120s"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	fi

	# Set gas limit in genesis
	jq '.consensus_params["block"]["max_gas"]="10000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# set custom pruning settings
	sed -i.bak 's/pruning = "default"/pruning = "custom"/g' "$APP_TOML"
	sed -i.bak 's/pruning-keep-recent = "0"/pruning-keep-recent = "2"/g' "$APP_TOML"
	sed -i.bak 's/pruning-interval = "0"/pruning-interval = "10"/g' "$APP_TOML"


	# ===========================================================
	# Update rpc laddr in config with node port
	
	node_rpc_port=27657
	node_api_port=27656
	node_pprof_port=7060
	node_proxy_app_port=27658
	node_grpc_port=10090
	node_grpc_web_port=10091
	node_evm_port=9545
	node_evm_socket_port=9546
	node_api_port=2317

	sed -i "s/^laddr = \"tcp:\/\/127.0.0.1:26657\"/laddr = \"tcp:\/\/127.0.0.1:${node_rpc_port}\"/" "$CONFIG"

	# Update p2p laddr in config with node port
	sed -i "s/^laddr = \"tcp:\/\/0.0.0.0:26656\"/laddr = \"tcp:\/\/0.0.0.0:${node_api_port}\"/" "$CONFIG"

	# Update p2p laddr in config with node port
	sed -i "s/^pprof_laddr =.*/pprof_laddr = \"localhost:${node_pprof_port}\"/" "$CONFIG"


	sed -i.bak "s/:26658/:${node_proxy_app_port}/g" "$CONFIG"


	# Update grpc in app with node port
	sed -i "s/^address = \"localhost:9090\"/address = \"0.0.0.0:${node_grpc_port}\"/" "$APP_TOML"

	# Update grpc web in app with node port
	sed -i "s/^address = \"localhost:9091\"/address = \"0.0.0.0:${node_grpc_web_port}\"/" "$APP_TOML"

	sed -i.bak "s/127.0.0.1:8545/0.0.0.0:${node_evm_port}/g" "$APP_TOML"
	sed -i.bak "s/127.0.0.1:8546/0.0.0.0:${node_evm_socket_port}/g" "$APP_TOML"

	sed -i.bak "s/:1317/:${node_api_port}/g" "$APP_TOML"
	# ===========================================================

	# Allocate genesis accounts (cosmos formatted addresses)
	for KEY in "${KEYS[@]}"; do
		cascadiad add-genesis-account "$KEY" 100000000000000000000000000aCC --keyring-backend $KEYRING --home "$HOMEDIR"
	done

	# bc is required to add these big numbers
	total_supply=$(echo "${#KEYS[@]} * 100000000000000000000000000" | bc)
	jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"][0]["amount"]=$total_supply' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# Sign genesis transaction
	cascadiad gentx "${KEYS[0]}" 1000000000000000000000aCC --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR"
	## In case you want to create multiple validators at genesis
	## 1. Back to `cascadiad keys add` step, init more keys
	## 2. Back to `cascadiad add-genesis-account` step, add balance for those
	## 3. Clone this ~/.cascadiad home directory into some others, let's say `~/.clonedCascadiad`
	## 4. Run `gentx` in each of those folders
	## 5. Copy the `gentx-*` folders under `~/.clonedCascadiad/config/gentx/` folders into the original `~/.cascadiad/config/gentx`

	# Collect genesis tx
	cascadiad collect-gentxs --home "$HOMEDIR"

	# Run this to ensure everything worked and that the genesis file is setup correctly
	cascadiad validate-genesis --home "$HOMEDIR"

	if [[ $1 == "pending" ]]; then
		echo "pending mode is on, please wait for the first block committed."
	fi
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
cascadiad start --metrics "$TRACE" --log_level debug --minimum-gas-prices=0.0001aCC --json-rpc.api eth,txpool,personal,net,debug,web3 --api.enable --home "$HOMEDIR" --rpc.laddr "tcp://0.0.0.0:26657" --json-rpc.address 0.0.0.0:8545 --json-rpc.ws-address 0.0.0.0:8546
