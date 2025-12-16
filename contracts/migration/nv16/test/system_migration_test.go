package test

import (
	"context"
	"fmt"
	"strings"

	vm7 "github.com/post-quantumqoin/specs-contracts/support/vm"
	"github.com/post-quantumqoin/specs-contracts/support/vm7Util"

	"github.com/post-quantumqoin/core-types/rt"

	"github.com/post-quantumqoin/specs-contracts/contracts/states"
	"github.com/post-quantumqoin/specs-contracts/contracts/util/adt"

	"testing"

	"github.com/post-quantumqoin/specs-contracts/contracts/builtin"
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/exported"
	manifest8 "github.com/post-quantumqoin/specs-contracts/contracts/builtin/manifest"
	system8 "github.com/post-quantumqoin/specs-contracts/contracts/builtin/system"
	"github.com/post-quantumqoin/specs-contracts/contracts/migration/nv16"
	"github.com/post-quantumqoin/specs-contracts/support/ipld"
	vm8 "github.com/post-quantumqoin/specs-contracts/support/vm"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	mh "github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/require"
)

func makeTestManifest(t *testing.T, store cbor.IpldStore) cid.Cid {
	adtStore := adt.WrapStore(context.Background(), store)
	builder := cid.V1Builder{Codec: cid.Raw, MhType: mh.IDENTITY}

	manifestData := manifest8.ManifestData{}
	for _, name := range []string{"system", "init", "cron", "account", "storagepower", "storageminer", "storagemarket", "paymentchannel", "multisig", "reward", "verifiedregistry"} {
		codeCid, err := builder.Sum([]byte(fmt.Sprintf("fil/8/%s", name)))
		if err != nil {
			t.Fatal(err)
		}

		manifestData.Entries = append(manifestData.Entries,
			manifest8.ManifestEntry{
				Name: name,
				Code: codeCid,
			})
	}

	manifestDataCid, err := adtStore.Put(context.Background(), &manifestData)
	if err != nil {
		t.Fatal(err)
	}

	manifest := manifest8.Manifest{
		Version: 1,
		Data:    manifestDataCid,
	}

	manifestCid, err := adtStore.Put(context.Background(), &manifest)
	if err != nil {
		t.Fatal(err)
	}

	return manifestCid
}

func TestNv16Migration(t *testing.T) {
	ctx := context.Background()
	bs := ipld.NewBlockStoreInMemory()
	v := vm7.NewVMWithSingletons(ctx, t, bs)
	ctxStore := adt.WrapBlockStore(ctx, bs)
	manifestCid := makeTestManifest(t, ctxStore)
	log := nv16.TestLogger{TB: t}

	v = vm7Util.AdvanceToEpochWithCron(t, v, 200)

	startRoot := v.StateRoot()
	cache := nv16.NewMemMigrationCache()
	_, err := nv16.MigrateStateTree(ctx, ctxStore, manifestCid, startRoot, v.GetEpoch(), nv16.Config{MaxWorkers: 1}, log, cache)
	require.NoError(t, err)

	cacheRoot, err := nv16.MigrateStateTree(ctx, ctxStore, manifestCid, v.StateRoot(), v.GetEpoch(), nv16.Config{MaxWorkers: 1}, log, cache)
	require.NoError(t, err)

	noCacheRoot, err := nv16.MigrateStateTree(ctx, ctxStore, manifestCid, v.StateRoot(), v.GetEpoch(), nv16.Config{MaxWorkers: 1}, log, nv16.NewMemMigrationCache())
	require.NoError(t, err)
	require.True(t, cacheRoot.Equals(noCacheRoot))

	lookup := map[cid.Cid]rt.VMActor{}
	for _, ba := range exported.BuiltinActors() {
		lookup[ba.Code()] = ba
	}

	v8, err := vm8.NewVMAtEpoch(ctx, lookup, ctxStore, noCacheRoot, v.GetEpoch())
	require.NoError(t, err)

	stateTree, err := v8.GetStateTree()
	require.NoError(t, err)
	totalBalance, err := v8.GetTotalActorBalance()
	require.NoError(t, err)
	acc, err := states.CheckStateInvariants(stateTree, totalBalance, v8.GetEpoch()-1)
	require.NoError(t, err)
	require.True(t, acc.IsEmpty(), strings.Join(acc.Messages(), "\n"))

	actor, found, err := stateTree.GetActor(builtin.SystemActorAddr)
	require.NoError(t, err)
	require.True(t, found, "system actor not found")

	var manifest manifest8.Manifest
	err = ctxStore.Get(ctx, manifestCid, &manifest)
	require.NoError(t, err)

	err = manifest.Load(ctx, ctxStore)
	require.NoError(t, err)

	systemActorCodeCid, ok := manifest.Get("system")
	require.True(t, ok, "system actor not in manifest")
	require.Equal(t, systemActorCodeCid, actor.Code)

	var state system8.State
	err = ctxStore.Get(ctx, actor.Head, &state)
	require.NoError(t, err)
	require.Equal(t, manifest.Data, state.BuiltinActors)
}
