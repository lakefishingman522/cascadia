#!/bin/bash

KEYS[0]="dev0"
KEYS[1]="dev1"
KEYS[2]="dev2"
CHAINID="cascadia_6102-1"
MONIKER="localtestnet"
MULTISIG_ADDRESS="cascadia1duc20j5qccrawl9n7url89lk5t23hur3w0rhem"
# Remember to change to other types of keyring like 'file' in-case exposing to outside world,
# otherwise your balance will be wiped quickly
# The keyring test does not require private key to steal tokens from you
KEYRING="test"
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
# Set dedicated home directory for the cascadiad instance
HOMEDIR="$HOME/.testcascadiad"
# to trace evm
#TRACE="--trace"
TRACE=""

# Path variables
CONFIG=$HOMEDIR/config/config.toml
APP_TOML=$HOMEDIR/config/app.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json

MNEMONIC_LOG="$HOME/utility/testcascadiad/log"

# validate dependencies are installed
  command -v jq >/dev/null 2>&1 || {
    echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"
    exit 1
  }

  # used to exit on first error (any non-zero exit code)
  set -e

  

init() {
  # Reinstall daemon
#   make install 

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
	mkdir "$HOMEDIR"
	mkdir "$HOMEDIR/keyring-test"
	
    # Set client config
    # cascadiad config keyring-backend $KEYRING --home "$HOMEDIR"
    cp -r  ~/utility/testcascadiad/keyring-test/ $HOMEDIR
	cascadiad config chain-id $CHAINID --home "$HOMEDIR"

    # If keys exist they should be deleted
    # for KEY in "${KEYS[@]}"; do
    #   cascadiad keys add "$KEY" --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR" --output json >>"$MNEMONIC_LOG"
    # done
	cascadiad keys add dev0 --key --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR"

    # Set moniker and chain-id for Cascadia (Moniker can be anything, chain-id must be an integer)
    cascadiad init $MONIKER -o --chain-id $CHAINID --home "$HOMEDIR"

    jq '.app_state["staking"]["params"]["bond_denom"]="aCC"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
    jq '.app_state["crisis"]["constant_fee"]["denom"]="aCC"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
    jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="aCC"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
    jq '.app_state["feemarket"]["block_gas"]="10000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

    # Set gas limit in genesis
    jq '.consensus_params["block"]["max_gas"]="10000000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

    # set custom pruning settings
    sed -i.bak 's/pruning = "default"/pruning = "custom"/g' "$APP_TOML"
    sed -i.bak 's/pruning-keep-recent = "0"/pruning-keep-recent = "2"/g' "$APP_TOML"
    sed -i.bak 's/pruning-interval = "0"/pruning-interval = "10"/g' "$APP_TOML"

    sed -i.bak 's/127.0.0.1:26657/0.0.0.0:27657/g' "$CONFIG"

    # avoid confilct existing cascadia
    sed -i.bak 's/:9090/:10090/g' "$APP_TOML"
    sed -i.bak 's/:9091/:10091/g' "$APP_TOML"
    sed -i.bak 's/127.0.0.1:8545/0.0.0.0:9545/g' "$APP_TOML"
    sed -i.bak 's/:8546/:9546/g' "$APP_TOML"

    sed -i.bak 's/:6065/:7065/g' "$APP_TOML"
    sed -i.bak 's/:1317/:2317/g' "$APP_TOML"
    sed -i.bak 's/:8080/:9080/g' "$APP_TOML"

    sed -i.bak 's/:26656/:27656/g' "$CONFIG"
    sed -i.bak 's/:26658/:27658/g' "$CONFIG"
    sed -i.bak 's/:26660/:27660/g' "$CONFIG"
    sed -i.bak 's/:6060/:7060/g' "$CONFIG"

    sed -i.bak 's/cors_allowed_origins\s*=\s*\[\]/cors_allowed_origins = ["*",]/g' "$CONFIG"

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

  # Add the MultiSig address to the genesis.json file
  jq --arg ms_addr "$MULTISIG_ADDRESS" '.app_state["multisig_address"]=$ms_addr' "$GENESIS" >"$TMP_GENESIS"
  cp "$TMP_GENESIS" "$GENESIS"
  rm "$TMP_GENESIS"

  # Start the node (remove the --pruning=nothing flag if historical queries are not needed)
  cascadiad start --metrics "$TRACE" --log_level info --minimum-gas-prices=0.0001aCC --json-rpc.api eth,txpool,personal,net,debug,web3 --api.enable --home "$HOMEDIR" --rpc.laddr "tcp://0.0.0.0:27657" --json-rpc.address 0.0.0.0:9545 --json-rpc.ws-address 0.0.0.0:9546
}

goon() {
  # Reinstall daemon
  # make install

  # Start the node (remove the --pruning=nothing flag if historical queries are not needed)
  cascadiad start --metrics "$TRACE" --log_level info --minimum-gas-prices=0.0001aCC --json-rpc.api eth,txpool,personal,net,debug,web3 --api.enable --home "$HOMEDIR" --rpc.laddr "tcp://0.0.0.0:27657" --json-rpc.address 0.0.0.0:9545 --json-rpc.ws-address 0.0.0.0:9546

}

confirm(){

  total=$(cascadiad --home ~/.testcascadiad/ --node "tcp://localhost:27657" query bank total --output json | jq -r '.supply[0].amount')
  
  r0=$(cascadiad --home ~/.testcascadiad/ --node "tcp://localhost:27657" query bank balances cascadia1ahje4knm458pzhqntlvv6krpq8cnwmd7e3unmc --output json | jq -r '.balances[0].amount')
  r1=$(cascadiad --home ~/.testcascadiad/ --node "tcp://localhost:27657" query bank balances cascadia19sv2xv9qknp5vcmcswe0m3q4ghgtk89gqe7zdx --output json | jq -r '.balances[0].amount')
  r2=$(cascadiad --home ~/.testcascadiad/ --node "tcp://localhost:27657" query bank balances cascadia1fzdhv2hkq78aaar02dnmcaj48pres8pm37ag07 --output json | jq -r '.balances[0].amount')
  
  validator_outstanding_reward=$(cascadiad --home ~/.testcascadiad/ --node "tcp://localhost:27657" query distribution validator-outstanding-rewards cascadiavaloper1r4cdphjp4fph500jayxymve8sccu3jcryc5r39 --output json | jq -r '.rewards[0].amount' | sed 's/\..*//')
  distribution_reward=$(cascadiad --home ~/.testcascadiad/ --node "tcp://localhost:27657" query distribution rewards cascadia1r4cdphjp4fph500jayxymve8sccu3jcrm8e94k cascadiavaloper1r4cdphjp4fph500jayxymve8sccu3jcryc5r39 --output json | jq -r '.rewards[0].amount' | sed 's/\..*//')
  
  d0=$(cascadiad --home ~/.testcascadiad/ --node "tcp://localhost:27657" query bank balances cascadia1r4cdphjp4fph500jayxymve8sccu3jcrm8e94k --output json | jq -r '.balances[0].amount')
  d1=$(cascadiad --home ~/.testcascadiad/ --node "tcp://localhost:27657" query bank balances cascadia1nx0cqra2gz9ang548yx7x8hly7h7ld97t7qkat --output json | jq -r '.balances[0].amount')
  d2=$(cascadiad --home ~/.testcascadiad/ --node "tcp://localhost:27657" query bank balances cascadia1mmw9ykcgsj4fdfwmpnhep0znvje3nh9pjduta7 --output json | jq -r '.balances[0].amount')
  
  ff=$(cascadiad --home ~/.testcascadiad/ --node "tcp://localhost:27657" query distribution community-pool --output json | jq -r '.pool[0].amount'  | sed 's/\..*//')
  gg=$(cascadiad --home ~/.testcascadiad/ --node "tcp://localhost:27657" query distribution validator-outstanding-rewards cascadiavaloper1r4cdphjp4fph500jayxymve8sccu3jcryc5r39 --output json | jq -r '.rewards[0].amount' | sed 's/\..*//')
  
  result=$(echo "$total-$r0-$r1-$r2-$distribution_reward-$d0-$d1-$d2" | bc)
  echo "sum:" $(echo "$d0+$d1+$d2" | bc)
  echo "total, d0, d1, d2:" $total $d0 $d1 $d2 
  echo "validator_outstanding_reward, distribution_reward, r1:" $validator_outstanding_reward $distribution_reward $r1
  
  echo $result
  echo "total:" $total
  echo "$(echo "scale=4; ($r1-$distribution_reward)/$r1" | bc)"
  echo "$(echo "scale=4; ($r1-$distribution_reward)" | bc)"

}

case "$1" in
    init)
        init
        ;;
    goon)
        goon
        ;;
    confirm)
        confirm
        ;;
    *)
        echo "Usage: $0 {init|goon|confirm}"
        exit 1
esac
