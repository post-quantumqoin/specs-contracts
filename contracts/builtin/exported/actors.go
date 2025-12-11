package exported

import (
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/account"
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/cron"
	init_ "github.com/post-quantumqoin/specs-contracts/contracts/builtin/init"
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/market"
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/miner"
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/multisig"
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/paych"
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/power"
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/reward"
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/system"
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/verifreg"
	"github.com/post-quantumqoin/specs-contracts/contracts/runtime"
)

func BuiltinActors() []runtime.VMActor {
	return []runtime.VMActor{
		account.Actor{},
		cron.Actor{},
		init_.Actor{},
		market.Actor{},
		miner.Actor{},
		multisig.Actor{},
		paych.Actor{},
		power.Actor{},
		reward.Actor{},
		system.Actor{},
		verifreg.Actor{},
	}
}
