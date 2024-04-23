package service

import (
	"context"
	"fmt"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
)

func GetBlockResultsFromRPC(RPC string, height int64) (*coretypes.ResultBlockResults, error) {
	rpcClient, err := rpchttp.New(RPC, "/websocket")
	if err != nil {
		return nil, fmt.Errorf("rpchttp.New: %w", err)
	}
	results, err := rpcClient.BlockResults(context.Background(), &height)
	if err != nil {
		return nil, fmt.Errorf("rpcClient.BlockResults: %w", err)
	}
	return results, nil
}
