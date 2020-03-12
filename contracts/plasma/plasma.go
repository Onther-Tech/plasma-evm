package plasma

//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/RootChain.sol --pkg rootchain --out rootchain/rootchain.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/handlers/EpochHandler.sol --pkg epochhandler --out epochhandler/epochhandler.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/handlers/SubmitHandler.sol --pkg submithandler --out submithandler/submithandler.go

//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/stake/tokens//TON.sol --pkg ton --out ton/ton.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/stake/tokens/WTON.sol --pkg wton --out wton/wton.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/stake/managers/DepositManager.sol --pkg depositmanager --out depositmanager/depositmanager.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/stake/managers/SeigManager.sol --pkg seigmanager --out seigmanager/seigmanager.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/stake/RootChainRegistry.sol --pkg rootchainregistry --out rootchainregistry/rootchainregistry.go

//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/RequestableSimpleToken.sol --pkg token --out token/token.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/node_modules/openzeppelin-solidity/contracts/token/ERC20/ERC20Mintable.sol --pkg mintabletoken --out mintabletoken/mintabletoken.go

//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/EtherToken.sol --pkg ethertoken --out ethertoken/ethertoken.go

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/depositmanager"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/epochhandler"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/ethertoken"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/mintabletoken"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchain"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchainregistry"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/seigmanager"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/submithandler"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/ton"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/wton"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/rawdb"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/pls"
)

func DeployPlasmaContracts(opt *bind.TransactOpts, backend *ethclient.Client, cfg *pls.Config, withPETH bool, development bool, NRELength *big.Int) (common.Address, error) {
	opt.GasLimit = 7000000

	wait := func(hash common.Hash) {
		<-time.NewTimer(1 * time.Second).C

		for receipt, _ := backend.TransactionReceipt(context.Background(), hash); receipt == nil; {
			<-time.NewTimer(1 * time.Second).C

			receipt, _ = backend.TransactionReceipt(context.Background(), hash)
		}
	}

	dummyDB := rawdb.NewMemoryDatabase()
	defer dummyDB.Close()
	var dummyBlock *types.Block

	if withPETH {
		// give PETH to operator
		dummyBlock = core.DeveloperGenesisBlock(
			1,
			common.HexToAddress("0xdead"),
			opt.From,
			cfg.StaminaConfig,
		).ToBlock(dummyDB)
	} else {
		// Nobody has PETH in genesis block
		dummyBlock = core.DefaultGenesisBlock(
			common.HexToAddress("0xdead"),
			opt.From,
			cfg.StaminaConfig,
		).ToBlock(dummyDB)
	}

	// 1. deploy MintableToken
	mintableTokenContract, tx, _, err := mintabletoken.DeployERC20Mintable(opt, backend)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy MintableToken contract: %v", err))
	}
	log.Info("Deploy MintableToken contract", "hash", tx.Hash(), "address", mintableTokenContract)

	log.Info("Wait until deploy transaction is mined")
	wait(tx.Hash())

	// 2. deploy EtherToken
	etherTokenContract, tx, etherToken, err := ethertoken.DeployEtherToken(opt, backend, development, mintableTokenContract, false)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy EtherToken contract: %v", err))
	}
	log.Info("Deploy EtherToken contract", "hash", tx.Hash(), "address", etherTokenContract)

	log.Info("Wait until deploy transaction is mined")
	wait(tx.Hash())

	// 3. deploy EpochHandler
	epochHandlerContract, tx, _, err := epochhandler.DeployEpochHandler(opt, backend)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy EpochHandler contract: %v", err))
	}
	log.Info("Deploy EpochHandler contract", "hash", tx.Hash(), "address", epochHandlerContract)

	log.Info("Wait until deploy transaction is mined")
	wait(tx.Hash())

	// 4. deploy SubmitHandler
	submitHandlerContract, tx, _, err := submithandler.DeploySubmitHandler(opt, backend, epochHandlerContract)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy SubmitHandler contract: %v", err))
	}
	log.Info("Deploy EpochHandler contract", "hash", tx.Hash(), "address", epochHandlerContract)

	log.Info("Wait until deploy transaction is mined")
	wait(tx.Hash())

	// 5. deploy RootChain
	rootchainContract, tx, _, err := rootchain.DeployRootChain(opt, backend, epochHandlerContract, submitHandlerContract, etherTokenContract, development, NRELength, dummyBlock.Root(), dummyBlock.TxHash(), dummyBlock.ReceiptHash())
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy RootChain contract: %v", err))
	}
	log.Info("Deploy RootChain contract", "hash", tx.Hash(), "address", rootchainContract)
	wait(tx.Hash())

	// 6. initialize EtherToken
	tx, err = etherToken.Init(opt, rootchainContract)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to initialize EtherToken: %v", err))
	}
	log.Info("Initialize EtherToken", "hash", tx.Hash())
	wait(tx.Hash())

	return rootchainContract, nil
}

type seigManagerSetter interface {
	SetSeigManager(opts *bind.TransactOpts, _seigManager common.Address) (*types.Transaction, error)
	SeigManager(opts *bind.CallOpts) (common.Address, error)
}

func DeployManagers(
	opt *bind.TransactOpts,
	backend *ethclient.Client,
	withdrawalDelay *big.Int,
	seigPerBlock *big.Int,
	_tonAddr common.Address,
	_wtonAddr common.Address,
) (
	tonAddr common.Address,
	wtonAddr common.Address,
	registryAddr common.Address,
	depositManagerAddr common.Address,
	seigManagerAddr common.Address,
	err error,
) {
	opt.GasLimit = 7000000

	var (
		//TON            *ton.TON
		WTON *wton.WTON
		//registry       *stakingmanager.RootChainRegistry
		depositManager *depositmanager.DepositManager
		//seigManager    *stakingmanager.SeigManager

		tx *types.Transaction
	)

	// 1. deploy TON
	log.Info("1. deploy TON contract")
	if (_tonAddr == common.Address{}) {
		if tonAddr, tx, _, err = ton.DeployTON(opt, backend); err != nil {
			err = errors.New(fmt.Sprintf("Failed to deploy TON: %v", err))
			return
		}

		if err = WaitTx(backend, tx.Hash()); err != nil {
			err = errors.New(fmt.Sprintf("Failed to deploy TON: %v", err))
			return
		}

		log.Info("TON deployed", "addr", tonAddr.String(), "tx", tx.Hash())
	} else {
		tonAddr = _tonAddr
		log.Warn("use TON contract at %s", tonAddr.String())
		if _, err = ton.NewTON(tonAddr, backend); err != nil {
			err = errors.New(fmt.Sprintf("Failed to instantiate TON: %v", err))
			return
		}
	}

	// 2. deploy WTON
	log.Info("2. deploy WTON contract")
	if (_wtonAddr == common.Address{}) {
		if wtonAddr, tx, WTON, err = wton.DeployWTON(opt, backend, tonAddr); err != nil {
			err = errors.New(fmt.Sprintf("Failed to deploy WTON: %v", err))
			return
		}

		if err = WaitTx(backend, tx.Hash()); err != nil {
			err = errors.New(fmt.Sprintf("Failed to deploy WTON: %v", err))
			return
		}
		log.Info("WTON deployed", "addr", wtonAddr.String(), "tx", tx.Hash())
	} else {
		wtonAddr = _wtonAddr
		log.Warn("use WTON contract at %s", wtonAddr.String())
		if WTON, err = wton.NewWTON(wtonAddr, backend); err != nil {
			err = errors.New(fmt.Sprintf("Failed to instantiate WTON: %v", err))
			return
		}
	}

	if addr, _ := WTON.SeigManager(&bind.CallOpts{Pending: false}); (addr != common.Address{}) {
		err = errors.New("WTON already set SeigManager")
		return
	}

	// 3. deploy RootChainRegistry
	log.Info("3. deploy RootChainRegistry")
	if registryAddr, tx, _, err = rootchainregistry.DeployRootChainRegistry(opt, backend); err != nil {
		err = errors.New(fmt.Sprintf("Failed to deploy RootChainRegistry: %v", err))
		return
	}

	if err = WaitTx(backend, tx.Hash()); err != nil {
		err = errors.New(fmt.Sprintf("Failed to deploy RootChainRegistry: %v", err))
		return
	}
	log.Info("RootChainRegistry deployed", "addr", registryAddr.String(), "tx", tx.Hash())

	// 4. deploy DepositManager
	log.Info("4. deploy DepositManager")
	if depositManagerAddr, tx, depositManager, err = depositmanager.DeployDepositManager(opt, backend, wtonAddr, registryAddr, withdrawalDelay); err != nil {
		err = errors.New(fmt.Sprintf("Failed to deploy DepositManager: %v", err))
		return
	}

	if err = WaitTx(backend, tx.Hash()); err != nil {
		err = errors.New(fmt.Sprintf("Failed to deploy DepositManager: %v", err))
		return
	}
	log.Info("DepositManager deployed", "addr", depositManagerAddr.String(), "tx", tx.Hash())

	// 5. deploy SeigManager
	log.Info("5. deploy SeigManager")
	if seigManagerAddr, tx, _, err = seigmanager.DeploySeigManager(opt, backend, tonAddr, wtonAddr, registryAddr, depositManagerAddr, seigPerBlock); err != nil {
		err = errors.New(fmt.Sprintf("Failed to deploy SeigManager: %v", err))
		return
	}

	if err = WaitTx(backend, tx.Hash()); err != nil {
		err = errors.New(fmt.Sprintf("Failed to deploy SeigManager: %v", err))
		return
	}
	log.Info("SeigManager deployed", "addr", seigManagerAddr.String(), "tx", tx.Hash())

	// 6. add WTON minter role to SeigManager
	log.Info("6. add WTON minter role to SeigManager")
	if tx, err = WTON.AddMinter(opt, seigManagerAddr); err != nil {
		err = errors.New(fmt.Sprintf("Failed to add WTON minter role to SeigManager: %v", err))
		return
	}
	if err = WaitTx(backend, tx.Hash()); err != nil {
		err = errors.New(fmt.Sprintf("Failed to add WTON minter role to SeigManager: %v", err))
		return
	}
	log.Info("Add WTON minter role to SeigManager", "tx", tx.Hash())

	// 7. set seig manager to contracts
	var contracts = []seigManagerSetter{depositManager, WTON}
	targets := []string{"DepositManager", "WTON"}

	log.Info("Setting SeigManager address to target contracts", "targets", targets)
	for i, c := range contracts {
		target := targets[i]

		if tx, err = c.SetSeigManager(opt, seigManagerAddr); err != nil {
			err = errors.New(fmt.Sprintf("Failed to set SeigManager to %s: %v", target, err))
			return
		}
		if err = WaitTx(backend, tx.Hash()); err != nil {
			err = errors.New(fmt.Sprintf("Failed to set SeigManager to %s: %v", target, err))
			return
		}

		log.Info("Set SeigManager to target cotnract", "target", target, "tx", tx.Hash())
	}

	return
}

func WaitTx(backend *ethclient.Client, hash common.Hash) error {
	var receipt *types.Receipt

	<-time.NewTimer(1 * time.Second).C

	for receipt, _ = backend.TransactionReceipt(context.Background(), hash); receipt == nil; {
		<-time.NewTimer(1 * time.Second).C

		receipt, _ = backend.TransactionReceipt(context.Background(), hash)
	}

	if receipt.Status == 0 {
		return errors.New(fmt.Sprintf("Transaction reverted: %s", receipt.TxHash.String()))
	}
	return nil
}
