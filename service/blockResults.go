package service

import (
	"context"
	"fmt"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
)

func GetBlockResultsFromRPC(RPC string, height int64) (*coretypes.ResultBlockResults, error) {
	rpcClient, err := rpchttp.New(RPC, "/websocket")
	if err != nil {
		// log.Logger.Error.Println("connectRPC Failed : ", err)
		fmt.Println("connectRPC Failed : ", err)
		return nil, err
	}
	results, err := rpcClient.BlockResults(context.Background(), &height)
	if err != nil {
		// log.Logger.Error.Println("GetBlockByHeight Failed : ", err)
		fmt.Println("GetBlockByHeight Failed : ", err)
		return nil, err
	}
	return results, nil
}

func EventFinder(results *coretypes.ResultBlockResults) ([]abcitypes.Event, []abcitypes.Event, []*abcitypes.ResponseDeliverTx) {
	// log.Logger.Trace.Println("BlockResults", results)
	var beginEvents []abcitypes.Event
	var endEvents []abcitypes.Event
	var execTxResults []*abcitypes.ResponseDeliverTx
	for _, BeginBlockEvents := range results.BeginBlockEvents {
		beginEvents = append(beginEvents, BeginBlockEvents)
	}
	for _, EndBlockEvents := range results.EndBlockEvents {
		endEvents = append(endEvents, EndBlockEvents)
	}
	for _, txsResult := range results.TxsResults {
		execTxResults = append(execTxResults, txsResult)
		for i, event := range txsResult.Events {
			// log.Logger.Trace.Println("txsResult.Event", i, event)
			fmt.Println("txsResult.Event", i, event)
		}
	}
	return beginEvents, endEvents, execTxResults
}
