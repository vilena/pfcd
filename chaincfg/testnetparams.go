// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2015-2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/picfight/pfcd/wire"
	"github.com/picfight/pfcd/picfight"
)

// TestNet3Params defines the network parameters for the test currency network.
// This network is sometimes simply called "testnet".
// This is the third public iteration of testnet.
var TestNet3Params = Params{
	Name:        "testnet3",
	Net:         wire.TestNet3,
	DefaultPort: "19708",
	DNSSeeds: []DNSSeed{
		{"testnet-seed.picfight.org", true},
	},

	// Chain parameters
	GenesisBlock:             &testNet3GenesisBlock,
	GenesisHash:              &testNet3GenesisHash,
	PowLimit:                 testNetPowLimit,
	PowLimitBits:             picfight.BigToCompact(testNetPowLimit),
	ReduceMinDifficulty:      true,
	MinDiffReductionTime:     time.Minute * 10, // ~99.3% chance to be mined before reduction
	GenerateSupported:        true,
	MaximumBlockSizes:        []int{1310720},
	MaxTxSize:                1000000,
	TargetTimePerBlock:       time.Minute * 2,
	WorkDiffAlpha:            1,
	WorkDiffWindowSize:       144,
	WorkDiffWindows:          20,
	TargetTimespan:           time.Minute * 2 * 144, // TimePerBlock * WindowSize
	RetargetAdjustmentFactor: 4,

	// Subsidy parameters.
	BaseSubsidy:              2500000000, // 25 Coin
	MulSubsidy:               100,
	DivSubsidy:               101,
	SubsidyReductionInterval: 2048,
	WorkRewardProportion:     6,
	StakeRewardProportion:    3,
	BlockTaxProportion:       1,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{},

	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationQuorum:     2520, // 10 % of RuleChangeActivationInterval * TicketsPerBlock
	RuleChangeActivationMultiplier: 3,    // 75%
	RuleChangeActivationDivisor:    4,
	RuleChangeActivationInterval:   5040, // 1 week

	// Enforce current block version once majority of the network has
	// upgraded.
	// 51% (51 / 100)
	// Reject previous block versions once a majority of the network has
	// upgraded.
	// 75% (75 / 100)
	BlockEnforceNumRequired: 51,
	BlockRejectNumRequired:  75,
	BlockUpgradeNumToCheck:  100,

	// AcceptNonStdTxs is a mempool param to either accept and relay
	// non standard txs to the network or reject them
	AcceptNonStdTxs: true,

	// Address encoding magics
	NetworkAddressPrefix: "R",
	PubKeyAddrID:         [2]byte{0x0d, 0xef}, // starts with Rk
	PubKeyHashAddrID:     [2]byte{0xc9, 0x57}, // starts with Rs
	PKHEdwardsAddrID:     [2]byte{0xc7, 0x86}, // starts with Re
	PKHSchnorrAddrID:     [2]byte{0xc5, 0xd9}, // starts with RS
	ScriptHashAddrID:     [2]byte{0xc7, 0x3f}, // starts with Rc
	PrivateKeyID:         [2]byte{0x0c, 0xdd}, // starts with Pr

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x97}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xd1}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	SLIP0044CoinType: 1,  // SLIP0044, Testnet (all coins)
	LegacyCoinType:   11, // for backwards compatibility

	// PicFight PoS parameters
	MinimumStakeDiff:        20000000, // 0.2 Coin
	TicketPoolSize:          1024,
	TicketsPerBlock:         5,
	TicketMaturity:          16,
	TicketExpiry:            6144, // 6*TicketPoolSize
	CoinbaseMaturity:        16,
	SStxChangeMaturity:      1,
	TicketPoolSizeWeight:    4,
	StakeDiffAlpha:          1,
	StakeDiffWindowSize:     144,
	StakeDiffWindows:        20,
	StakeVersionInterval:    144 * 2 * 7, // ~1 week
	MaxFreshStakePerBlock:   20,          // 4*TicketsPerBlock
	StakeEnabledHeight:      16 + 16,     // CoinbaseMaturity + TicketMaturity
	StakeValidationHeight:   768,         // Arbitrary
	StakeBaseSigScript:      []byte{0x00, 0x00},
	StakeMajorityMultiplier: 3,
	StakeMajorityDivisor:    4,

	// PicFight organization related parameters.
	// Organization address is TcrypGAcGCRVXrES7hWqVZb5oLJKCZEtoL1.
	OrganizationPkScript:        hexDecode("a914d585cd7426d25b4ea5faf1e6987aacfeda3db94287"),
	OrganizationPkScriptVersion: 0,
	BlockOneLedger:              BlockOneLedgerTestNet3,
}
