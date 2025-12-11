package nv3

import (
	"context"

	cid "github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/post-quantumqoin/address"
	"github.com/post-quantumqoin/core-types/abi"

	miner "github.com/post-quantumqoin/specs-contracts/contracts/builtin/miner"
	"github.com/post-quantumqoin/specs-contracts/contracts/states"
)

type minerMigrator struct {
}

func (m *minerMigrator) MigrateState(ctx context.Context, store cbor.IpldStore, head cid.Cid, _ abi.ChainEpoch, _ address.Address, _ *states.Tree) (cid.Cid, error) {
	var st miner.State
	if err := store.Get(ctx, head, &st); err != nil {
		return cid.Undef, err
	}

	//  - repair broken partitions, deadline info:
	//  - fix power actor claim with any power delta

	newHead, err := store.Put(ctx, &st)
	return newHead, err
}
