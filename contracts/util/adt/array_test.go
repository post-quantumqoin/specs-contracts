package adt_test

import (
	"testing"

	"github.com/post-quantumqoin/address"
	"github.com/stretchr/testify/require"

	"github.com/post-quantumqoin/specs-contracts/contracts/util/adt"
	"github.com/post-quantumqoin/specs-contracts/support/mock"
)

func TestArrayNotFound(t *testing.T) {
	rt := mock.NewBuilder(address.Undef).Build(t)
	store := adt.AsStore(rt)
	arr, err := adt.MakeEmptyArray(store, 3)
	require.NoError(t, err)

	found, err := arr.Get(7, nil)
	require.NoError(t, err)
	require.False(t, found)
}
