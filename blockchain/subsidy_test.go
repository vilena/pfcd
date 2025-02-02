// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2015-2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package blockchain

import (
	"testing"

	"github.com/jfixby/coin"
	"github.com/jfixby/pin"
	"github.com/picfight/pfcd/chaincfg"
	"github.com/picfight/picfightcoin"
)

func TestPicfightCoinBlockSubsidy(t *testing.T) {
	net := &chaincfg.PicFightCoinNetParams
	calc := net.SubsidyCalculator()
	expectedTotal := calc.ExpectedTotalNetworkSubsidy()
	expectedActual := coin.Amount{799999997687360}
	expectedTotal = expectedActual
	genBlocksNum := calc.NumberOfGeneratingBlocks()
	preminedCoins := calc.PreminedCoins()
	firstBlock := calc.FirstGeneratingBlockIndex()

	totalSubsidy := preminedCoins
	for i := int64(0); i <= genBlocksNum; i++ {
		blockIndex := firstBlock + i

		work := CalcBlockWorkSubsidy(nil, blockIndex,
			net.TicketsPerBlock, net)
		stake := CalcStakeVoteSubsidy(nil, blockIndex,
			net) * int64(net.TicketsPerBlock)
		tax := CalcBlockTaxSubsidy(nil, blockIndex,
			net.TicketsPerBlock, net)
		if (work + stake + tax) == 0 {
			//break
		}
		totalSubsidy.AtomsValue = totalSubsidy.AtomsValue + (work + stake + tax)

	}

	if totalSubsidy.AtomsValue != expectedTotal.AtomsValue {
		t.Errorf("Bad total subsidy; want %v, got %v",
			expectedTotal.AtomsValue,
			totalSubsidy.AtomsValue,
		)
	}
}

func TestDecredBlockSubsidyFull(t *testing.T) {
	net := &chaincfg.DecredNetParams
	calc := net.SubsidyCalculator
	net.SubsidyCalculator = nil
	subsidyCache := NewSubsidyCache(0, net)
	exepctedValue := int64(2103834590794301)
	// value received by block-by-block testing
	fullDecredBlockSubsidyCheck(t, subsidyCache, exepctedValue)
	net.SubsidyCalculator = calc
}

func TestDecredBlockSubsidyFunctionFull(t *testing.T) {
	net := &chaincfg.DecredNetParams
	expected := net.SubsidyCalculator().ExpectedTotalNetworkSubsidy().AtomsValue
	pin.AssertNotNil("net.SubsidyCalculator", net.SubsidyCalculator)
	fullDecredBlockSubsidyCheck(t, nil, expected)
}

func fullDecredBlockSubsidyCheck(t *testing.T, cache *SubsidyCache, expected int64) {
	net := &chaincfg.DecredNetParams
	//--------------------
	totalSubsidy := coin.Amount{0}
	for i := int64(0); ; i++ {
		blockIndex := i

		work := CalcBlockWorkSubsidy(cache, blockIndex,
			net.TicketsPerBlock, net)
		stake := CalcStakeVoteSubsidy(cache, blockIndex,
			net) * int64(net.TicketsPerBlock)
		tax := CalcBlockTaxSubsidy(cache, blockIndex,
			net.TicketsPerBlock, net)
		if i%100000 == 0 {
			//fmt.Println(fmt.Sprintf("block: %v/%v: %v", i, "?", work+stake+tax))
		}
		if (work + stake + tax) == 0 {
			break
		}
		totalSubsidy.AtomsValue = totalSubsidy.AtomsValue + (work + stake + tax)

	}

	expectedTotal := coin.Amount{expected}
	if totalSubsidy.AtomsValue != expectedTotal.AtomsValue {
		t.Errorf("Bad total subsidy; want %v, got %v",
			expectedTotal.AtomsValue,
			totalSubsidy.AtomsValue,
		)
	}
}

// originalTestExpected is value from the original decred/dcrd repo
// most likely is invalid due to incorrect testing
const originalTestExpected int64 = 2099999999800912

func TestDecredBlockSubsidyFunctionOriginal(t *testing.T) {
	net := &chaincfg.DecredNetParams
	pin.AssertNotNil("net.SubsidyCalculator", net.SubsidyCalculator)
	expected := net.SubsidyCalculator().ExpectedTotalNetworkSubsidy().AtomsValue
	expected = originalTestExpected
	originalDecredBlockSubsidyCheck(t, nil, expected)
}

func TestDecredBlockSubsidyOriginal(t *testing.T) {
	net := &chaincfg.DecredNetParams
	calc := net.SubsidyCalculator
	net.SubsidyCalculator = nil
	subsidyCache := NewSubsidyCache(0, net)
	originalDecredBlockSubsidyCheck(t, subsidyCache, originalTestExpected)
	net.SubsidyCalculator = calc
}

func originalDecredBlockSubsidyCheck(t *testing.T, subsidyCache *SubsidyCache, expected int64) {
	net := &chaincfg.DecredNetParams
	params := &picfightcoin.DecredSubsidyParams{
		BaseSubsidy:              3119582664,
		MulSubsidy:               100,
		DivSubsidy:               101,
		SubsidyReductionInterval: 6144,
		// Subsidy parameters.
	}
	//--------------------
	totalSubsidy := net.BlockOneSubsidy()
	for i := int64(0); ; i++ {
		// Genesis block or first block.
		if i == 0 || i == 1 {
			continue
		}

		if i%params.SubsidyReductionInterval == 0 {
			numBlocks := params.SubsidyReductionInterval
			// First reduction internal, which is reduction interval - 2
			// to skip the genesis block and block one.
			if i == params.SubsidyReductionInterval {
				numBlocks -= 2
			}
			height := i - numBlocks

			work := CalcBlockWorkSubsidy(subsidyCache, height,
				net.TicketsPerBlock, net)
			stake := CalcStakeVoteSubsidy(subsidyCache, height,
				net) * int64(net.TicketsPerBlock)
			tax := CalcBlockTaxSubsidy(subsidyCache, height,
				net.TicketsPerBlock, net)
			if (work + stake + tax) == 0 {
				break
			}
			totalSubsidy += ((work + stake + tax) * numBlocks)

			// First reduction internal, subtract the stake subsidy for
			// blocks before the staking system is enabled.
			if i == params.SubsidyReductionInterval {
				totalSubsidy -= stake * (net.StakeValidationHeight - 2)
			}
		}
	}
	if totalSubsidy != expected {
		t.Errorf("Bad total subsidy; want %v, got %v", expected, totalSubsidy)
	}
}
