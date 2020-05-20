#!/bin/bash

DATADIR_1=$HOME/.pls.dev

for((i=0;i<500;i++))
do
  # send empty transaction
  build/bin/geth --exec "web3.eth.sendTransaction({from: eth.accounts[0], to:eth.accounts[0], value: 0})" --datadir $DATADIR_1 attach
done

echo "tx1 done"
