package main_test

import (
	"fmt"
	"testing"

	"github.com/stakescanpoc/config"
)

func Test(t *testing.T) {

	ctx, _, err := config.InitConfig()
	if err != nil {
		fmt.Println(err)
	}
	ctx.Telegram.SendTelegramMsg("HI")
	fmt.Println(ctx.DirsMap)
	// // for _, chain := range ctx.Config.ChainInfos {
	// // 	height := int64(19816165)
	// // 	Blocks, err := modules.GetBlockByHeightByRPC(chain.RPC, height)
	// // 	if err != nil {
	// // 		// log.Logger.Error.Println("GetBlockByHeightByRPC Failed : ", err)
	// // 		fmt.Println("GetBlockByHeightByRPC Failed : ", err)
	// // 	}
	// // 	txs, err := modules.TxFinder(Blocks)
	// // 	if err != nil {
	// // 		// log.Logger.Error.Println("TxFinder Failed : ", err)
	// // 		fmt.Println("TxFinder Failed : ", err)
	// // 	}
	// // 	// log.Logger.Trace.Println("txs", txs)
	// // 	fmt.Println("txs", txs)
	// // 	BlockResults, err := modules.GetBlockResultsByRPC(chain.RPC, int64(19816165))
	// // 	if err != nil {
	// // 		// log.Logger.Error.Println("GetBlockResultsByRPC Failed : ", err)
	// // 		fmt.Println("GetBlockResultsByRPC Failed : ", err)
	// // 	}
	// // 	modules.EventFinder(BlockResults)
	// // }
}

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
