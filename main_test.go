package main_test

import (
	"fmt"
	"github.com/stakescanpoc/service"
	"testing"

	"github.com/stakescanpoc/config"
)

func Test(t *testing.T) {

	ctx, logger, err := config.InitContext()
	if err != nil {
		fmt.Println(err)
	}

	//ctx.Telegram.SendTelegramMsg("HI")
	for _, chain := range ctx.ChainInfos {
		height := int64(20106032)
		block, err := service.GetBlockByHeightFromRPC(chain.RPC, height)
		if err != nil {
			logger.Error.Fatalln("GetBlockByHeightByRPC Failed : ", err)
		}
		err = service.InsertBlock(ctx, block)
		if err != nil {
			logger.Error.Fatalln("InsertBlock Failed : ", err)
		}

		txs, err := service.QueryTx(chain.RPC, block)
		if err != nil {
			logger.Error.Fatalln("QueryTx Failed : ", err)
		}
		err = service.InsertTxs(ctx, txs)
		if err != nil {
			logger.Error.Fatalln("InsertTxs Failed : ", err)
		}

		//_, err = service.FindTx(block)
		//if err != nil {
		//	logger.Trace.Fatalln("service.FindTx Failed : ", err)
		//}
		//fmt.Println(txs)

		//for _, tx := range txs {
		//	dbTx := models.Transactions{
		//		TxID:        0,
		//		TxIDX:       0,
		//		Code:        tx.,
		//		TxHash:      "",
		//		Height:      0,
		//		Messages:    nil,
		//		MessageType: "",
		//		Events:      nil,
		//		GasWanted:   "",
		//		GasUsed:     "",
		//	}
		//}
		// log.Logger.Trace.Println("txs", txs)
		// 	fmt.Println("txs", txs)
		// 	BlockResults, err := modules.GetBlockResultsByRPC(chain.RPC, int64(19816165))
		// 	if err != nil {
		// 		// log.Logger.Error.Println("GetBlockResultsByRPC Failed : ", err)
		// 		fmt.Println("GetBlockResultsByRPC Failed : ", err)
		// 	}
		// 	modules.EventFinder(BlockResults)
	}
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
