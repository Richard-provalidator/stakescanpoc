package service

import (
	"context"
	"fmt"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/stakescanpoc/config"
	"github.com/stakescanpoc/models"
	"gorm.io/gorm"
)

func GetBlockByHeightFromRPC(ctx config.Context) (*coretypes.ResultBlock, error) {
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

func InsertBlock(DB *gorm.DB, block models.Block) error {
	err := models.InsertBlocks(DB, block)
	if err != nil {
		return fmt.Errorf("InsertBlocks Failed : %w", err)
	}
	return nil
}
