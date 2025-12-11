package ipld

import (
	"bytes"

	"github.com/ipfs/go-cid"
	"github.com/post-quantumqoin/core-types/abi"
	"github.com/post-quantumqoin/core-types/cbor"
)

// Marshals an object to bytes for storing in state.
func MarshalCBOR(o cbor.Marshaler) (cid.Cid, []byte, error) {
	r := bytes.Buffer{}
	err := o.MarshalCBOR(&r)
	if err != nil {
		return cid.Undef, nil, err
	}
	data := r.Bytes()
	key, err := abi.CidBuilder.Sum(data)
	if err != nil {
		return cid.Undef, nil, err
	}
	return key, data, nil
}
