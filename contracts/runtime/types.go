package runtime

import (
	"github.com/post-quantumqoin/core-types/rt"
	addr "github.com/post-quantumqoin/address"
	"github.com/post-quantumqoin/core-types/abi"
)

// Concrete types associated with the runtime interface.

// Result of checking two headers for a consensus fault.
// type ConsensusFault = runtime0.ConsensusFault
type ConsensusFaultType int64

type ConsensusFault struct {
	// Address of the miner at fault (always an ID address).
	Target addr.Address
	// Epoch of the fault, which is the higher epoch of the two blocks causing it.
	Epoch abi.ChainEpoch
	// Type of fault.
	Type ConsensusFaultType
}

// type ConsensusFaultType = runtime0.ConsensusFaultType
const (
	//ConsensusFaultNone             ConsensusFaultType = 0
	ConsensusFaultDoubleForkMining ConsensusFaultType = 1
	ConsensusFaultParentGrinding   ConsensusFaultType = 2
	ConsensusFaultTimeOffsetMining ConsensusFaultType = 3
)
// const (
// 	ConsensusFaultDoubleForkMining = runtime0.ConsensusFaultDoubleForkMining
// 	ConsensusFaultParentGrinding   = runtime0.ConsensusFaultParentGrinding
// 	ConsensusFaultTimeOffsetMining = runtime0.ConsensusFaultTimeOffsetMining
// )

type VMActor = rt.VMActor
