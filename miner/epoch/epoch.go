package epoch

import (
	"math/big"
	"sync"
)

type EpochEnvironment struct {
	IsRequest          bool
	UserActivated      bool
	Rebase             bool
	Completed          bool
	NumBlockMined      *big.Int
	EpochLength        *big.Int
	CurrentFork        *big.Int
	LastFinalizedBlock *big.Int

	lock sync.Mutex
}

func (self *EpochEnvironment) SetIsRequest(b bool) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.IsRequest = b
}

func (self *EpochEnvironment) SetUserActivated(b bool) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.UserActivated = b
}

func (self *EpochEnvironment) SetRebase(b bool) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.Rebase = b
}

func (self *EpochEnvironment) SetCompleted(b bool) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.Completed = b
}

func (self *EpochEnvironment) SetNumBlockMined(n *big.Int) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.NumBlockMined = n
}

func (self *EpochEnvironment) SetEpochLength(l *big.Int) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.EpochLength = l
}

func (self *EpochEnvironment) SetCurrentFork(f *big.Int) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.CurrentFork = f
}

func (self *EpochEnvironment) SetLastFinalizedBlock(n *big.Int) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.LastFinalizedBlock = n
}

func (self *EpochEnvironment) Lock() {
	self.lock.Lock()
}

func (self *EpochEnvironment) Unlock() {
	self.lock.Unlock()
}
