package miner_test

import (
	"testing"

	"github.com/post-quantumqoin/core-types/abi"
	"github.com/stretchr/testify/assert"

	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/miner"
)

func TestProvingPeriodDeadlines(t *testing.T) {

	t.Run("quantization spec rounds to the next deadline", func(t *testing.T) {
		periodStart := abi.ChainEpoch(2)
		curr := periodStart + miner.WPoStProvingPeriod
		d := miner.NewDeadlineInfo(periodStart, 10, curr)
		quant := miner.QuantSpecForDeadline(d)
		assert.Equal(t, d.NextNotElapsed().Last(), quant.QuantizeUp(curr))
	})
}
