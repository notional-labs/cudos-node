#!/bin/bash

# This script is used to initialize, and run a localnet with 3 validators
CURDIR=$(pwd)
DENOM=${2:-"acudos"}
ERC20_ADDR="0x817bbDbC3e8A1204f3691d14bB44992841e3dB35"
rm -rf $CURDIR/cudos-data/
if [[ "$OSTYPE" == "darwin"* ]]; then
    killall SCREEN
else
    killall screen
fi

# start a testnet
build/cudos-noded testnet --keyring-backend=test

# change staking denom to acudos
cat $CURDIR/cudos-data/node0/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="acudos"' > $CURDIR/cudos-data/node0/config/tmp_genesis.json && mv $CURDIR/cudos-data/node0/config/tmp_genesis.json $CURDIR/cudos-data/node0/config/genesis.json

# update crisis variable to acudos
cat $CURDIR/cudos-data/node0/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="acudos"' > $CURDIR/cudos-data/node0/config/tmp_genesis.json && mv $CURDIR/cudos-data/node0/config/tmp_genesis.json $CURDIR/cudos-data/node0/config/genesis.json

# udpate gov genesis
cat $CURDIR/cudos-data/node0/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"]=[{"denom":"'$DENOM'","amount": "30000000000000000000000000"}]' > $CURDIR/cudos-data/node0/config/tmp_genesis.json && mv $CURDIR/cudos-data/node0/config/tmp_genesis.json $CURDIR/cudos-data/node0/config/genesis.json

# update mint genesis
cat $CURDIR/cudos-data/node0/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="acudos"' > $CURDIR/cudos-data/node0/config/tmp_genesis.json && mv $CURDIR/cudos-data/node0/config/tmp_genesis.json $CURDIR/cudos-data/node0/config/genesis.json

# update erc20 mapping
cat $CURDIR/cudos-data/node0/config/genesis.json | jq '.app_state["gravity"]["erc20_to_denoms"]=[{"erc20": "'$ERC20_ADDR'", "denom": "'$DENOM'" } ]' > $CURDIR/cudos-data/node0/config/tmp_genesis.json && mv $CURDIR/cudos-data/node0/config/tmp_genesis.json $CURDIR/cudos-data/node0/config/genesis.json

# update static valset update_test_genesis '.app_state["gravity"]["static_val_cosmos_addrs"]=[ "'$TEST0_ADDRESS'" ]'
TEST0_ADDRESS=$(build/cudos-noded keys show node0 -a --keyring-backend=test --home=$CURDIR/cudos-data/node0)
cat $CURDIR/cudos-data/node0/config/genesis.json | jq '.app_state["gravity"]["static_val_cosmos_addrs"]=[ "'$TEST0_ADDRESS'" ]' > $CURDIR/cudos-data/node0/config/tmp_genesis.json && mv $CURDIR/cudos-data/node0/config/tmp_genesis.json $CURDIR/cudos-data/node0/config/genesis.json
# change app.toml values

# validator 1
sed -i -E 's|swagger = false|swagger = true|g' $CURDIR/cudos-data/node0/config/app.toml

# validator2
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1316|g' $CURDIR/cudos-data/node1/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9088|g' $CURDIR/cudos-data/node1/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9089|g' $CURDIR/cudos-data/node1/config/app.toml
sed -i -E 's|swagger = false|swagger = true|g' $CURDIR/cudos-data/node1/config/app.toml

# validator3
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1315|g' $CURDIR/cudos-data/node2/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9086|g' $CURDIR/cudos-data/node2/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9087|g' $CURDIR/cudos-data/node2/config/app.toml
sed -i -E 's|swagger = false|swagger = true|g' $CURDIR/cudos-data/node2/config/app.toml

# change config.toml values

# validator1
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $CURDIR/cudos-data/node0/config/config.toml
# validator2
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26655|g' $CURDIR/cudos-data/node1/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $CURDIR/cudos-data/node1/config/config.toml
# validator3
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26652|g' $CURDIR/cudos-data/node2/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $CURDIR/cudos-data/node2/config/config.toml

# copy validator1 genesis file to validator2-3
cp $CURDIR/cudos-data/node0/config/genesis.json $CURDIR/cudos-data/node1/config/genesis.json
cp $CURDIR/cudos-data/node0/config/genesis.json $CURDIR/cudos-data/node2/config/genesis.json

echo "start all three validators"
screen -S validator1 -d -m build/cudos-noded start --home=$CURDIR/cudos-data/node0 
screen -S validator2 -d -m build/cudos-noded start --home=$CURDIR/cudos-data/node1 
screen -S validator3 -d -m build/cudos-noded start --home=$CURDIR/cudos-data/node2

echo $(build/cudos-noded keys show node0 -a --keyring-backend=test --home=$CURDIR/cudos-data/node0)
echo $(build/cudos-noded keys show node1 -a --keyring-backend=test --home=$CURDIR/cudos-data/node1)
echo $(build/cudos-noded keys show node2 -a --keyring-backend=test --home=$CURDIR/cudos-data/node2)