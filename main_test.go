package main_test

import (
	"testing"

	"github.com/stakescanpoc/models"
	"github.com/stakescanpoc/modules"
	"github.com/stakescanpoc/util"
)

func TestMain(t *testing.T) {
	util.Init()
	for _, chain := range models.Config.ChainInfo {
		modules.GetBlockByHeight(chain.RPC, int64(19816165))
	}
}
