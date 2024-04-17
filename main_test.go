package main_test

import (
	"fmt"
	"testing"

	"github.com/stakescanpoc/util"
)

func TestMain(m *testing.M) {
	ctx := util.NewContext()
	ctx.InitContext()
	fmt.Println(ctx)
	// for _, chain := range models.Config.ChainInfo {
	// 	Blocks, err := modules.GetBlockByHeightByRPC(chain.RPC, int64(19816165))
	// 	if err != nil {
	// 		log.Logger.Error.Println("GetBlockByHeightByRPC Failed : ", err)
	// 	}
	// 	txs, err := modules.TxFinder(Blocks)
	// 	if err != nil {
	// 		log.Logger.Error.Println("TxFinder Failed : ", err)
	// 	}
	// 	log.Logger.Trace.Println("txs", txs)
	// 	BlockResults, err := modules.GetBlockResultsByRPC(chain.RPC, int64(19816165))
	// 	if err != nil {
	// 		log.Logger.Error.Println("GetBlockResultsByRPC Failed : ", err)
	// 	}
	// 	modules.EventFinder(BlockResults)
	// }
}
