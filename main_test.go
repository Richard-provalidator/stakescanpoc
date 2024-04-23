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
			logger.Error.Fatalln("GetBlockByHeightByRPC Failed: ", err)
		}
		//err = service.InsertBlock(ctx, block)
		//if err != nil {
		//	logger.Error.Fatalln("InsertBlock Failed: ", err)
		//}

		txs, err := service.QueryTx(ctx.DB, chain, block)
		if err != nil {
			logger.Error.Fatalln("QueryTx Failed: ", err)
		}
		_ = txs
		//err = service.InsertTxs(ctx, txs)
		//if err != nil {
		//	logger.Error.Fatalln("InsertTxs Failed: ", err)
		//}

		blockResults, err := service.GetBlockResultsFromRPC(chain.RPC, height)
		if err != nil {
			logger.Error.Fatalln("GetBlockResultsFromRPC Failed: ", err)
		}
		beginEvents, endEvents := service.FindBlockEvents(blockResults)
		_ = beginEvents
		_ = endEvents

		err = service.ChangeBalance(ctx.DB, height, chain.Denom, beginEvents)
		if err != nil {
			logger.Error.Fatalln("ChangeBalance Failed: ", err)
		}
		err = service.ChangeBalance(ctx.DB, height, chain.Denom, endEvents)
		if err != nil {
			logger.Error.Fatalln("ChangeBalance Failed: ", err)
		}
		//fmt.Println("beginEvents", beginEvents)
		//fmt.Println("endEvents", endEvents)
		//fmt.Println("txsResults", txsResults)
	}
}
