// Copyright 2019 The multi-geth Authors
// This file is part of the multi-geth library.
//
// The multi-geth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The multi-geth library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the multi-geth library. If not, see <http://www.gnu.org/licenses/>.

package ctypes

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// This file holds the Configurator interfaces.
// Interface methods follow distinct naming and signature patterns
// to enable abstracted logic.
//
// All methods are in pairs; Get'ers and Set'ers.
// Some Set methods are prefixed with Must, ie MustSet. These methods
// are allowed to return errors for debugging and logging, but
// any non-nil errors returned should stop program execution.
//
// All Forking methods (getting and setting Hard-fork requiring protocol changes)
// are suffixed with "Transition", and use *uint64 as in- and out-put variable types.
// A pointer is used because it's important that the fields can be nil, to signal
// being unset (as opposed to zero value uint64 == 0, ie from Genesis).

type Configurator interface {
	ChainConfigurator
	GenesisBlocker
}

type ChainConfigurator interface {
	String() string

	ProtocolSpecifier
	Forker
	ConsensusEnginator // Consensus Engine
	// CHTer
}

// ProtocolSpecifier defines protocol interfaces that are agnostic of consensus engine.
type ProtocolSpecifier interface {
	GetAccountStartNonce() *uint64
	SetAccountStartNonce(n *uint64) error
	GetMaximumExtraDataSize() *uint64
	SetMaximumExtraDataSize(n *uint64) error
	GetMinGasLimit() *uint64
	SetMinGasLimit(n *uint64) error
	GetGasLimitBoundDivisor() *uint64
	SetGasLimitBoundDivisor(n *uint64) error
	GetNetworkID() *uint64
	SetNetworkID(n *uint64) error
	GetChainID() *big.Int
	SetChainID(i *big.Int) error
	GetSupportedProtocolVersions() []uint
	SetSupportedProtocolVersions(p []uint) error
	GetMaxCodeSize() *uint64
	SetMaxCodeSize(n *uint64) error

	GetElasticityMultiplier() uint64
	SetElasticityMultiplier(n uint64) error
	GetBaseFeeChangeDenominator() uint64
	SetBaseFeeChangeDenominator(n uint64) error

	// Be careful with EIP2.
	// It is a messy EIP, specifying diverse changes, like difficulty, intrinsic gas costs for contract creation,
	// txpool management, and contract OoG handling.
	// It is both Ethash-specific and _not_.
	GetEIP2Transition() *uint64
	SetEIP2Transition(n *uint64) error

	GetEIP7Transition() *uint64
	SetEIP7Transition(n *uint64) error
	GetEIP150Transition() *uint64
	SetEIP150Transition(n *uint64) error
	GetEIP152Transition() *uint64
	SetEIP152Transition(n *uint64) error
	GetEIP160Transition() *uint64
	SetEIP160Transition(n *uint64) error
	GetEIP161abcTransition() *uint64
	SetEIP161abcTransition(n *uint64) error
	GetEIP161dTransition() *uint64
	SetEIP161dTransition(n *uint64) error
	GetEIP170Transition() *uint64
	SetEIP170Transition(n *uint64) error
	GetEIP155Transition() *uint64
	SetEIP155Transition(n *uint64) error
	GetEIP140Transition() *uint64
	SetEIP140Transition(n *uint64) error
	GetEIP198Transition() *uint64
	SetEIP198Transition(n *uint64) error
	GetEIP211Transition() *uint64
	SetEIP211Transition(n *uint64) error
	GetEIP212Transition() *uint64
	SetEIP212Transition(n *uint64) error
	GetEIP213Transition() *uint64
	SetEIP213Transition(n *uint64) error
	GetEIP214Transition() *uint64
	SetEIP214Transition(n *uint64) error
	GetEIP658Transition() *uint64
	SetEIP658Transition(n *uint64) error
	GetEIP145Transition() *uint64
	SetEIP145Transition(n *uint64) error
	GetEIP1014Transition() *uint64
	SetEIP1014Transition(n *uint64) error
	GetEIP1052Transition() *uint64
	SetEIP1052Transition(n *uint64) error
	GetEIP1283Transition() *uint64
	SetEIP1283Transition(n *uint64) error
	GetEIP1283DisableTransition() *uint64
	SetEIP1283DisableTransition(n *uint64) error
	GetEIP1108Transition() *uint64
	SetEIP1108Transition(n *uint64) error
	GetEIP2200Transition() *uint64
	SetEIP2200Transition(n *uint64) error
	GetEIP2200DisableTransition() *uint64
	SetEIP2200DisableTransition(n *uint64) error
	GetEIP1344Transition() *uint64
	SetEIP1344Transition(n *uint64) error
	GetEIP1884Transition() *uint64
	SetEIP1884Transition(n *uint64) error
	GetEIP2028Transition() *uint64
	SetEIP2028Transition(n *uint64) error
	GetECIP1080Transition() *uint64
	SetECIP1080Transition(n *uint64) error
	GetEIP1706Transition() *uint64
	SetEIP1706Transition(n *uint64) error
	GetEIP2537Transition() *uint64
	SetEIP2537Transition(n *uint64) error

	GetECBP1100Transition() *uint64
	SetECBP1100Transition(n *uint64) error
	GetECBP1100DeactivateTransition() *uint64
	SetECBP1100DeactivateTransition(n *uint64) error

	GetEIP2315Transition() *uint64
	SetEIP2315Transition(n *uint64) error

	// ModExp gas cost
	GetEIP2565Transition() *uint64
	SetEIP2565Transition(n *uint64) error

	// Gas cost increases for state access opcodes
	GetEIP2929Transition() *uint64
	SetEIP2929Transition(n *uint64) error

	// Optional access lists
	GetEIP2930Transition() *uint64
	SetEIP2930Transition(n *uint64) error

	// Typed transaction envelope
	GetEIP2718Transition() *uint64
	SetEIP2718Transition(n *uint64) error

	GetEIP1559Transition() *uint64
	SetEIP1559Transition(n *uint64) error

	GetEIP3541Transition() *uint64
	SetEIP3541Transition(n *uint64) error

	GetEIP3529Transition() *uint64
	SetEIP3529Transition(n *uint64) error

	GetEIP3198Transition() *uint64
	SetEIP3198Transition(n *uint64) error

	// EIP4399 is the RANDOM opcode.
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-4399.md
	GetEIP4399Transition() *uint64
	SetEIP4399Transition(n *uint64) error

	// Shanghai:
	//
	// EIP3651: Warm COINBASE
	GetEIP3651TransitionTime() *uint64
	SetEIP3651TransitionTime(n *uint64) error
	// EIP3855: PUSH0 instruction
	GetEIP3855TransitionTime() *uint64
	SetEIP3855TransitionTime(n *uint64) error
	// EIP3860: Limit and meter initcode
	GetEIP3860TransitionTime() *uint64
	SetEIP3860TransitionTime(n *uint64) error
	// EIP4895: Beacon chain push withdrawals as operations
	GetEIP4895TransitionTime() *uint64
	SetEIP4895TransitionTime(n *uint64) error
	// EIP6049: Deprecate SELFDESTRUCT
	GetEIP6049TransitionTime() *uint64
	SetEIP6049TransitionTime(n *uint64) error

	// Shanghai expressed as block activation numbers:
	GetEIP3651Transition() *uint64
	SetEIP3651Transition(n *uint64) error
	GetEIP3855Transition() *uint64
	SetEIP3855Transition(n *uint64) error
	GetEIP3860Transition() *uint64
	SetEIP3860Transition(n *uint64) error
	GetEIP4895Transition() *uint64
	SetEIP4895Transition(n *uint64) error
	GetEIP6049Transition() *uint64
	SetEIP6049Transition(n *uint64) error

	// GetMergeVirtualTransition is a Virtual fork after The Merge to use as a network splitter
	GetMergeVirtualTransition() *uint64
	SetMergeVirtualTransition(n *uint64) error

	// Cancun:
	// EIP4844 - Shard Blob Transactions - https://eips.ethereum.org/EIPS/eip-4844
	GetEIP4844TransitionTime() *uint64
	SetEIP4844TransitionTime(n *uint64) error

	// EIP1153 - Transient Storage opcodes - https://eips.ethereum.org/EIPS/eip-1153
	GetEIP1153TransitionTime() *uint64
	SetEIP1153TransitionTime(n *uint64) error

	// EIP5656 - MCOPY - Memory copying instruction - https://eips.ethereum.org/EIPS/eip-5656
	GetEIP5656TransitionTime() *uint64
	SetEIP5656TransitionTime(n *uint64) error

	// EIP6780 - SELFDESTRUCT only in same transaction - https://eips.ethereum.org/EIPS/eip-6780
	GetEIP6780TransitionTime() *uint64
	SetEIP6780TransitionTime(n *uint64) error

	GetEticaSmartContractv2Transition() *uint64
	SetEticaSmartContractv2Transition(n *uint64) error

}

type Forker interface {
	// IsEnabled tells if interface has met or exceeded a fork block number.
	// eg. IsEnabled(c.GetEIP1108Transition, big.NewInt(42)))
	IsEnabled(fn func() *uint64, n *big.Int) bool
	IsEnabledByTime(fn func() *uint64, n *uint64) bool

	// ForkCanonHash yields arbitrary number/hash pairs.
	// This is an abstraction derived from the original EIP150 implementation.
	GetForkCanonHash(n uint64) common.Hash
	SetForkCanonHash(n uint64, h common.Hash) error
	GetForkCanonHashes() map[uint64]common.Hash
}

type ConsensusEnginator interface {
	GetConsensusEngineType() ConsensusEngineT
	MustSetConsensusEngineType(t ConsensusEngineT) error
	GetIsDevMode() bool
	SetDevMode(devMode bool) error

	EthashConfigurator
	CliqueConfigurator
	Lyra2Configurator
}

type EthashConfigurator interface {
	GetEthashMinimumDifficulty() *big.Int
	SetEthashMinimumDifficulty(i *big.Int) error
	GetEthashDifficultyBoundDivisor() *big.Int
	SetEthashDifficultyBoundDivisor(i *big.Int) error
	GetEthashDurationLimit() *big.Int
	SetEthashDurationLimit(i *big.Int) error
	GetEthashHomesteadTransition() *uint64
	SetEthashHomesteadTransition(n *uint64) error

	// GetEthashEIP779Transition should return the block if the node wants the fork.
	// Otherwise, nil should be returned.
	GetEthashEIP779Transition() *uint64 // DAO

	// SetEthashEIP779Transition should turn DAO support on (nonnil) or off (nil).
	SetEthashEIP779Transition(n *uint64) error
	GetEthashEIP649Transition() *uint64
	SetEthashEIP649Transition(n *uint64) error
	GetEthashEIP1234Transition() *uint64
	SetEthashEIP1234Transition(n *uint64) error
	GetEthashEIP2384Transition() *uint64
	SetEthashEIP2384Transition(n *uint64) error
	GetEthashEIP3554Transition() *uint64
	SetEthashEIP3554Transition(n *uint64) error
	GetEthashEIP4345Transition() *uint64
	SetEthashEIP4345Transition(n *uint64) error
	GetEthashECIP1010PauseTransition() *uint64
	SetEthashECIP1010PauseTransition(n *uint64) error
	GetEthashECIP1010ContinueTransition() *uint64
	SetEthashECIP1010ContinueTransition(n *uint64) error
	GetEthashECIP1017Transition() *uint64
	SetEthashECIP1017Transition(n *uint64) error
	GetEthashECIP1017EraRounds() *uint64
	SetEthashECIP1017EraRounds(n *uint64) error
	GetEthashEIP100BTransition() *uint64
	SetEthashEIP100BTransition(n *uint64) error
	GetEthashECIP1041Transition() *uint64
	SetEthashECIP1041Transition(n *uint64) error
	GetEthashECIP1099Transition() *uint64
	SetEthashECIP1099Transition(n *uint64) error
	GetEthashEIP5133Transition() *uint64 // Gray Glacier difficulty bomb delay
	SetEthashEIP5133Transition(n *uint64) error

	GetEthashTerminalTotalDifficulty() *big.Int
	SetEthashTerminalTotalDifficulty(n *big.Int) error

	GetEthashTerminalTotalDifficultyPassed() bool
	SetEthashTerminalTotalDifficultyPassed(t bool) error

	IsTerminalPoWBlock(parentTotalDiff *big.Int, totalDiff *big.Int) bool

	GetEthashDifficultyBombDelaySchedule() Uint64BigMapEncodesHex
	SetEthashDifficultyBombDelaySchedule(m Uint64BigMapEncodesHex) error
	GetEthashBlockRewardSchedule() Uint64BigMapEncodesHex
	SetEthashBlockRewardSchedule(m Uint64BigMapEncodesHex) error

	GetEthashETIP1017Transition() *uint64
	SetEthashETIP1017Transition(n *uint64) error
}

type CliqueConfigurator interface {
	GetCliquePeriod() uint64
	SetCliquePeriod(n uint64) error
	GetCliqueEpoch() uint64
	SetCliqueEpoch(n uint64) error
}

type Lyra2Configurator interface {
	GetLyra2NonceTransition() *uint64
	SetLyra2NonceTransition(n *uint64) error
}

type BlockSealer interface {
	GetSealingType() BlockSealingT
	SetSealingType(t BlockSealingT) error
	BlockSealerEthereum
}

type BlockSealerEthereum interface {
	GetGenesisSealerEthereumNonce() uint64
	SetGenesisSealerEthereumNonce(n uint64) error
	GetGenesisSealerEthereumMixHash() common.Hash
	SetGenesisSealerEthereumMixHash(h common.Hash) error
}

type GenesisBlocker interface {
	BlockSealer
	Accounter
	GetGenesisDifficulty() *big.Int
	SetGenesisDifficulty(i *big.Int) error
	GetGenesisAuthor() common.Address
	SetGenesisAuthor(a common.Address) error
	GetGenesisTimestamp() uint64
	SetGenesisTimestamp(u uint64) error
	GetGenesisParentHash() common.Hash
	SetGenesisParentHash(h common.Hash) error
	GetGenesisExtraData() []byte
	SetGenesisExtraData(b []byte) error
	GetGenesisGasLimit() uint64
	SetGenesisGasLimit(u uint64) error
}

type Accounter interface {
	ForEachAccount(fn func(address common.Address, bal *big.Int, nonce uint64, code []byte, storage map[common.Hash]common.Hash) error) error
	UpdateAccount(address common.Address, bal *big.Int, nonce uint64, code []byte, storage map[common.Hash]common.Hash) error
}
