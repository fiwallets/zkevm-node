package constants

import "github.com/fiwallets/zkevm-node/state"

var (
	EffectivePercentage     = []uint8{state.MaxEffectivePercentage}
	TwoEffectivePercentages = []uint8{state.MaxEffectivePercentage, state.MaxEffectivePercentage}
)
