package service

import (
	"context"
	"fmt"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"gorm.io/gorm"

	"github.com/provalidator/stakescan-indexer/model"
)

func GetBlockByHeightFromRPC(ctx Context) (*coretypes.ResultBlock, error) {
	rpcClient, err := rpchttp.New(ctx.Chain.RPC, "/websocket")
	if err != nil {
		return nil, fmt.Errorf("rpchttp.New: %w", err)
	}
	results, err := rpcClient.Block(context.Background(), &ctx.Height)
	if err != nil {
		return nil, fmt.Errorf("rpcClient.Block: %w", err)
	}

	return results, nil
}

func InsertBlock(db *gorm.DB, block model.Block) error {
	err := model.InsertBlocks(db, block)
	if err != nil {
		return fmt.Errorf("InsertBlocks: %w", err)
	}
	return nil
}

func GetBlockResultsFromRPC(ctx Context) (*coretypes.ResultBlockResults, error) {
	rpcClient, err := rpchttp.New(ctx.Chain.RPC, "/websocket")
	if err != nil {
		return nil, fmt.Errorf("rpchttp.New: %w", err)
	}
	results, err := rpcClient.BlockResults(ctx, &ctx.Height)
	if err != nil {
		return nil, fmt.Errorf("rpcClient.BlockResults: %w", err)
	}
	return results, nil
}
