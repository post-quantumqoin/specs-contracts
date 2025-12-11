package system

import (
	"github.com/ipfs/go-cid"
	"github.com/post-quantumqoin/core-types/abi"
	"github.com/post-quantumqoin/core-types/cbor"

	"github.com/post-quantumqoin/specs-contracts/contracts/builtin"
	"github.com/post-quantumqoin/specs-contracts/contracts/runtime"
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

	rt.StateCreate(&State{})
	return nil
}

type State struct{}
