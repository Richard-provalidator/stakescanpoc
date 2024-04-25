package service

import (
	"context"
	"fmt"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/stakescanpoc/config"
)

func GetBlockResultsFromRPC(ctx config.Context) (*coretypes.ResultBlockResults, error) {
	rpcClient, err := rpchttp.New(ctx.Chain.RPC, "/websocket")
	if err != nil {
		return nil, fmt.Errorf("rpchttp.New: %w", err)
	}
	results, err := rpcClient.BlockResults(context.Background(), &ctx.Height)
	if err != nil {
		return nil, fmt.Errorf("rpcClient.BlockResults: %w", err)
	}
	return results, nil
}
