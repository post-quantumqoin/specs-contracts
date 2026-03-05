package system_test

import (
	"testing"

	"github.com/post-quantumqoin/specs-contracts/contracts/builtin"
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/manifest"
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/system"
	"github.com/post-quantumqoin/specs-contracts/support/mock"
)

func TestExports(t *testing.T) {
	mock.CheckActorExports(t, system.Actor{})
}

func TestConstruction(t *testing.T) {
	rt := mock.NewBuilder(builtin.SystemActorAddr).Build(t)
	a := system.Actor{}

	rt.ExpectValidateCallerAddr(builtin.SystemActorAddr)
	rt.SetCaller(builtin.SystemActorAddr, builtin.SystemActorCodeID)
	rt.Call(a.Constructor, nil)
	rt.Verify()

	var st system.State
	rt.GetState(&st)

	var manifestData manifest.ManifestData
	if ok := rt.StoreGet(st.BuiltinActors, &manifestData); !ok {
		t.Fatal("missing manifest data")
	}

	if len(manifestData.Entries) != 0 {
		t.Fatal("expected empty manifest data")
	}
}
