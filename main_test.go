package main_test

import (
	"testing"

	"github.com/stakescanpoc/log"
	"github.com/stakescanpoc/models"
	"github.com/stakescanpoc/modules"
	"github.com/stakescanpoc/util"
)

func TestMain(t *testing.T) {
	util.Init()
	for _, chain := range models.Config.ChainInfo {
		result, err := modules.GetBlockByHeight(chain.RPC, int64(19816165))
		if err != nil {
			log.Logger.Error.Println("GetBlockByHeight Failed : ", err)
		}
		modules.TxFinder(result)
	}
}
