package system

import (
	"context"

	cid "github.com/ipfs/go-cid"
	xerrors "golang.org/x/xerrors"

	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/manifest"
	"github.com/post-quantumqoin/specs-contracts/contracts/util/adt"
)

type State struct {
	BuiltinActors cid.Cid // ManifestData
}

func ConstructState(store adt.Store) (*State, error) {
	empty, err := store.Put(context.TODO(), &manifest.ManifestData{})
	if err != nil {
		return nil, xerrors.Errorf("failed to create empty manifest: %w", err)
	}

	return &State{BuiltinActors: empty}, nil
}
