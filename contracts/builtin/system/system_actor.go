package system

import (
	"github.com/post-quantumqoin/core-types/abi"
	"github.com/post-quantumqoin/core-types/cbor"
	"github.com/post-quantumqoin/core-types/exitcode"
	"github.com/ipfs/go-cid"

	"github.com/post-quantumqoin/specs-contracts/contracts/builtin"
	"github.com/post-quantumqoin/specs-contracts/contracts/runtime"
	"github.com/post-quantumqoin/specs-contracts/contracts/util/adt"
)

type Actor struct{}

func (a Actor) Exports() []interface{} {
	return []interface{}{
		builtin.MethodConstructor: a.Constructor,
	}
}

func (a Actor) Code() cid.Cid {
	return builtin.SystemActorCodeID
}

func (a Actor) IsSingleton() bool {
	return true
}

func (a Actor) State() cbor.Er {
	return new(State)
}

var _ runtime.VMActor = Actor{}

func (a Actor) Constructor(rt runtime.Runtime, _ *abi.EmptyValue) *abi.EmptyValue {
	rt.ValidateImmediateCallerIs(builtin.SystemActorAddr)
	st, err := ConstructState(adt.AsStore(rt))
	builtin.RequireNoErr(rt, err, exitcode.ErrIllegalState, "failed to construct state")
	rt.StateCreate(st)
	return nil
}
