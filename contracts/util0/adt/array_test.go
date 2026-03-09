package adt_test

import (
	"context"
	"testing"

	"github.com/post-quantumqoin/address"
	"github.com/stretchr/testify/require"

	"github.com/post-quantumqoin/specs-contracts/contracts/util/adt"
	"github.com/post-quantumqoin/specs-contracts/support/mock"
)

func TestArrayNotFound(t *testing.T) {
	rt := mock.NewBuilder(context.Background(), address.Undef).Build(t)
	store := adt.AsStore(rt)
	arr := adt.MakeEmptyArray(store)

	found, err := arr.Get(7, nil)
	require.NoError(t, err)
	require.False(t, found)
}
