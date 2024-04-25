package service

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
)

func FindBlockEvents(results *coretypes.ResultBlockResults) ([]abcitypes.Event, []abcitypes.Event) {
	var beginEvents []abcitypes.Event
	var endEvents []abcitypes.Event
	for _, BeginBlockEvents := range results.BeginBlockEvents {
		beginEvents = append(beginEvents, BeginBlockEvents)
	}
	for _, EndBlockEvents := range results.EndBlockEvents {
		endEvents = append(endEvents, EndBlockEvents)
	}
	return beginEvents, endEvents
}
