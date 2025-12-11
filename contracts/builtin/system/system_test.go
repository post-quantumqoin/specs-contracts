package system_test

import (
	"testing"

	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/system"
	"github.com/post-quantumqoin/specs-contracts/support/mock"
)

func TestExports(t *testing.T) {
	mock.CheckActorExports(t, system.Actor{})
}
