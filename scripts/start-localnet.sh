#!/bin/bash

# This script is used to run an existing localnet with 3 validators

# Stop any running nodes if they exist
./scripts/stop-nodes.sh
CURDIR=$(pwd)

echo "start all three validators"
screen -S validator1 -d -m build/cudos-noded start --home=$CURDIR/cudos-data/node0 
screen -S validator2 -d -m build/cudos-noded start --home=$CURDIR/cudos-data/node1 
screen -S validator3 -d -m build/cudos-noded start --home=$CURDIR/cudos-data/node2

echo $(build/cudos-noded keys show node0 -a --keyring-backend=test --home=$CURDIR/cudos-data/node0)
echo $(build/cudos-noded keys show node1 -a --keyring-backend=test --home=$CURDIR/cudos-data/node1)
echo $(build/cudos-noded keys show node2 -a --keyring-backend=test --home=$CURDIR/cudos-data/node2)