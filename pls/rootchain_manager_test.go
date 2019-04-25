package pls

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/accounts/keystore"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/consensus"
	"github.com/Onther-Tech/plasma-evm/consensus/ethash"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/epochhandler"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/ethertoken"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/mintabletoken"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchain"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/token"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/bloombits"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/core/vm"
	"github.com/Onther-Tech/plasma-evm/crypto"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/ethdb"
	"github.com/Onther-Tech/plasma-evm/event"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/miner"
	"github.com/Onther-Tech/plasma-evm/miner/epoch"
	"github.com/Onther-Tech/plasma-evm/node"
	"github.com/Onther-Tech/plasma-evm/p2p"
	"github.com/Onther-Tech/plasma-evm/params"
	"github.com/Onther-Tech/plasma-evm/pls/gasprice"
	"github.com/Onther-Tech/plasma-evm/plsclient"
	"github.com/Onther-Tech/plasma-evm/rpc"
	"github.com/Onther-Tech/plasma-evm/tx"
	"github.com/mattn/go-colorable"
)

var (
	loglevel = flag.Int("loglevel", 4, "verbosity of logs")

	rootchainUrl   = "ws://localhost:8546"
	plasmachainUrl = "http://localhost:8547"

	operatorKey, _   = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	operator         = crypto.PubkeyToAddress(operatorKey.PublicKey)
	challengerKey, _ = crypto.HexToECDSA("78ae75d1cd5960d87e76a69760cb451a58928eee7890780c352186d23094a114")
	challenger       = crypto.PubkeyToAddress(challengerKey.PublicKey)
	operatorOpt      = bind.NewKeyedTransactor(operatorKey)

	addr1 = common.HexToAddress("0x5df7107c960320b90a3d7ed9a83203d1f98a811d")
	addr2 = common.HexToAddress("0x3cd9f729c8d882b851f8c70fb36d22b391a288cd")
	addr3 = common.HexToAddress("0x57ab89f4eabdffce316809d790d5c93a49908510")
	addr4 = common.HexToAddress("0x6c278df36922fea54cf6f65f725267e271f60dd9")
	addrs = []common.Address{addr1, addr2, addr3, addr4}

	key1, _ = crypto.HexToECDSA("78ae75d1cd5960d87e76a69760cb451a58928eee7890780c352186d23094a115")
	key2, _ = crypto.HexToECDSA("bfaa65473b85b3c33b2f5ddb511f0f4ef8459213ada2920765aaac25b4fe38c5")
	key3, _ = crypto.HexToECDSA("067394195895a82e685b000e592f771f7899d77e87cc8c79110e53a2f0b0b8fc")
	key4, _ = crypto.HexToECDSA("ae03e057a5b117295db86079ba4c8505df6074cdc54eec62f2050e677e5d4e66")
	keys    = []*ecdsa.PrivateKey{key1, key2, key3, key4}

	locks = map[common.Address]*sync.Mutex{
		operator: &sync.Mutex{},
		addr1:    &sync.Mutex{},
		addr2:    &sync.Mutex{},
		addr3:    &sync.Mutex{},
		addr4:    &sync.Mutex{},
	}

	operatorNonceRootChain uint64 = 0
	addr1NonceRootChain    uint64 = 0
	addr2NonceRootChain    uint64 = 0
	addr3NonceRootChain    uint64 = 0
	addr4NonceRootChain    uint64 = 0
	noncesRootChain               = map[common.Address]*uint64{
		operator: &operatorNonceRootChain,
		addr1:    &addr1NonceRootChain,
		addr2:    &addr2NonceRootChain,
		addr3:    &addr3NonceRootChain,
		addr4:    &addr4NonceRootChain,
	}

	operatorNonceChildChain uint64 = 0
	addr1NonceChildChain    uint64 = 0
	addr2NonceChildChain    uint64 = 0
	addr3NonceChildChain    uint64 = 0
	addr4NonceChildChain    uint64 = 0
	noncesChildChain               = map[common.Address]*uint64{
		operator: &operatorNonceChildChain,
		addr1:    &addr1NonceChildChain,
		addr2:    &addr2NonceChildChain,
		addr3:    &addr3NonceChildChain,
		addr4:    &addr4NonceChildChain,
	}

	opt1 = bind.NewKeyedTransactor(key1)
	opt2 = bind.NewKeyedTransactor(key2)
	opt3 = bind.NewKeyedTransactor(key3)
	opt4 = bind.NewKeyedTransactor(key4)

	opts = map[common.Address]*bind.TransactOpts{
		operator: operatorOpt,
		addr1:    opt1,
		addr2:    opt2,
		addr3:    opt3,
		addr4:    opt4,
	}

	empty32Bytes = common.Hash{}

	// contracts
	mintableToken     *mintabletoken.MintableToken
	mintableTokenAddr common.Address

	etherToken     *ethertoken.EtherToken
	etherTokenAddr common.Address

	etherTokenInChildChain     *ethertoken.EtherToken
	etherTokenAddrInChildChain common.Address

	// blockchain
	canonicalSeed = 1
	engine        = ethash.NewFaker()

	// node
	testNodeKey, _ = crypto.GenerateKey()
	testNodeConfig = &node.Config{
		Name: "test node",
		P2P:  p2p.Config{PrivateKey: testNodeKey},
	}

	// pls ~ rootchain
	testVmConfg   = vm.Config{EnablePreimageRecording: true}
	testPlsConfig = &DefaultConfig
	ethClient     *ethclient.Client

	// pls ~ plasmachain
	plsClient *plsclient.Client

	testTxPoolConfig = &core.DefaultTxPoolConfig

	// rootchain contract
	NRELength               = big.NewInt(2)
	development             = false
	swapEnabledInRootChain  = false
	swapEnabledInChildChain = true

	// transaction
	defaultGasPrice        = big.NewInt(1) // 1 Gwei
	defaultValue           = big.NewInt(0)
	defaultGasLimit uint64 = 7000000
	maxTxFee        *big.Int

	err error
)

func init() {
	log.PrintOrigins(true)
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*loglevel), log.StreamHandler(colorable.NewColorableStderr(), log.TerminalFormat(true))))

	testTxPoolConfig.Journal = ""
	testPlsConfig.TxPool = *testTxPoolConfig
	testPlsConfig.Operator = accounts.Account{Address: operator}
	testPlsConfig.Challenger = accounts.Account{Address: challenger}
	testPlsConfig.NodeMode = ModeOperator

	testPlsConfig.RootChainURL = rootchainUrl

	testPlsConfig.TxConfig.Interval = 1 * time.Second
	testPlsConfig.MinerRecommit = 10 * time.Second

	ethClient, err = ethclient.Dial(testPlsConfig.RootChainURL)
	if err != nil {
		log.Error("Failed to connect rootchian provider", "err", err)
	}

	networkId, err := ethClient.NetworkID(context.Background())
	if err != nil {
		log.Error("Failed to get network id", "err", err)
	}
	testPlsConfig.RootChainNetworkID = networkId.Uint64()
	testPlsConfig.TxConfig.ChainId = networkId

	keys = []*ecdsa.PrivateKey{key1, key2, key3, key4}
	addrs = []common.Address{addr1, addr2, addr3, addr4}

	operatorNonceRootChain, err = ethClient.NonceAt(context.Background(), operator, nil)
	addr1NonceRootChain, _ = ethClient.NonceAt(context.Background(), addr1, nil)
	addr2NonceRootChain, _ = ethClient.NonceAt(context.Background(), addr2, nil)
	addr3NonceRootChain, _ = ethClient.NonceAt(context.Background(), addr3, nil)
	addr4NonceRootChain, _ = ethClient.NonceAt(context.Background(), addr4, nil)

	for _, opt := range opts {
		opt.GasLimit = defaultGasLimit
	}

	maxTxFee = new(big.Int).Mul(defaultGasPrice, big.NewInt(int64(defaultGasLimit)))
}

func TestScenario1(t *testing.T) {
	rcm, stopFn, err := makeManager()
	defer stopFn()

	if err != nil {
		t.Fatalf("Failed to make rootchian manager: %v", err)
	}

	NRELength, err := rcm.NRELength()

	if err != nil {
		t.Fatalf("Failed to get NRELength: %v", err)
	}

	startETHDeposit(t, rcm, key1, ether(1))
	startETHDeposit(t, rcm, key2, ether(1))
	startETHDeposit(t, rcm, key3, ether(1))
	startETHDeposit(t, rcm, key4, ether(1))

	wait(3)

	numEROs, _ := rcm.rootchainContract.GetNumEROs(baseCallOpt)

	if numEROs.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("numEROs should not be 0")
	}

	events := rcm.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer events.Unsubscribe()

	if err = rcm.Start(); err != nil {
		t.Fatalf("Failed to start rootchain manager: %v", err)
	}

	timer := time.NewTimer(1 * time.Minute)
	go func() {
		<-timer.C
		t.Fatal("Out of time")
	}()

	var i uint64

	for i = 0; i < NRELength.Uint64(); {
		makeSampleTx(rcm)
		i++
		ev := <-events.Chan()

		blockInfo := ev.Data.(core.NewMinedBlockEvent)

		if rcm.minerEnv.IsRequest {
			t.Fatal("Block should not be request block, but it is not. blockNumber:", blockInfo.Block.NumberU64())
		}
	}

	ev := <-events.Chan()
	blockInfo := ev.Data.(core.NewMinedBlockEvent)
	if !rcm.minerEnv.IsRequest {
		t.Fatal("Block should be request block", "blockNumber", blockInfo.Block.NumberU64())
	}

	for i = 0; i < NRELength.Uint64()*2; {
		makeSampleTx(rcm)
		i++
		ev := <-events.Chan()
		blockInfo := ev.Data.(core.NewMinedBlockEvent)
		makeSampleTx(rcm)

		if rcm.minerEnv.IsRequest {
			t.Fatal("Block should not be request block", "blockNumber", blockInfo.Block.NumberU64())
		}
	}

	log.Info("test finished")
	return
}

// TestScenario2 tests enter and exit between root chain & plasma chain
func TestScenario2(t *testing.T) {
	pls, rpcServer, dir, err := makePls()
	defer os.RemoveAll(dir)

	if err != nil {
		t.Fatalf("Failed to make pls service: %v", err)
	}
	defer pls.Stop()
	defer rpcServer.Stop()

	if err := pls.rootchainManager.Start(); err != nil {
		t.Fatalf("Failed to start RootChainManager: %v", err)
	}
	pls.protocolManager.Start(1)

	rpcClient := rpc.DialInProc(rpcServer)

	// assign to global variable
	plsClient = plsclient.NewClient(rpcClient)

	rcm := pls.rootchainManager

	NRELength, err := rcm.NRELength()
	if err != nil {
		t.Fatalf("Failed to get NRELength: %v", err)
	}

	// balance check in root chain before enter
	balances1 := getEtherTokenBalances(addrs)

	enterAmount := ether(1)

	// make enter request
	startETHDeposit(t, rcm, key1, enterAmount)
	startETHDeposit(t, rcm, key2, enterAmount)
	startETHDeposit(t, rcm, key3, enterAmount)
	startETHDeposit(t, rcm, key4, enterAmount)

	balances2 := getEtherTokenBalances(addrs)

	for i, balance1 := range balances1 {
		balance2 := balances2[i]
		if err := checkBalance(balance1, balance2, new(big.Int).Neg(enterAmount), nil, "check enter request result"); err != nil {
			t.Fatal(err)
		}
	}

	numEROs, _ := rcm.rootchainContract.GetNumEROs(baseCallOpt)

	if numEROs.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("numEROs should not be 0")
	}

	events := rcm.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer events.Unsubscribe()

	timer := time.NewTimer(120 * time.Second)
	go func() {
		<-timer.C
		t.Fatal("Out of time")
	}()

	var i uint64

	// NRB#1 : deploy EtherToken in child chain
	deployEtherTokenInChildChain(t)
	ev := <-events.Chan()

	blockInfo := ev.Data.(core.NewMinedBlockEvent)

	if rcm.minerEnv.IsRequest {
		t.Fatal("Block should not be request block, but it is not. blockNumber:", blockInfo.Block.NumberU64())
	}

	// map EtherToken address
	setNonce(operatorOpt, &operatorNonceRootChain) // for NRB#1 submit
	setNonce(operatorOpt, &operatorNonceRootChain) // for ether token address map

	tx, err := rcm.rootchainContract.MapRequestableContractByOperator(operatorOpt, etherTokenAddr, etherTokenAddrInChildChain)
	if err != nil {
		t.Errorf("Failed to map EtherToken and PEtherToken: %v", err)
	}
	waitTx(tx.Hash())

	// #1 NRE
	for i = 0; i < NRELength.Uint64()-1; i++ {
		makeSampleTx(rcm)
		ev := <-events.Chan()

		blockInfo := ev.Data.(core.NewMinedBlockEvent)

		if rcm.minerEnv.IsRequest {
			t.Fatal("Block should not be request block, but it is not. blockNumber:", blockInfo.Block.NumberU64())
		}
	}

	// #2 empty ORE

	// #3 NRE
	for i = 0; i < NRELength.Uint64(); i++ {
		makeSampleTx(rcm)
		ev := <-events.Chan()
		blockInfo := ev.Data.(core.NewMinedBlockEvent)

		if rcm.minerEnv.IsRequest {
			t.Fatal("Block should not be request block", "blockNumber", blockInfo.Block.NumberU64())
		}
	}

	// #4 ORE
	withdrawalAmount := ether(0.9)
	for i = 0; i < 1; i++ {
		ev := <-events.Chan()
		blockInfo := ev.Data.(core.NewMinedBlockEvent)
		if !rcm.minerEnv.IsRequest {
			t.Fatal("Block should be request block", "blockNumber", blockInfo.Block.NumberU64())
		}
	}

	// make exit request
	startETHWithdraw(t, rcm, key1, withdrawalAmount, big.NewInt(int64(rcm.state.costERO)))
	startETHWithdraw(t, rcm, key2, withdrawalAmount, big.NewInt(int64(rcm.state.costERO)))
	startETHWithdraw(t, rcm, key3, withdrawalAmount, big.NewInt(int64(rcm.state.costERO)))
	startETHWithdraw(t, rcm, key4, withdrawalAmount, big.NewInt(int64(rcm.state.costERO)))

	// #5 NRE
	// swap PETH to PEtherToken
	setNonce(opt1, noncesChildChain[addr1])
	setNonce(opt2, noncesChildChain[addr2])
	setNonce(opt3, noncesChildChain[addr3])
	setNonce(opt4, noncesChildChain[addr4])

	opt1.Value = withdrawalAmount
	opt2.Value = withdrawalAmount
	opt3.Value = withdrawalAmount
	opt4.Value = withdrawalAmount

	if _, err := etherTokenInChildChain.SwapFromEth(opt1); err != nil {
		log.Error("Failed to swap PEtherToken", "err", err)
	}
	if _, err := etherTokenInChildChain.SwapFromEth(opt2); err != nil {
		log.Error("Failed to swap PEtherToken", "err", err)
	}
	if _, err := etherTokenInChildChain.SwapFromEth(opt3); err != nil {
		log.Error("Failed to swap PEtherToken", "err", err)
	}
	if _, err := etherTokenInChildChain.SwapFromEth(opt4); err != nil {
		log.Error("Failed to swap PEtherToken", "err", err)
	}

	ev = <-events.Chan()
	blockInfo = ev.Data.(core.NewMinedBlockEvent)

	if rcm.minerEnv.IsRequest {
		t.Fatal("Block should not be request block", "blockNumber", blockInfo.Block.NumberU64())
	}

	// NRE#5 - rest blocks
	for i = 0; i < NRELength.Uint64()-1; i++ {
		makeSampleTx(rcm)
		ev := <-events.Chan()
		blockInfo := ev.Data.(core.NewMinedBlockEvent)

		if rcm.minerEnv.IsRequest {
			t.Fatal("Block should not be request block", "blockNumber", blockInfo.Block.NumberU64())
		}
	}

	// #6 empty ORE

	// #7 NRE
	for i = 0; i < NRELength.Uint64(); i++ {
		makeSampleTx(rcm)
		ev := <-events.Chan()
		blockInfo := ev.Data.(core.NewMinedBlockEvent)

		if rcm.minerEnv.IsRequest {
			t.Fatal("Block should not be request block", "blockNumber", blockInfo.Block.NumberU64())
		}
	}

	// #8 ORE
	ev = <-events.Chan()
	blockInfo = ev.Data.(core.NewMinedBlockEvent)
	if !rcm.minerEnv.IsRequest {
		t.Fatal("Block should be request block", "blockNumber", blockInfo.Block.NumberU64())
	}

	finalizeBlocks(t, pls.rootchainManager.rootchainContract, 10)

	// wait challenge period ends
	wait(20)

	applyRequests(t, rcm.rootchainContract, key2)

	balances3 := getEtherTokenBalances(addrs)

	for i, balance2 := range balances2 {
		balance3 := balances3[i]
		if err := checkBalance(balance2, balance3, withdrawalAmount, nil, "check exit result"); err != nil {
			t.Fatal(err)
		}
	}

	log.Info("test finished")
	return
}

// TestScenario3 tests enter & exit with token transfer in child chain.
func TestScenario3(t *testing.T) {
	pls, rpcServer, dir, err := makePls()
	defer os.RemoveAll(dir)

	if err != nil {
		t.Fatalf("Failed to make pls service: %v", err)
	}
	defer pls.Stop()
	defer rpcServer.Stop()

	if err := pls.rootchainManager.Start(); err != nil {
		t.Fatalf("Failed to start RootChainManager: %v", err)
	}
	pls.protocolManager.Start(1)

	rpcClient := rpc.DialInProc(rpcServer)

	// assign to global variable
	plsClient = plsclient.NewClient(rpcClient)

	plasmaBlockMinedEvents := pls.rootchainManager.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer plasmaBlockMinedEvents.Unsubscribe()

	blockSubmitEvents := make(chan *rootchain.RootChainBlockSubmitted)
	blockSubmitWatchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}
	blockFilterer, _ := pls.rootchainManager.rootchainContract.WatchBlockSubmitted(blockSubmitWatchOpts, blockSubmitEvents)
	defer blockFilterer.Unsubscribe()

	var tx *types.Transaction

	wait(3)

	log.Info("All backends are set up")

	// NRE#1 / Block#1 (1/2)
	tokenInRootChain, tokenInChildChain, tokenAddrInRootChain, tokenAddrInChildChain := deployTokenContracts(t)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 1); err != nil {
		t.Fatal(err)
	}

	setNonce(operatorOpt, &operatorNonceRootChain)
	tx, err = tokenInRootChain.Mint(operatorOpt, addr1, ether(100))
	if err != nil {
		t.Fatalf("Failed to mint token: %v", err)
	}

	waitTx(tx.Hash())

	ts1, err := tokenInRootChain.TotalSupply(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get total supply from root chain: %v", err)
	}

	ts2, err := tokenInChildChain.TotalSupply(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get total supply from child chain: %v", err)
	}

	log.Info("Token total supply", "rootchain", ts1, "childchain", ts2)

	setNonce(operatorOpt, &operatorNonceRootChain)
	tx, err = pls.rootchainManager.rootchainContract.MapRequestableContractByOperator(operatorOpt, tokenAddrInRootChain, tokenAddrInChildChain)
	if err != nil {
		t.Fatalf("Failed to map token addresses to RootChain contract: %v", err)
	}

	waitTx(tx.Hash())

	tokenAddr, err := pls.rootchainManager.rootchainContract.RequestableContracts(baseCallOpt, tokenAddrInRootChain)
	if err != nil {
		t.Fatalf("Failed to fetch token address from RootChain contract: %v", err)
	} else if tokenAddr != tokenAddrInChildChain {
		t.Fatalf("RootChain doesn't know requestable contract address in child chain: %v != %v", tokenAddrInChildChain.Hex(), tokenAddr.Hex())
	}

	// NRE#1 -> ORE#4 -- deposit 1 ether for each account
	ETHBalances1 := getETHBalances(addrs)
	PETHBalances1 := getPETHBalances(addrs)

	depositAmount := ether(1)
	depositAmountNeg := new(big.Int).Neg(depositAmount)

	for _, key := range keys {
		startETHDeposit(t, pls.rootchainManager, key, depositAmount)
	}

	// NRE#1 / Block#2 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 2); err != nil {
		t.Fatal(err)
	}

	// ORE#2 is empty

	// NRE#3 -> ORE#6 -- deposit 1 token from addr1
	tokenAmount := ether(1)
	tokenAmountNeg := new(big.Int).Neg(tokenAmount)

	TokenBalances1 := getTokenBalances(addrs, tokenInRootChain)
	PTokenBalances1 := getTokenBalances(addrs, tokenInChildChain)

	startTokenDeposit(t, pls.rootchainManager, tokenInRootChain, tokenAddrInRootChain, key1, tokenAmount)

	data1, err := pls.rootchainManager.rootchainContract.GetEROTxData(baseCallOpt, big.NewInt(4))
	if err != nil {
		t.Fatalf("Failed to get ero tx data: %v", err)
	}

	// NRE#3 / Block#3 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 3); err != nil {
		t.Fatal(err)
	}

	// NRE#3 / Block#4 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 4); err != nil {
		t.Fatal(err)
	}

	// ORE#4 / Block#5 (1/1): deposit 1 ether for each account
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, true, 0, 5); err != nil {
		t.Fatal(err)
	}

	ETHBalances2 := getETHBalances(addrs)
	PETHBalances2 := getPETHBalances(addrs)

	// check eth balance in root chain
	for i := range keys {
		if err := checkBalance(ETHBalances1[i], ETHBalances2[i], depositAmountNeg, maxTxFee, "Failed to check ETH balance(1)"); err != nil {
			t.Fatal(err)
		}
	}

	// check peth balance in child chain
	for i := range keys {
		if err := checkBalance(PETHBalances1[i], PETHBalances2[i], depositAmount, nil, "Failed to check PETH balance(1)"); err != nil {
			t.Fatal(err)
		}
	}

	// NRE#5 / Block#6 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 6); err != nil {
		t.Fatal(err)
	}

	// NRE#5 / Block#7 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 7); err != nil {
		t.Fatal(err)
	}

	// ORE#6/ Block#8 (1/1) -- deposit 1 token from addr1
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, true, 0, 8); err != nil {
		t.Fatal(err)
	}

	TokenBalances2 := getTokenBalances(addrs, tokenInRootChain)
	PTokenBalances2 := getTokenBalances(addrs, tokenInChildChain)

	// check Token balance
	if err := checkBalance(TokenBalances1[0], TokenBalances2[0], tokenAmountNeg, nil, "Failed to check Token Balance (1)"); err != nil {
		t.Fatal(err)
	}

	// check PToken balance
	if err := checkBalance(PTokenBalances1[0], PTokenBalances2[0], tokenAmount, nil, "Failed to check PToken Balance (1)"); err != nil {
		t.Fatal(err)
	}

	// transfer token from addr1 to addr2, in child chain
	tokenAmountToTransfer := new(big.Int).Div(ether(1), big.NewInt(2))
	tokenAmountToTransferNeg := new(big.Int).Neg(tokenAmountToTransfer)

	// NRE#7 / Block#9 (1/2)
	transferToken(t, tokenInChildChain, key1, addr2, tokenAmountToTransfer, false)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 9); err != nil {
		t.Fatal(err)
	}

	PTokenBalances3 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances2[0], PTokenBalances3[0], tokenAmountToTransferNeg, nil, "Failed to check PToken Balance (2-1)"); err != nil {
		t.Fatal(err)
	}
	if err := checkBalance(PTokenBalances2[1], PTokenBalances3[1], tokenAmountToTransfer, nil, "Failed to check PToken Balance (2-2)"); err != nil {
		t.Fatal(err)
	}

	// NRE#7 / Block#10 (2/2)
	transferToken(t, tokenInChildChain, key1, addr2, tokenAmountToTransfer, false)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 10); err != nil {
		t.Fatal(err)
	}

	PTokenBalances4 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances3[0], PTokenBalances4[0], tokenAmountToTransferNeg, nil, "Failed to check PToken Balance (3-1)"); err != nil {
		t.Fatal(err)
	}
	if err := checkBalance(PTokenBalances3[1], PTokenBalances4[1], tokenAmountToTransfer, nil, "Failed to check PToken Balance (3-2)"); err != nil {
		t.Fatal(err)
	}

	// ORE#8 is empty

	// NRE#9 / Block#11 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 11); err != nil {
		t.Fatal(err)
	}

	// NRE#9 -> ORE#12 -- (1/4) withdraw addr2's token to root chain
	tokenAmountToWithdraw := new(big.Int).Div(tokenAmount, big.NewInt(4)) // 4 witndrawal requests
	tokenAmountToWithdrawNeg := new(big.Int).Neg(tokenAmountToWithdraw)

	PTokenBalances5 := getTokenBalances(addrs, tokenInChildChain)
	startTokenWithdraw(t, pls.rootchainManager.rootchainContract, tokenInRootChain, tokenAddrInRootChain, key2, tokenAmountToWithdraw, big.NewInt(int64(pls.rootchainManager.state.costERO)))

	// NRE#9 / Block#12 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 12); err != nil {
		t.Fatal(err)
	}

	// ORE#10 is empty

	// NRE#11 / Block#13 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 13); err != nil {
		t.Fatal(err)
	}

	// NRE#11 / Block#14 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 14); err != nil {
		t.Fatal(err)
	}

	// ORE#12 / Block#15 (1/1) -- (1/4) withdraw addr2's token to root chain
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, true, 0, 15); err != nil {
		t.Fatal(err)
	}

	PTokenBalances6 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances5[1], PTokenBalances6[1], tokenAmountToWithdrawNeg, nil, "Failed to check PToken Balance - token exit (1/4)"); err != nil {
		t.Fatal(err)
	}

	// NRE#13/ Block#16 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 16); err != nil {
		t.Fatal(err)
	}

	// NRE#13 -> ORE#16 -- (2/4) withdraw addr2's token to root chain
	startTokenWithdraw(t, pls.rootchainManager.rootchainContract, tokenInRootChain, tokenAddrInRootChain, key2, tokenAmountToWithdraw, big.NewInt(int64(pls.rootchainManager.state.costERO)))

	// NRE#13/ Block#17 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 17); err != nil {
		t.Fatal(err)
	}

	// ORE#14 is empty

	// NRE#15/ Block#18 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 18); err != nil {
		t.Fatal(err)
	}

	// NRE#15 -> ORE#18 -- (3/4) withdraw addr2's token to root chain
	startTokenWithdraw(t, pls.rootchainManager.rootchainContract, tokenInRootChain, tokenAddrInRootChain, key2, tokenAmountToWithdraw, big.NewInt(int64(pls.rootchainManager.state.costERO)))

	// NRE#15/ Block#19 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 19); err != nil {
		t.Fatal(err)
	}

	// ORE#16 / Block#20 -- (2/4) withdraw addr2's token to root chain
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, true, 0, 20); err != nil {
		t.Fatal(err)
	}

	PTokenBalances7 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances6[1], PTokenBalances7[1], tokenAmountToWithdrawNeg, nil, "Failed to check PToken Balance - token exit (2/4)"); err != nil {
		t.Fatal(err)
	}

	// NRE#17/ Block#21 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 21); err != nil {
		t.Fatal(err)
	}

	// NRE#17 -> ORE#20 -- (4/4) withdraw addr2's token to root chain
	startTokenWithdraw(t, pls.rootchainManager.rootchainContract, tokenInRootChain, tokenAddrInRootChain, key2, tokenAmountToWithdraw, big.NewInt(int64(pls.rootchainManager.state.costERO)))

	// NRE#17/ Block#22 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 22); err != nil {
		t.Fatal(err)
	}

	// ORE#18 / Block#23 -- (3/4) withdraw addr2's token to root chain
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, true, 0, 23); err != nil {
		t.Fatal(err)
	}

	PTokenBalances8 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances7[1], PTokenBalances8[1], tokenAmountToWithdrawNeg, nil, "Failed to check PToken Balance - token exit (3/4)"); err != nil {
		t.Fatal(err)
	}

	// NRE#19/ Block#24 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 24); err != nil {
		t.Fatal(err)
	}

	// NRE#19/ Block#25 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 25); err != nil {
		t.Fatal(err)
	}

	// ORE#20 / Block#26 -- (4/4) withdraw addr2's token to root chain
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, true, 0, 26); err != nil {
		t.Fatal(err)
	}

	PTokenBalances9 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances8[1], PTokenBalances9[1], tokenAmountToWithdrawNeg, nil, "Failed to check PToken Balance - token exit (4/4)"); err != nil {
		t.Fatal(err)
	}

	// finalize until block#26
	finalizeBlocks(t, pls.rootchainManager.rootchainContract, 26)

	// apply requests (4 ETH deposits, 1 Token deposits, 4 Token withdrawals)
	for i := 0; i < 4+1; i++ {
		applyRequest(t, pls.rootchainManager.rootchainContract, operatorKey)
	}
	for i := 0; i < 4; i++ {
		tokenBalanceBefore, _ := tokenInRootChain.Balances(baseCallOpt, addr2)
		applyRequest(t, pls.rootchainManager.rootchainContract, operatorKey)
		tokenBalanceAfter, _ := tokenInRootChain.Balances(baseCallOpt, addr2)

		if tokenBalanceAfter.Cmp(new(big.Int).Add(tokenBalanceBefore, tokenAmountToWithdraw)) != 0 {
			t.Fatalf("applyRequest() does not increase token balance")
		}

		log.Info("Exit request applied")
	}

	t.Log("Test finished")
}

// test challenge invalid exit
func TestScenario4(t *testing.T) {
	pls, rpcServer, dir, err := makePls()
	defer os.RemoveAll(dir)

	if err != nil {
		t.Fatalf("Failed to make pls service: %v", err)
	}
	defer pls.Stop()
	defer rpcServer.Stop()

	// pls.Start()
	pls.protocolManager.Start(1)

	if err := pls.rootchainManager.Start(); err != nil {
		t.Fatalf("Failed to start RootChainManager: %v", err)
	}

	pls.StartMining(runtime.NumCPU())

	rpcClient := rpc.DialInProc(rpcServer)

	// assign to global variable
	plsClient = plsclient.NewClient(rpcClient)

	plasmaBlockMinedEvents := pls.rootchainManager.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer plasmaBlockMinedEvents.Unsubscribe()

	blockSubmitEvents := make(chan *rootchain.RootChainBlockSubmitted)
	blockSubmitWatchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}
	blockFilterer, _ := pls.rootchainManager.rootchainContract.WatchBlockSubmitted(blockSubmitWatchOpts, blockSubmitEvents)
	defer blockFilterer.Unsubscribe()

	wait(3)

	log.Info("All backends are set up")

	// NRE#1 / Block#1 (1/2)
	tokenInRootChain, tokenInChildChain, tokenAddrInRootChain, tokenAddrInChildChain := deployTokenContracts(t)

	wait(4)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 1); err != nil {
		t.Fatal(err)
	}

	opt := makeTxOpt(operatorKey, 0, nil, nil)

	_, err = tokenInRootChain.Mint(opt, addr1, ether(100))
	if err != nil {
		t.Fatalf("Failed to mint token: %v", err)
	}
	wait(2)

	ts1, err := tokenInRootChain.TotalSupply(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get total supply from root chain: %v", err)
	}

	ts2, err := tokenInChildChain.TotalSupply(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get total supply from child chain: %v", err)
	}

	log.Info("Token total supply", "rootchain", ts1, "childchain", ts2)

	wait(3)

	_, err = pls.rootchainManager.rootchainContract.MapRequestableContractByOperator(opt, tokenAddrInRootChain, tokenAddrInChildChain)
	if err != nil {
		t.Fatalf("Failed to map token addresses to RootChain contract: %v", err)
	}
	wait(2)

	tokenAddr, err := pls.rootchainManager.rootchainContract.RequestableContracts(baseCallOpt, tokenAddrInRootChain)
	wait(2)
	if err != nil {
		t.Fatalf("Failed to fetch token address from RootChain contract: %v", err)
	} else if tokenAddr != tokenAddrInChildChain {
		t.Fatalf("RootChain doesn't know requestable contract address in child chain: %v != %v", tokenAddrInChildChain, tokenAddr)
	}

	// NRE#1 -> ORE#4 -- deposit 1 ether for each account
	ETHBalances1 := getETHBalances(addrs)
	PETHBalances1 := getPETHBalances(addrs)

	depositAmount := ether(10)
	depositAmountNeg := new(big.Int).Neg(depositAmount)

	for _, key := range keys {
		startETHDeposit(t, pls.rootchainManager, key, depositAmount)
	}

	// NRE#1 / Block#2 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 2); err != nil {
		t.Fatal(err)
	}

	// ORB#2 is empty

	// NRE#3 -> ORE#6 -- deposit 1 token from addr1
	tokenAmount := ether(1)
	tokenAmountNeg := new(big.Int).Neg(tokenAmount)

	TokenBalances1 := getTokenBalances(addrs, tokenInRootChain)
	PTokenBalances1 := getTokenBalances(addrs, tokenInChildChain)

	startTokenDeposit(t, pls.rootchainManager, tokenInRootChain, tokenAddrInRootChain, key1, tokenAmount)

	// NRE#3 / Block#3 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 3); err != nil {
		t.Fatal(err)
	}

	// NRE#3 / Block#4 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 4); err != nil {
		t.Fatal(err)
	}

	// ORE#4 / Block#5 (1/1) -- deposit 1 ether for each account
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, true, 0, 5); err != nil {
		t.Fatal(err)
	}

	ETHBalances2 := getETHBalances(addrs)
	PETHBalances2 := getPETHBalances(addrs)
	// check eth balance in root chain
	for i := range keys {
		if err := checkBalance(ETHBalances1[i], ETHBalances2[i], depositAmountNeg, maxTxFee, "Failed to check ETH balance(1)"); err != nil {
			t.Fatal(err)
		}
	}

	// check peth balance in child chain
	for i := range keys {
		if err := checkBalance(PETHBalances1[i], PETHBalances2[i], depositAmount, nil, "Failed to check ETH balance(1)"); err != nil {
			t.Fatal(err)
		}
	}

	// NRE#5 / Block#6 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 6); err != nil {
		t.Fatal(err)
	}

	// NRE#5 / Block#7 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 7); err != nil {
		t.Fatal(err)
	}

	// ORE#6 / Block#8 (1/1) -- deposit 1 token from addr1
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, true, 0, 8); err != nil {
		t.Fatal(err)
	}

	TokenBalances2 := getTokenBalances(addrs, tokenInRootChain)
	PTokenBalances2 := getTokenBalances(addrs, tokenInChildChain)

	// check Token balance
	if err := checkBalance(TokenBalances1[0], TokenBalances2[0], tokenAmountNeg, nil, "Failed to check Token Balance (1)"); err != nil {
		t.Fatal(err)
	}

	// check PToken balance
	if err := checkBalance(PTokenBalances1[0], PTokenBalances2[0], tokenAmount, nil, "Failed to check PToken Balance (1)"); err != nil {
		t.Fatal(err)
	}

	// NRE#7 -> ORB#9 -- invalid withdrawal
	startTokenWithdraw(t, pls.rootchainManager.rootchainContract, tokenInRootChain, tokenAddrInRootChain, key2, ether(100), big.NewInt(int64(pls.rootchainManager.state.costERO)))

	// NRE#7 / Block#9 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 9); err != nil {
		t.Fatal(err)
	}

	// NRE#7 / Block#10 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 10); err != nil {
		t.Fatal(err)
	}

	// ORB#8 is empty

	// NRE#9 / Block#11 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 11); err != nil {
		t.Fatal(err)
	}

	// NRE#9 / Block#12 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 12); err != nil {
		t.Fatal(err)
	}

	// ORE#10 / Block#13 (1/1) -- invalid withdrawal
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, true, 0, 13); err != nil {
		t.Fatal(err)
	}

	// NRE#11 / Block#14 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 14); err != nil {
		t.Fatal(err)
	}

	// NRE#11 / Block#15 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 15); err != nil {
		t.Fatal(err)
	}

	// ORB#12 is empty

	// NRE#13 / Block#16 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 16); err != nil {
		t.Fatal(err)
	}

	// NRE#13 / Block#17 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 17); err != nil {
		t.Fatal(err)
	}

	// ORB#14 is empty

	// NRE#15 / Block#18 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 18); err != nil {
		t.Fatal(err)
	}

	// NRE#15 / Block#19 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 19); err != nil {
		t.Fatal(err)
	}

	wait(10)

	ERO, err := pls.rootchainManager.rootchainContract.EROs(baseCallOpt, big.NewInt(5))
	if err != nil {
		t.Fatal("failed to get ERO")
	}

	if !ERO.Challenged {
		t.Fatal("ERO is not challenged successfully")
	}
}

func TestStress(t *testing.T) {

	testPlsConfig.MinerRecommit = 10 * time.Second
	testPlsConfig.TxConfig.Interval = 10 * time.Second
	timeout := testPlsConfig.MinerRecommit / 100
	targetBlockNumber := 10

	pls, rpcServer, dir, err := makePls()
	defer os.RemoveAll(dir)

	if err != nil {
		t.Fatalf("Failed to make pls service: %v", err)
	}
	defer pls.Stop()
	defer rpcServer.Stop()

	if err := pls.rootchainManager.Start(); err != nil {
		t.Fatalf("Failed to start RootChainManager: %v", err)
	}
	pls.protocolManager.Start(1)

	rpcClient := rpc.DialInProc(rpcServer)

	// assign to global variable
	plsClient = plsclient.NewClient(rpcClient)

	plasmaBlockMinedEvents := pls.rootchainManager.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer plasmaBlockMinedEvents.Unsubscribe()

	blockSubmitEvents := make(chan *rootchain.RootChainBlockSubmitted)
	blockSubmitWatchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}
	blockFilterer, _ := pls.rootchainManager.rootchainContract.WatchBlockSubmitted(blockSubmitWatchOpts, blockSubmitEvents)
	defer blockFilterer.Unsubscribe()

	blockNumber := 0

	timer := time.NewTimer(5 * time.Minute)
	go func() {
		<-timer.C
		t.Fatal("Out of time")
	}()

	txs := types.Transactions{}
	nTxsInBlocks := 0

	blockNumber++

	for _, addr := range addrs {
		var tx *types.Transaction
		if tx, err = transferETH(operatorKey, addr, ether(1), false); err != nil {
			t.Fatalf("Failed to transfer PETH: %v", err)
		}
		txs = append(txs, tx)
	}
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, int64(blockNumber)); err != nil {
		t.Fatal(err)
	}

	for blockNumber < targetBlockNumber {
		blockNumber++

		var tx *types.Transaction
		for _, addr := range addrs {
			if tx, err = transferETH(operatorKey, addr, ether(1), false); err != nil {
				t.Fatalf("Failed to transfer PETH: %v", err)
			}
		}

		txs = append(txs, tx)

		done := make(chan struct{})
		wg := sync.WaitGroup{}

		wg.Add(1)
		go func(t *testing.T) {
			for {
				timer := time.NewTimer(timeout)
				select {
				case <-timer.C:
					timer.Reset(timeout)

					for i, addr := range addrs {
						go func(t *testing.T, i int, addr common.Address) {
							var tx *types.Transaction

							if tx, err = transferETH(operatorKey, addr, ether(1), false); err != nil {
								t.Fatalf("Failed to transfer PETH: %v", err)
							}
							txs = append(txs, tx)

							key := keys[i]

							if tx, err = transferETH(key, addr, ether(0.0001), false); err != nil {
								t.Fatalf("Failed to transfer PETH: %v", err)
							}
							txs = append(txs, tx)
						}(t, i, addr)
					}

				case <-done:
					wg.Done()
					return
				}
			}
		}(t)

		// check new block is mined
		wg.Add(1)
		go func() {
			if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, int64(blockNumber)); err != nil {
				t.Fatal(err)
			}
			close(done)
			b := pls.blockchain.GetBlockByNumber(uint64(blockNumber))
			nTxsInBlocks += len(b.Transactions())
			wg.Done()
		}()

		wg.Wait()
	}

	wait(testPlsConfig.MinerRecommit * 2 / time.Second)

	for _, tx := range txs {
		r, isPending, err := plsClient.TransactionByHash(context.Background(), tx.Hash())
		signer := types.NewEIP155Signer(params.PlasmaChainConfig.ChainID)
		msg, _ := r.AsMessage(signer)
		from := msg.From()

		if isPending {
			t.Fatalf("Transaction %s is pending (from: %s, nonce: %d)", r.Hash().String(), from.String(), r.Nonce())
		}

		if err != nil {
			t.Fatalf("failed to get transaction receipt: %v", err)
		}

		if r == nil {
			t.Fatalf("Transaction %s not mined (from: %s, nonce: %d)", r.Hash().String(), from.String(), r.Nonce())
		}
	}

	nTxs := 0
	lastBlockNumber := pls.blockchain.CurrentBlock().NumberU64()
	for i := 1; i <= int(lastBlockNumber); i++ {
		block := pls.blockchain.GetBlockByNumber(uint64(i))

		if block.Transactions().Len() == 0 {
			t.Fatalf("Block#%d has no transaction", i)
		}

		nTxs += block.Transactions().Len()
	}

	firstBlock := pls.blockchain.GetBlockByNumber(uint64(1))
	lastBlock := pls.blockchain.GetBlockByNumber(uint64(lastBlockNumber))

	elapsed := new(big.Int).Sub(lastBlock.Time(), firstBlock.Time())
	tps := float64(nTxs) / float64(elapsed.Int64())
	t.Logf("Elapsed time: %s nTxs: %d TPS: %6.3f", elapsed.String(), nTxs, tps)

}

//func TestAdjustGasPrice(t *testing.T) {
//	quit := make(chan bool, 1)
//	pls, rpcServer, dir, err := makePls()
//	if err != nil {
//		t.Fatalf("Failed to make pls service: %v", err)
//	}
//	wait(3)
//
//	defer os.RemoveAll(dir)
//	defer pls.Stop()
//	defer rpcServer.Stop()
//	defer func() {
//		quit <- true
//	}()
//
//	originalGasPrice := big.NewInt(1 * params.GWei)
//	newGasPrice := big.NewInt(1 * params.GWei)
//
//	pls.rootchainManager.state.gasPrice = new(big.Int).Set(originalGasPrice)
//	pls.rootchainManager.config.TxConfig.MaxGasPrice = big.NewInt(100 * params.GWei)
//	pls.config.TxConfig.Interval = 300 * time.Millisecond
//
//	go func() {
//		nonce, _ := ethClient.NonceAt(context.Background(), addr1, nil)
//		opt1.GasPrice = big.NewInt(2 * params.GWei)
//		for {
//			select {
//			case <-quit:
//				return
//			default:
//				opt1.Nonce = big.NewInt(int64(nonce))
//				_, _, _, err := epochhandler.DeployEpochHandler(opt1, ethClient)
//				if err != nil {
//					nonce++
//				}
//				nonce++
//			}
//		}
//	}()
//
//	pls.protocolManager.Start(1)
//
//	if err := pls.rootchainManager.Start(); err != nil {
//		t.Fatalf("Failed to start RootChainManager: %v", err)
//	}
//
//	pls.StartMining(runtime.NumCPU())
//
//	// assign to global variable
//	rpcClient := rpc.DialInProc(rpcServer)
//	plsClient = plsclient.NewClient(rpcClient)
//
//	//plasmaBlockMinedEvents := pls.rootchainManager.eventMux.Subscribe(core.NewMinedBlockEvent{})
//	//defer plasmaBlockMinedEvents.Unsubscribe()
//
//	blockSubmitEvents := make(chan *rootchain.RootChainBlockSubmitted)
//	blockSubmitWatchOpts := &bind.WatchOpts{
//		Start:   nil,
//		Context: context.Background(),
//	}
//	blockFilterer, _ := pls.rootchainManager.rootchainContract.WatchBlockSubmitted(blockSubmitWatchOpts, blockSubmitEvents)
//	defer blockFilterer.Unsubscribe()
//
//	log.Info("All backends are set up")
//
//	timerInterval := 20 * time.Second
//	timer := time.NewTimer(timerInterval)
//
//	for i := 0; i < 10; i++ {
//		makeSampleTx(pls.rootchainManager)
//		//<-plasmaBlockMinedEvents.Chan()
//
//		select {
//		case <-blockSubmitEvents:
//			timer.Reset(timerInterval)
//		case _, ok := <-timer.C:
//			if ok {
//				t.Fatal("out of time")
//			}
//		}
//
//		originalGasPrice = new(big.Int).Set(newGasPrice)
//		newGasPrice = new(big.Int).Set(pls.rootchainManager.state.gasPrice)
//
//		if originalGasPrice.Cmp(newGasPrice) == 0 {
//			t.Fatalf("originalGasPrice: %v, new: %v", originalGasPrice, newGasPrice)
//		}
//	}
//}

func TestRestart(t *testing.T) {
	timer := time.NewTimer(2 * time.Minute)
	go func() {
		<-timer.C
		t.Fatal("Out of time")
	}()

	pls, rpcServer, dir, err := makePls()
	defer os.RemoveAll(dir)

	if err != nil {
		t.Fatalf("Failed to make pls service: %v", err)
	}
	defer pls.Stop()
	defer rpcServer.Stop()

	if err := pls.rootchainManager.Start(); err != nil {
		t.Fatalf("Failed to start RootChainManager: %v", err)
	}
	pls.protocolManager.Start(1)

	rpcClient := rpc.DialInProc(rpcServer)

	// assign to global variable
	plsClient = plsclient.NewClient(rpcClient)

	plasmaBlockMinedEvents := pls.rootchainManager.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer plasmaBlockMinedEvents.Unsubscribe()

	blockSubmitEvents := make(chan *rootchain.RootChainBlockSubmitted)
	blockSubmitWatchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}
	blockFilterer, _ := pls.rootchainManager.rootchainContract.WatchBlockSubmitted(blockSubmitWatchOpts, blockSubmitEvents)
	defer blockFilterer.Unsubscribe()

	rcm := pls.rootchainManager

	wait(3)

	log.Info("All backends are set up")

	// ORE#4 make enter request
	enterAmount := ether(1)

	startETHDeposit(t, rcm, key1, enterAmount)
	startETHDeposit(t, rcm, key2, enterAmount)
	startETHDeposit(t, rcm, key3, enterAmount)
	startETHDeposit(t, rcm, key4, enterAmount)

	// NRE#1 / NRB#1
	// deploy EtherToken in child chain
	deployEtherTokenInChildChain(t)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 1); err != nil {
		t.Fatal(err)
	}
	// NRE#1 / NRB#2
	//tokenInRootChain, tokenInChildChain, tokenAddrInRootChain, tokenAddrInChildChain := deployTokenContracts(t)
	deployTokenContracts(t)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 2); err != nil {
		t.Fatal(err)
	}

	// ORE#2 is empty

	// NRE#3 / NRB#3
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 3); err != nil {
		t.Fatal(err)
	}
	// NRE#3 / NRB#4
	makeSampleTx(pls.rootchainManager)
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 4); err != nil {
		t.Fatal(err)
	}

	// ORE#4 / ORB#5 : enter request
	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, true, 0, 5); err != nil {
		t.Fatal(err)
	}

	// stop pls service
}

func TestMinerRestart(t *testing.T) {
	pls, rpcServer, dir, err := makePls()
	defer os.RemoveAll(dir)

	if err != nil {
		t.Fatalf("Failed to make pls service: %v", err)
	}
	defer pls.Stop()
	defer rpcServer.Stop()

	// pls.Start()
	pls.protocolManager.Start(1)

	if err := pls.rootchainManager.Start(); err != nil {
		t.Fatalf("Failed to start RootChainManager: %v", err)
	}

	pls.StartMining(runtime.NumCPU())

	rpcClient := rpc.DialInProc(rpcServer)

	// assign to global variable
	plsClient = plsclient.NewClient(rpcClient)

	plasmaBlockMinedEvents := pls.rootchainManager.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer plasmaBlockMinedEvents.Unsubscribe()

	blockSubmitEvents := make(chan *rootchain.RootChainBlockSubmitted)
	blockSubmitWatchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}
	blockFilterer, _ := pls.rootchainManager.rootchainContract.WatchBlockSubmitted(blockSubmitWatchOpts, blockSubmitEvents)
	defer blockFilterer.Unsubscribe()

	log.Info("All backends are set up")

	// NRE#1 / Block#1 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 1); err != nil {
		t.Fatal(err)
	}

	pls.StopMining()
	blockBeforeStop := pls.blockchain.CurrentBlock()

	wait(5)
	pls.StartMining(runtime.NumCPU())

	// NRE#1 / Block#2 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 2); err != nil {
		t.Fatal(err)
	}

	blockAfterStop := pls.blockchain.CurrentBlock()
	if diff := blockAfterStop.NumberU64() - blockBeforeStop.NumberU64(); diff != 1 {
		t.Fatal("failed to resume current epoch", "difference", diff)
	}
}

func startETHDeposit(t *testing.T, rcm *RootChainManager, key *ecdsa.PrivateKey, amount *big.Int) {
	if amount.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("Cannot deposit 0 ETH")
	}

	watchOpt := &bind.WatchOpts{Start: nil, Context: context.Background()}
	event := make(chan *rootchain.RootChainRequestCreated)
	filterer, _ := rcm.rootchainContract.WatchRequestCreated(watchOpt, event)
	defer filterer.Unsubscribe()

	addr := crypto.PubkeyToAddress(key.PublicKey)
	opt := opts[addr]
	setNonce(opt, noncesRootChain[addr])
	opt.Nonce = nil

	trieKey, err := etherToken.GetBalanceTrieKey(baseCallOpt, addr)
	if err != nil {
		t.Fatalf("Failed to get trie key: %v", err)
	}
	trieValue := amount.Bytes()
	trieValue32Bytes := common.BytesToHash(trieValue)

	tx, err := rcm.rootchainContract.StartEnter(opt, etherTokenAddr, trieKey, trieValue32Bytes[:])

	if err != nil {
		t.Fatalf("Failed to make an ETH deposit request: %v", err)
	}

	if err = waitTx(tx.Hash()); err != nil {
		t.Fatalf("Failed to make an ETH deposit request: %v", err)
	}

	receipt, err := rcm.backend.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("Failed to send eth deposit tx: %v", err)
	} else if receipt.Status == 0 {
		t.Fatal("ETH deposit tx is reverted")
	}

	<-event
}

func startTokenDeposit(t *testing.T, rcm *RootChainManager, tokenContract *token.RequestableSimpleToken, tokenAddress common.Address, key *ecdsa.PrivateKey, amount *big.Int) {
	if amount.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("Cannot deposit 0 Token")
	}

	watchOpt := &bind.WatchOpts{Start: nil, Context: context.Background()}
	event := make(chan *rootchain.RootChainRequestCreated)
	filterer, _ := rcm.rootchainContract.WatchRequestCreated(watchOpt, event)
	defer filterer.Unsubscribe()

	addr := crypto.PubkeyToAddress(key.PublicKey)
	opt := opts[addr]
	setNonce(opt, noncesRootChain[addr])

	trieKey, err := tokenContract.GetBalanceTrieKey(baseCallOpt, addr)
	if err != nil {
		t.Fatalf("Failed to get trie key: %v", err)
	}
	trieValue := amount.Bytes()
	trieValue = common.LeftPadBytes(trieValue, 32)
	trieValue32Bytes := common.BytesToHash(trieValue)

	tx, err := rcm.rootchainContract.StartEnter(opt, tokenAddress, trieKey, trieValue32Bytes[:])

	if err != nil {
		log.Error("Failed to make an token deposit request", "err", err, "hash", tx.Hash())
		t.Fatalf("Failed to make an token deposit request: %v", err)
	}

	waitTx(tx.Hash())

	request := <-event
	log.Debug("Token deposit request", "request", request)

	receipt, err := rcm.backend.TransactionReceipt(context.Background(), tx.Hash())

	log.Debug("Token deposit", "hash", tx.Hash().String(), "receipt", receipt)

	if err != nil {
		t.Fatalf("Failed to send token deposit tx: %v", err)
	} else if receipt.Status == 0 {
		t.Fatal("Token deposit tx is reverted")
	}
}

func startETHWithdraw(t *testing.T, rcm *RootChainManager, key *ecdsa.PrivateKey, value, cost *big.Int) {
	if value.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("Cannot deposit 0 ETH")
	}

	watchOpt := &bind.WatchOpts{Start: nil, Context: context.Background()}
	event := make(chan *rootchain.RootChainRequestCreated)
	filterer, _ := rcm.rootchainContract.WatchRequestCreated(watchOpt, event)
	defer filterer.Unsubscribe()

	addr := crypto.PubkeyToAddress(key.PublicKey)
	opt := makeTxOpt(key, 2000000, nil, cost)
	setNonce(opt, noncesRootChain[addr])
	opt.Nonce = nil

	trieKey, err := etherToken.GetBalanceTrieKey(baseCallOpt, addr)
	if err != nil {
		t.Fatalf("Failed to get trie key: %v", err)
	}
	trieValue := value.Bytes()
	trieValue32Bytes := common.BytesToHash(trieValue)

	tx, err := rcm.rootchainContract.StartExit(opt, etherTokenAddr, trieKey, trieValue32Bytes[:])

	if err != nil {
		t.Fatalf("Failed to make an ETH withdraw request: %v", err)
	}

	waitTx(tx.Hash())

	receipt, err := rcm.backend.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("Failed to send eth deposit tx: %v", err)
	} else if receipt.Status == 0 {
		t.Fatal("ETH withdraw tx is reverted")
	}

	<-event
}

func startTokenWithdraw(t *testing.T, rootchainContract *rootchain.RootChain, tokenContract *token.RequestableSimpleToken, tokenAddress common.Address, key *ecdsa.PrivateKey, amount, cost *big.Int) {
	opt := makeTxOpt(key, 0, nil, cost)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	setNonce(opt, noncesRootChain[addr])

	trieKey, err := tokenContract.GetBalanceTrieKey(baseCallOpt, addr)
	if err != nil {
		t.Fatalf("Failed to get trie key: %v", err)
	}
	trieValue := amount.Bytes()
	trieValue = common.LeftPadBytes(trieValue, 32)
	trieValue32Bytes := common.BytesToHash(trieValue)

	tx, err := rootchainContract.StartExit(opt, tokenAddress, trieKey, trieValue32Bytes[:])

	if err != nil {
		t.Fatalf("Failed to make an token withdrawal request: %v", err)
	}

	if err := waitTx(tx.Hash()); err != nil {
		t.Fatalf("failed to make exit request for token withdrawal")
	}
}

func transferToken(t *testing.T, tokenContract *token.RequestableSimpleToken, key *ecdsa.PrivateKey, to common.Address, amount *big.Int, isRootChain bool) {
	opt := makeTxOpt(key, 0, nil, nil)
	addr := crypto.PubkeyToAddress(key.PublicKey)

	if isRootChain {
		setNonce(opt, noncesRootChain[addr])
	} else {
		setNonce(opt, noncesChildChain[addr])
	}

	_, err := tokenContract.Transfer(opt, to, amount)
	if err != nil {
		t.Fatalf("Failed to transfer toekn: %v", err)
	}
}

func transferETH(key *ecdsa.PrivateKey, to common.Address, amount *big.Int, isRootChain bool) (*types.Transaction, error) {
	opt := makeTxOpt(key, 21000, defaultGasPrice, amount)
	addr := crypto.PubkeyToAddress(key.PublicKey)

	locks[addr].Lock()
	defer locks[addr].Unlock()

	if isRootChain {
		setNonce(opt, noncesRootChain[addr])
	} else {
		setNonce(opt, noncesChildChain[addr])
	}

	tx := types.NewTransaction(opt.Nonce.Uint64(), to, amount, 21000, defaultGasPrice, []byte{})

	var err error

	chainId := params.PlasmaChainConfig.ChainID

	if isRootChain {
		chainId = testPlsConfig.TxConfig.ChainId
	}

	signer := types.NewEIP155Signer(chainId)
	signedTx, err := types.SignTx(tx, signer, key)
	if err != nil {
		return nil, err
	}

	if isRootChain {
		err = ethClient.SendTransaction(context.Background(), signedTx)
		return signedTx, err
	}

	err = plsClient.SendTransaction(context.Background(), signedTx)
	return signedTx, err
}

func finalizeBlocks(t *testing.T, rootchainContract *rootchain.RootChain, targetNumber int64) {
	target := big.NewInt(targetNumber)

	last, err := rootchainContract.GetLastFinalizedBlock(baseCallOpt, big.NewInt(0))
	if err != nil {
		t.Fatalf("Failed to GetLastFinalizedBlock: %v", err)
	}

	for last.Cmp(target) < 0 {
		opt := opts[addr1]
		opt.Value = nil
		setNonce(opt, noncesRootChain[addr1])
		opt.GasLimit = 6000000

		log.Info("Try to finalize block", "lastFinalizedBlock", last, "lastBlock", target)

		tx, err := rootchainContract.FinalizeBlock(opt)
		if err != nil {
			t.Errorf("Failed to fianlize block: %v", err)
		}

		waitTx(tx.Hash())

		receipt, _ := ethClient.TransactionReceipt(context.Background(), tx.Hash())
		if receipt.Status == 0 {
			log.Error("FinalizeBlock transaction is failed")
		}

		last, err = rootchainContract.GetLastFinalizedBlock(baseCallOpt, big.NewInt(0))
		if err != nil {
			t.Fatalf("Failed to GetLastFinalizedBlock: %v", err)
		}

		wait(3)
	}

	log.Info("All blocks are fianlized")
}

// apply a single request
func applyRequest(t *testing.T, rootchainContract *rootchain.RootChain, key *ecdsa.PrivateKey) {
	opt := makeTxOpt(key, 2000000, nil, nil)
	addr := crypto.PubkeyToAddress(key.PublicKey)

	setNonce(opt, noncesRootChain[addr])

	wait(1)

	tx, err := rootchainContract.FinalizeRequest(opt)
	if err != nil {
		t.Fatalf("failed to apply requeest: %v", err)
	}

	if err := waitTx(tx.Hash()); err != nil {
		t.Fatalf("failed to apply requeest: %v", err)
	}
}

// apply all requests
func applyRequests(t *testing.T, rootchainContract *rootchain.RootChain, key *ecdsa.PrivateKey) {
	opt := makeTxOpt(key, 2000000, nil, nil)
	addr := crypto.PubkeyToAddress(key.PublicKey)

	last, err := rootchainContract.LastAppliedERO(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get last applied ERO: %v", err)
	}

	target, err := rootchainContract.GetNumEROs(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get number of EROs: %v", err)
	}

	target = new(big.Int).Sub(target, big.NewInt(1))

	for last.Cmp(target) < 0 {
		log.Info("Try to apply request", "last", last, "target", target)
		setNonce(opt, noncesRootChain[addr])

		wait(1)

		tx, err := rootchainContract.FinalizeRequest(opt)
		if err != nil {
			t.Fatalf("failed to apply requeest: %v", err)
		}

		if err := waitTx(tx.Hash()); err != nil {
			t.Fatalf("failed to apply requeest: %v", err)

		}

		last, _ = rootchainContract.LastAppliedERO(baseCallOpt)
		target, _ = rootchainContract.GetNumEROs(baseCallOpt)

	}
}

func deployRootChain(genesis *types.Block) (rootchainAddress common.Address, rootchainContract *rootchain.RootChain, err error) {

	dummyDB := ethdb.NewMemDatabase()
	defer dummyDB.Close()
	dummyBlock := core.DeveloperGenesisBlock(
		0,
		common.HexToAddress("0xdead"),
		operator,
		core.DefaultStaminaConfig,
	).ToBlock(dummyDB)

	var tx *types.Transaction
	log.Info("Deploying contracts for development mode")

	// 1. deploy MintableToken in root chain
	setNonce(operatorOpt, &operatorNonceRootChain)
	mintableTokenAddr, tx, mintableToken, err = mintabletoken.DeployMintableToken(operatorOpt, ethClient)

	if err != nil {
		return common.Address{}, nil, errors.New(fmt.Sprintf("Failed to deploy MintableToken contract: %v", err))
	}
	log.Info("Deploy MintableToken contract", "hash", tx.Hash(), "address", mintableTokenAddr)

	log.Info("Wait until deploy transaction is mined")
	waitTx(tx.Hash())

	// 2. deploy EtherToken in root chain
	setNonce(operatorOpt, &operatorNonceRootChain)
	etherTokenAddr, tx, etherToken, err = ethertoken.DeployEtherToken(operatorOpt, ethClient, development, mintableTokenAddr, swapEnabledInRootChain)

	if err != nil {
		return common.Address{}, nil, errors.New(fmt.Sprintf("Failed to deploy EtherToken contract: %v", err))
	}
	log.Info("Deploy EtherToken contract", "hash", tx.Hash(), "address", etherTokenAddr)

	log.Info("Wait until deploy transaction is mined")
	waitTx(tx.Hash())

	// 3. deploy EpochHandler in root chain
	setNonce(operatorOpt, &operatorNonceRootChain)
	epochHandlerAddr, tx, _, err := epochhandler.DeployEpochHandler(operatorOpt, ethClient)

	if err != nil {
		return common.Address{}, nil, errors.New(fmt.Sprintf("Failed to deploy EpochHandler contract: %v", err))
	}
	log.Info("Deploy EpochHandler contract", "hash", tx.Hash(), "address", epochHandlerAddr)

	log.Info("Wait until deploy transaction is mined")
	waitTx(tx.Hash())

	// 4. deploy RootChain in root chain
	setNonce(operatorOpt, &operatorNonceRootChain)
	rootchainAddr, tx, rootchainContract, err := rootchain.DeployRootChain(operatorOpt, ethClient, epochHandlerAddr, etherTokenAddr, development, NRELength, dummyBlock.Root(), dummyBlock.TxHash(), dummyBlock.ReceiptHash())
	if err != nil {
		return common.Address{}, nil, errors.New(fmt.Sprintf("Failed to deploy RootChain contract: %v", err))
	}
	log.Info("Deploy RootChain contract", "hash", tx.Hash(), "address", rootchainAddr)
	waitTx(tx.Hash())

	// 5. initialize EtherToken
	setNonce(operatorOpt, &operatorNonceRootChain)
	tx, err = etherToken.Init(operatorOpt, rootchainAddr)
	if err != nil {
		return common.Address{}, nil, errors.New(fmt.Sprintf("Failed to initialize EtherToken: %v", err))
	}
	log.Info("Initialize EtherToken", "hash", tx.Hash())
	waitTx(tx.Hash())

	// 6. mint tokens
	mintEvents := make(chan *mintabletoken.MintableTokenMint)
	mintWatchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}
	mintFilterrer, _ := mintableToken.WatchMint(mintWatchOpts, mintEvents, addrs)

	setNonce(operatorOpt, &operatorNonceRootChain)
	tx1, err := mintableToken.Mint(operatorOpt, addr1, ether(100))
	setNonce(operatorOpt, &operatorNonceRootChain)
	tx2, err := mintableToken.Mint(operatorOpt, addr2, ether(100))
	setNonce(operatorOpt, &operatorNonceRootChain)
	tx3, err := mintableToken.Mint(operatorOpt, addr3, ether(100))
	setNonce(operatorOpt, &operatorNonceRootChain)
	tx4, err := mintableToken.Mint(operatorOpt, addr4, ether(100))

	<-mintEvents
	<-mintEvents
	<-mintEvents
	<-mintEvents
	mintFilterrer.Unsubscribe()

	log.Info("Mint MintableToken to users")

	// 7. swap MintableToken to EtherToken
	setNonce(opt1, &addr1NonceRootChain)
	setNonce(opt2, &addr2NonceRootChain)
	setNonce(opt3, &addr3NonceRootChain)
	setNonce(opt4, &addr4NonceRootChain)

	if tx1, err = mintableToken.Approve(opt1, etherTokenAddr, ether(100)); err != nil {
		log.Error("Failed to approve MintableToken to EtherToken", "err", err)
	}
	if tx2, err = mintableToken.Approve(opt2, etherTokenAddr, ether(100)); err != nil {
		log.Error("Failed to approve MintableToken to EtherToken", "err", err)
	}
	if tx3, err = mintableToken.Approve(opt3, etherTokenAddr, ether(100)); err != nil {
		log.Error("Failed to approve MintableToken to EtherToken", "err", err)
	}
	if tx4, err = mintableToken.Approve(opt4, etherTokenAddr, ether(100)); err != nil {
		log.Error("Failed to approve MintableToken to EtherToken", "err", err)
	}

	log.Info("MintableToken is approved to EtherToken")

	waitTx(tx1.Hash())
	waitTx(tx2.Hash())
	waitTx(tx3.Hash())
	waitTx(tx4.Hash())

	setNonce(opt1, &addr1NonceRootChain)
	setNonce(opt2, &addr2NonceRootChain)
	setNonce(opt3, &addr3NonceRootChain)
	setNonce(opt4, &addr4NonceRootChain)

	tx1, _ = etherToken.Deposit(opt1, ether(100))
	tx2, _ = etherToken.Deposit(opt2, ether(100))
	tx3, _ = etherToken.Deposit(opt3, ether(100))
	tx4, _ = etherToken.Deposit(opt4, ether(100))

	waitTx(tx1.Hash())
	waitTx(tx2.Hash())
	waitTx(tx3.Hash())
	waitTx(tx4.Hash())

	log.Info("Swap MintableToken to EtherToken")

	for i, addr := range addrs {
		bal, err := etherToken.BalanceOf(baseCallOpt, addr)
		if err != nil {
			log.Error("Failed to get EtherToken balance", "err", err)
		}

		bal.Div(bal, ether(1))

		log.Info("EtherToken balance", "i", i, "balance", bal)
	}

	testPlsConfig.RootChainContract = rootchainAddr

	return rootchainAddr, rootchainContract, err
}

func newCanonical(n int, full bool) (ethdb.Database, *core.BlockChain, error) {
	gspec := core.DeveloperGenesisBlock(0, common.Address{}, operator, core.DefaultStaminaConfig)
	// Initialize a fresh chain with only a genesis block
	db := ethdb.NewMemDatabase()
	genesis := gspec.MustCommit(db)

	blockchain, _ := core.NewBlockChain(db, nil, params.PlasmaChainConfig, engine, testVmConfg, nil)
	// Create and inject the requested chain
	if n == 0 {
		return db, blockchain, nil
	}
	if full {
		// Full block-chain requested
		blocks := makeBlockChain(genesis, n, engine, db, canonicalSeed)
		_, err := blockchain.InsertChain(blocks)
		return db, blockchain, err
	}
	// Header-only chain requested
	headers := makeHeaderChain(genesis.Header(), n, engine, db, canonicalSeed)
	_, err := blockchain.InsertHeaderChain(headers, 1)
	return db, blockchain, err
}

func makeHeaderChain(parent *types.Header, n int, engine consensus.Engine, db ethdb.Database, seed int) []*types.Header {
	blocks := makeBlockChain(types.NewBlockWithHeader(parent), n, engine, db, seed)
	headers := make([]*types.Header, len(blocks))
	for i, block := range blocks {
		headers[i] = block.Header()
	}
	return headers
}

func makeBlockChain(parent *types.Block, n int, engine consensus.Engine, db ethdb.Database, seed int) []*types.Block {
	blocks, _ := core.GenerateChain(params.PlasmaChainConfig, parent, engine, db, n, func(i int, b *core.BlockGen) {
		b.SetCoinbase(common.Address{0: byte(seed), 19: byte(i)})
	})
	return blocks
}

func newTxPool(blockchain *core.BlockChain) *core.TxPool {
	pool := core.NewTxPool(*testTxPoolConfig, params.PlasmaChainConfig, blockchain)

	return pool
}

func tmpKeyStore() (string, *keystore.KeyStore) {
	d, err := ioutil.TempDir("/tmp", "eth-keystore-test")
	if err != nil {
		log.Error("Failed to set temporary keystore directory", "err", err)
	}
	ks := keystore.NewKeyStore(d, 2, 1)

	return d, ks
}

type testPlsBackend struct {
	acm        *accounts.Manager
	blockchain *core.BlockChain
	txPool     *core.TxPool
	db         ethdb.Database
}

func (b *testPlsBackend) AccountManager() *accounts.Manager { return b.acm }
func (b *testPlsBackend) BlockChain() *core.BlockChain      { return b.blockchain }
func (b *testPlsBackend) TxPool() *core.TxPool              { return b.txPool }
func (b *testPlsBackend) ChainDb() ethdb.Database           { return b.db }

func makePls() (*Plasma, *rpc.Server, string, error) {
	db, blockchain, err := newCanonical(0, true)

	if err != nil {
		log.Error("Failed to creaet blockchain", "err", err)
		return nil, nil, "", err
	}

	config := testPlsConfig
	chainConfig := params.PlasmaChainConfig

	rootchainAddress, rootchainContract, err := deployRootChain(blockchain.Genesis())

	if err != nil {
		log.Error("Failed to deploy rootchain contract", "err", err)
		return nil, nil, "", err
	}

	config.RootChainContract = rootchainAddress

	d, ks := tmpKeyStore()

	var oac accounts.Account
	var cac accounts.Account
	if oac, err = ks.ImportECDSA(operatorKey, ""); err != nil {
		log.Error("Failed to import operator account", "err", err)
	}

	if cac, err = ks.ImportECDSA(challengerKey, ""); err != nil {
		log.Error("Failed to import challenger account", "err", err)
	}

	if err = ks.Unlock(oac, ""); err != nil {
		log.Error("Failed to unlock operator account", "err", err)
	}

	if err = ks.Unlock(cac, ""); err != nil {
		log.Error("Failed to unlock challenger account", "err", err)
	}

	for _, key := range keys {
		var acc accounts.Account
		if acc, err = ks.ImportECDSA(key, ""); err != nil {
			log.Error("Failed to import user account", "err", err)
		}

		if err := ks.Unlock(acc, ""); err != nil {
			log.Error("Failed to unlock user  account", "err", err)
		}
	}

	config.Operator = oac
	config.Challenger = cac

	// configure account manager with temporary keystore backend
	backends := []accounts.Backend{
		ks,
	}
	accManager := accounts.NewManager(backends...)

	pls := &Plasma{
		config:         config,
		chainDb:        db,
		chainConfig:    chainConfig,
		eventMux:       new(event.TypeMux),
		accountManager: accManager,
		blockchain:     blockchain,
		engine:         engine,
		shutdownChan:   make(chan bool),
		networkID:      config.NetworkId,
		gasPrice:       config.MinerGasPrice,
		etherbase:      config.Etherbase,
		bloomRequests:  make(chan chan *bloombits.Retrieval),
		bloomIndexer:   NewBloomIndexer(db, params.BloomBitsBlocks, params.BloomConfirms),
	}
	pls.bloomIndexer.Start(pls.blockchain)
	pls.txPool = core.NewTxPool(config.TxPool, pls.chainConfig, pls.blockchain)

	if pls.protocolManager, err = NewProtocolManager(pls.chainConfig, config.SyncMode, config.NetworkId, pls.eventMux, pls.txPool, pls.engine, pls.blockchain, db, config.Whitelist); err != nil {
		return nil, nil, d, err
	}

	epochEnv := epoch.New()
	pls.miner = miner.New(pls, pls.chainConfig, pls.EventMux(), pls.engine, epochEnv, db, config.MinerRecommit, config.MinerGasFloor, config.MinerGasCeil, pls.isLocalBlock)
	pls.miner.SetExtra(makeExtraData(config.MinerExtraData))
	pls.APIBackend = &PlsAPIBackend{pls, nil}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.MinerGasPrice
	}
	pls.APIBackend.gpo = gasprice.NewOracle(pls.APIBackend, gpoParams)
	// Dial rootchain provider
	rootchainBackend, err := ethclient.Dial(config.RootChainURL)
	if err != nil {
		return nil, nil, d, err
	}
	log.Info("Rootchain provider connected", "url", config.RootChainURL)

	if err != nil {
		return nil, nil, d, err
	}

	stopFn := func() { pls.Stop() }
	txManager, err := tx.NewTransactionManager(ks, rootchainBackend, db, &config.TxConfig)

	if err != nil {
		return nil, nil, d, err
	}

	if pls.rootchainManager, err = NewRootChainManager(
		config,
		stopFn,
		pls.txPool,
		pls.blockchain,
		rootchainBackend,
		rootchainContract,
		pls.eventMux,
		pls.accountManager,
		txManager,
		pls.miner,
		epochEnv,
	); err != nil {
		return nil, nil, d, err
	}

	handler := rpc.NewServer()
	apis := pls.APIs()

	for _, api := range apis {
		if api.Service == nil || api.Namespace != "eth" {
			log.Debug("InProc skipped to register service", "service", api.Service, "namespace", api.Namespace)
			continue
		}

		if err := handler.RegisterName(api.Namespace, api.Service); err != nil {
			return nil, nil, d, err
		}
		log.Debug("InProc registered", "service", api.Service, "namespace", api.Namespace)
	}

	return pls, handler, d, nil
}

func setNonce(opt *bind.TransactOpts, nonce *uint64) {
	opt.Nonce = big.NewInt(int64(*nonce))
	*nonce++
}

func deployEtherTokenInChildChain(t *testing.T) {
	opt := makeTxOpt(operatorKey, 0, nil, nil)

	setNonce(opt, &operatorNonceChildChain)
	etherTokenAddrInChildChain, _, etherTokenInChildChain, err = ethertoken.DeployEtherToken(
		opt,
		plsClient,
		development,
		common.Address{},
		swapEnabledInChildChain,
	)

	if err != nil {
		t.Fatal("Failed to deploy EtherToken in child chain", "err", err)
	}

	log.Info("EtherToken deployed in child chain", "addr", etherTokenAddrInChildChain)

	return
}

func deployTokenContracts(t *testing.T) (*token.RequestableSimpleToken, *token.RequestableSimpleToken, common.Address, common.Address) {
	opt := makeTxOpt(operatorKey, 0, nil, nil)

	setNonce(opt, &operatorNonceRootChain)
	tokenAddrInRootChain, _, tokenInRootChain, err := token.DeployRequestableSimpleToken(
		opt,
		ethClient,
	)
	if err != nil {
		t.Fatal("Failed to deploy token contract in root chain", "err", err)
	}
	log.Info("Token deployed in root chain", "address", tokenAddrInRootChain)

	setNonce(opt, &operatorNonceChildChain)
	tokenAddrInChildChain, _, tokenInChildChain, err := token.DeployRequestableSimpleToken(
		opt,
		plsClient,
	)
	if err != nil {
		t.Fatal("Failed to deploy token contract in child chain", "err", err)
	}
	log.Info("Token deployed in child chain", "address", tokenAddrInChildChain)

	return tokenInRootChain, tokenInChildChain, tokenAddrInRootChain, tokenAddrInChildChain
}

func makeManager() (*RootChainManager, func(), error) {
	db, blockchain, _ := newCanonical(0, true)
	contractAddress, rootchainContract, err := deployRootChain(blockchain.Genesis())
	if err != nil {
		return nil, func() {}, err
	}
	wait(3)
	log.Info("Contract deployed at", "address", contractAddress)

	testPlsConfig.RootChainContract = contractAddress

	txPool := newTxPool(blockchain)
	minerBackend := &testPlsBackend{
		acm:        nil,
		blockchain: blockchain,
		txPool:     txPool,
		db:         db,
	}

	_, ks := tmpKeyStore()

	mux := new(event.TypeMux)
	epochEnv := epoch.New()
	miner := miner.New(minerBackend, params.PlasmaChainConfig, mux, engine, epochEnv, db, testPlsConfig.MinerRecommit, testPlsConfig.MinerGasFloor, testPlsConfig.MinerGasCeil, nil)

	account, err := ks.ImportECDSA(operatorKey, "")
	if err != nil {
		log.Error("Failed to import operator account", "err", err)
	}
	if err = ks.Unlock(account, ""); err != nil {
		log.Error("Failed to unlock operator account", "err", err)
	}
	// configure account manager with temporary keystore backend
	backends := []accounts.Backend{
		ks,
	}
	accManager := accounts.NewManager(backends...)
	txManager, err := tx.NewTransactionManager(ks, ethClient, db, &testPlsConfig.TxConfig)

	var rcm *RootChainManager

	stopFn := func() {
		blockchain.Stop()
		txPool.Stop()
		miner.Stop()
		mux.Stop()
		rcm.Stop()
	}
	rcm, err = NewRootChainManager(
		testPlsConfig,
		stopFn,
		txPool,
		blockchain,
		ethClient,
		rootchainContract,
		mux,
		accManager,
		txManager,
		miner,
		epochEnv,
	)

	if err != nil {
		return nil, func() {}, err
	}
	// TODO (aiden): there's no need to start miner in here, it starts when rcm connect to root chain contract by reading 1st NRE.
	//go Miner.Start(operator, &miner.NRE)
	return rcm, stopFn, nil
}

func makeTxOpt(key *ecdsa.PrivateKey, gasLimit uint64, gasPrice, value *big.Int) *bind.TransactOpts {
	opt := bind.NewKeyedTransactor(key)
	opt.GasLimit = defaultGasLimit
	opt.GasPrice = defaultGasPrice
	opt.Value = defaultValue

	if gasLimit != 0 {
		opt.GasLimit = gasLimit
	}

	if gasPrice != nil {
		opt.GasPrice = gasPrice
	}

	if value != nil {
		opt.Value = value
	}

	return opt
}

func ether(v float64) *big.Int {
	f := new(big.Float).Mul(big.NewFloat(v), big.NewFloat(1e18))
	out, _ := f.Int(nil)
	return out
}

func wait(t time.Duration) {
	timer := time.NewTimer(t * time.Second)
	<-timer.C
}

func waitTx(hash common.Hash) error {
	var receipt *types.Receipt
	for receipt, _ = ethClient.TransactionReceipt(context.Background(), hash); receipt == nil; {
		<-time.NewTimer(500 * time.Millisecond).C

		receipt, _ = ethClient.TransactionReceipt(context.Background(), hash)
	}

	if receipt.Status == 0 {
		log.Error("transaction reverted", "hash", hash)
		return errors.New("transaction reverted")
	}

	return nil
}

// TODO: any user sends tx
func makeSampleTx(rcm *RootChainManager) error {
	pool := rcm.txPool

	// self transfer
	var err error

	tx := types.NewTransaction(operatorNonceChildChain, operator, nil, 21000, nil, []byte{})
	operatorNonceChildChain++

	signer := types.NewEIP155Signer(params.PlasmaChainConfig.ChainID)

	tx, err = types.SignTx(tx, signer, operatorKey)
	if err != nil {
		log.Error("Failed to sign sample tx", "err", err)
		return err
	}

	if err = pool.AddLocal(tx); err != nil {
		log.Error("Failed to insert sample tx to tx pool", "err", err)

		return err
	}

	log.Debug("Sample transaction is submitted in child chian")

	return nil
}

func checkBlock(pls *Plasma, pbMinedEvents *event.TypeMuxSubscription, pbSubmitedEvents chan *rootchain.RootChainBlockSubmitted, expectedIsRequest bool, expectedFork, expectedBlockNumber int64) error {
	// TODO: delete below line after genesis.Difficulty is set 0
	expectedFork += 1

	setNonce(operatorOpt, &operatorNonceRootChain) // due to block submit

	outC := make(chan struct{})
	errC := make(chan error)
	defer close(outC)
	defer close(errC)

	timer := time.NewTimer((testPlsConfig.MinerRecommit + testPlsConfig.TxConfig.Interval) * 2)
	defer timer.Stop()

	log.Error("Check block", "expectedBlockNumber", expectedBlockNumber)

	quit := make(chan struct{})
	defer close(quit)

	go func() {
		select {
		case _, ok := <-timer.C:
			if ok {
				errC <- errors.New("Out of time")
			}
		case <-quit:
			return
		}
	}()

	go func() {
		outC2 := make(chan struct{})
		defer close(outC2)
		// use goroutine to read both events
		var blockInfo core.NewMinedBlockEvent
		go func() {
			e, ok := <-pbMinedEvents.Chan()
			if !ok {
				log.Error("cannot read from mined block channel")
				return
			}

			blockInfo = e.Data.(core.NewMinedBlockEvent)
			outC2 <- struct{}{}
		}()

		go func() {
			<-pbSubmitedEvents
			outC2 <- struct{}{}
		}()

		<-outC2
		<-outC2

		block := blockInfo.Block

		log.Warn("Check Block Number", "expectedBlockNumber", expectedBlockNumber, "minedBlockNumber", block.NumberU64(), "forkNumber", block.Difficulty().Uint64())

		// check block number.
		if expectedBlockNumber != block.Number().Int64() {
			errC <- errors.New(fmt.Sprintf("Expected block number: %d, actual block %d", expectedBlockNumber, block.Number().Int64()))
			return
		}

		// check fork number
		if expectedFork != block.Difficulty().Int64() {
			errC <- errors.New(fmt.Sprintf("Block Expected ForkNumber: %d, Actual ForkNumber %d", expectedFork, block.Difficulty().Int64()))
			return
		}

		// check isRequest.
		if block.IsRequest() != expectedIsRequest {
			log.Error("txs length check", "length", len(block.Transactions()))

			tx, _ := block.Transactions()[0].AsMessage(types.HomesteadSigner{})
			log.Error("tx sender address", "sender", tx.From())
			errC <- errors.New(fmt.Sprintf("Expected isRequest: %t, Actual isRequest %t", expectedIsRequest, block.IsRequest()))
			return
		}

		pb, _ := pls.rootchainManager.getBlock(big.NewInt(int64(pls.rootchainManager.state.currentFork)), block.Number())

		if pb.Timestamp == 0 {
			log.Debug("Submitted plasma block", "pb", pb)
			log.Debug("Mined plasma block", "b", block)
			errC <- errors.New("Plasma block is not submitted yet.")
			return
		}

		if pb.IsRequest != block.IsRequest() {
			errC <- errors.New(fmt.Sprintf("Block Expected isRequest: %t, Actual isRequest %t", pb.IsRequest, block.IsRequest()))
			return
		}

		pbStateRoot := pb.StatesRoot[:]
		bStateRoot := block.Header().Root.Bytes()
		if bytes.Compare(pbStateRoot, bStateRoot) != 0 {
			errC <- errors.New(fmt.Sprintf("Block Expected stateRoot: %s, Actual stateRoot: %s", common.Bytes2Hex(pbStateRoot), common.Bytes2Hex(bStateRoot)))
			return
		}

		pbTxRoot := pb.TransactionsRoot[:]
		bTxRoot := block.Header().TxHash.Bytes()
		if bytes.Compare(pbTxRoot, bTxRoot) != 0 {
			errC <- errors.New(fmt.Sprintf("Block Expected txRoot: %s, Actual txRoot: %s", common.Bytes2Hex(pbTxRoot), common.Bytes2Hex(bTxRoot)))
			return
		}

		pbReceiptsRoot := pb.ReceiptsRoot[:]
		bReceiptsRoot := block.Header().ReceiptHash.Bytes()
		if bytes.Compare(pbReceiptsRoot, bReceiptsRoot) != 0 {
			errC <- errors.New(fmt.Sprintf("Block Expected receiptsRoot: %s, Actual receiptsRoot: %s", common.Bytes2Hex(pbReceiptsRoot), common.Bytes2Hex(bReceiptsRoot)))
			return
		}
		log.Debug("Check block finished")
		outC <- struct{}{}
	}()

	select {
	case <-outC:
		return nil
	case err := <-errC:
		return err
	}
}

// checkBalance check after = before + diff if offset is nil.
// Otherwise, check -offset < after - (before + diff) < offset
func checkBalance(before, after *big.Int, diff, offset *big.Int, caption string) error {
	// truncate up to 100 ether
	truncate := func(v *big.Int) (*big.Int, *big.Int) {
		v0 := v

		if v.Sign() < 0 {
			v0.Abs(v0)
		}

		_, v1 := new(big.Int).DivMod(v0, ether(100), new(big.Int))

		v1, v2 := new(big.Int).DivMod(v1, ether(1), new(big.Int))
		v2.Div(v2, big.NewInt(1e9)) // remove 9 digits from wei

		if v.Sign() < 0 {
			v1.Neg(v1)
		}
		return v1, v2
	}

	toString := func(q, r *big.Int) string {

		return fmt.Sprintf("%s.%s", q.String(), r.String())
	}

	b := before
	a := after

	if offset == nil {
		if a.Cmp(new(big.Int).Add(b, diff)) != 0 {
			bt := toString(truncate(b))
			at := toString(truncate(a))
			dt := toString(truncate(diff))

			wait(1)
			return errors.New(fmt.Sprintf(caption+"\t : Expected %s (after) == %s (before) + %s (diff), but it isn't", at, bt, dt))
		}
		return nil
	}

	target := new(big.Int).Sub(new(big.Int).Sub(a, b), diff)
	e := new(big.Int).Abs(offset)
	s := new(big.Int).Neg(e)

	// out of range
	if s.Cmp(target) > 0 || target.Cmp(e) > 0 {
		st := toString(truncate(s))
		et := toString(truncate(e))
		targett := toString(truncate(target))

		wait(1)
		return errors.New(fmt.Sprintf(caption+"\t : Expected %s (-offset) < %s (after - before - diff) < %s (+offset), but it isn't", st, targett, et))
	}

	return nil
}

func getETHBalances(addrs []common.Address) []*big.Int {
	balances := make([]*big.Int, len(addrs))

	for i, addr := range addrs {
		// balances[i] would be nil if ethClient.BalanceAt fails
		balances[i], _ = etherToken.BalanceOf(baseCallOpt, addr)
	}

	return balances
}

func getPETHBalances(addrs []common.Address) []*big.Int {
	balances := make([]*big.Int, len(addrs))

	for i, addr := range addrs {
		// balances[i] would be nil if ethClient.BalanceAt fails
		balances[i], _ = plsClient.BalanceAt(context.Background(), addr, nil)
	}

	return balances
}

func getEtherTokenBalances(addrs []common.Address) []*big.Int {
	balances := make([]*big.Int, len(addrs))

	for i, addr := range addrs {
		// balances[i] would be nil if ethClient.BalanceAt fails
		balances[i], _ = etherToken.BalanceOf(baseCallOpt, addr)
	}

	return balances
}

func getPEtherTokenBalances(addrs []common.Address) []*big.Int {
	balances := make([]*big.Int, len(addrs))

	for i, addr := range addrs {
		// balances[i] would be nil if ethClient.BalanceAt fails
		balances[i], _ = etherToken.BalanceOf(baseCallOpt, addr)
	}

	return balances
}

func getTokenBalances(addrs []common.Address, tokenContract *token.RequestableSimpleToken) []*big.Int {
	balances := make([]*big.Int, len(addrs))

	for i, addr := range addrs {
		// balances[i] would be nil if ethClient.BalanceAt fails
		balances[i], _ = tokenContract.Balances(baseCallOpt, addr)
	}

	return balances
}
