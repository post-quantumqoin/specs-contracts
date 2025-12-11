package cron

import (
	addr "github.com/post-quantumqoin/address"
	"github.com/post-quantumqoin/core-types/abi"

	"github.com/post-quantumqoin/specs-contracts/contracts/builtin"
)

type State struct {
	Entries []Entry
}

type Entry struct {
	Receiver  addr.Address  // The actor to call (must be an ID-address)
	MethodNum abi.MethodNum // The method number to call (must accept empty parameters)
}

func ConstructState(entries []Entry) *State {
	return &State{Entries: entries}
}

// The default entries to install in the cron actor's state at genesis.
func BuiltInEntries() []Entry {
	return []Entry{
		{
			Receiver:  builtin.StoragePowerActorAddr,
			MethodNum: builtin.MethodsPower.OnEpochTickEnd,
		},
		{
			Receiver:  builtin.StorageMarketActorAddr,
			MethodNum: builtin.MethodsMarket.CronTick,
		},
	}
}
