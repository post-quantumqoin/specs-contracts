package miner_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/miner"
	"github.com/post-quantumqoin/specs-contracts/contracts/util/adt"
)

func sectorsArr(t *testing.T, store adt.Store, sectors []*miner.SectorOnChainInfo) miner.Sectors {
	sectorArr := miner.Sectors{adt.MakeEmptyArray(store)}
	require.NoError(t, sectorArr.Store(sectors...))
	return sectorArr
}
