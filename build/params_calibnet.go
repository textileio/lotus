// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeAssemblyHeight = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeTrustHeight = 600
const UpgradeNorwegianHeight = 114000

const UpgradeTurboHeight = 193789

const UpgradeHyperdriveHeight = 9999999

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
