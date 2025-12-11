package testing

import "github.com/post-quantumqoin/core-types/abi"

func MakePID(input string) abi.PeerID {
	return abi.PeerID([]byte(input))
}
