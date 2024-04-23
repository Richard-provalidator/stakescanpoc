package service

import (
	"context"
	"fmt"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/stakescanpoc/models"
	"gorm.io/gorm"
)

func GetBlockByHeightFromRPC(RPC string, height int64) (*coretypes.ResultBlock, error) {
	rpcClient, err := rpchttp.New(RPC, "/websocket")
	if err != nil {
		// log.Logger.Error.Println("connectRPC Failed : ", err)
		fmt.Println("connectRPC Failed : ", err)
		return nil, err
	}
	results, err := rpcClient.Block(context.Background(), &height)
	if err != nil {
		// log.Logger.Error.Println("GetBlockByHeight Failed : ", err)
		fmt.Println("GetBlockByHeight Failed : ", err)
		return nil, err
	}

	return results, nil
}

func InsertBlock(DB *gorm.DB, block *coretypes.ResultBlock) error {
	dbBlock := models.Block{
		Height:                   block.Block.Height,
		ProposerConsensusAddress: block.Block.ProposerAddress.String(),
		Block:                    *block.Block,
		Timestamp:                block.Block.Time,
	}
	err := models.InsertBlocks(DB, dbBlock)
	if err != nil {
		return fmt.Errorf("InsertBlocks Failed : %w", err)
	}
	return nil
}
