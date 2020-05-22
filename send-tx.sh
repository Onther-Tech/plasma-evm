#!/bin/bash

DATADIR_1=$HOME/.pls.dev

INIT_TIME=`date +%Y%m%d%H%M%S`
INIT_BALANCE=`build/bin/geth --exec "eth.getBalance(eth.accounts[0])" --datadir $DATADIR_1 attach`
INIT_BN=`build/bin/geth --exec "web3.eth.getBlock('latest').number" --datadir $DATADIR_1 attach`
INIT_TIME=`build/bin/geth --exec "eth.getBlock('latest').timestamp" --datadir $DATADIR_1 attach`

../chainload/chainload -id 16 -urls http://127.0.0.1:8547 -tps 5000 -dur 21s -senders 20
wait

echo "Sending tx is done"

TOT_TX=0

LAST_BN=`build/bin/geth --exec "eth.getBlock('latest').number" --datadir $DATADIR_1 attach`
LAST_TIME=`build/bin/geth --exec "eth.getBlock('latest').timestamp" --datadir $DATADIR_1 attach`
echo "Start BN is $INIT_BN"
echo "Last BN is $LAST_BN"

echo "Start timestamp is $INIT_TIME"
echo "LAST timestamp is $LAST_TIME"

TOTAL_BLOCK=``

echo "Calculate TPS"
for((i=$INIT_BN;i<$LAST_BN;i++))
do
  # send empty transaction
  TX=(`build/bin/geth --exec "eth.getBlock('latest').transactions" attach --datadir $DATADIR_1`)
  NUM_TX=${#TX[@]}
  TOT_TX=`expr $TOT_TX + $NUM_TX`
done
echo "Calculate TPS done"

CUR_TIME=`date +%Y%m%d%H%M%S`
CUR_BALANCE=`build/bin/geth --exec "eth.getBalance(eth.accounts[0])" --datadir $DATADIR_1 attach`

echo "Start Time is $INIT_TIME"
echo "Current Time is $CUR_TIME"
TIME_DIFF=`expr $CUR_TIME - $INIT_TIME`
echo "It takes $TIME_DIFF seconds."

# echo "Initial balance is $INIT_BALANCE"
# echo "Currnet balance is $CUR_BALANCE"

# BALANCE_DIFF=`expr $INIT_BALANCE - $CUR_BALANCE`
# echo "It use $BALANCE_DIFF"
echo "number of TPS is $TOT_TX"

DURATION=`expr $LAST_TIME - $INIT_TIME`
echo "Duration is $DURATION"

TPS=`expr $TOT_TX / $DURATION`
echo "TPS is $TPS"