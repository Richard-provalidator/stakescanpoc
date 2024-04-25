package main_test

import (
	gaia "github.com/cosmos/gaia/v15/app"
	"github.com/stakescanpoc/log"
	"github.com/stakescanpoc/service"
	"testing"

	"github.com/stakescanpoc/config"
)

func Test(t *testing.T) {

	rootPath := config.GetRootPath()
	logger, err := log.LogInit(rootPath)
	if err != nil {
		panic(err)
	}
	cfg, err := config.LoadYaml(rootPath)
	if err != nil {
		logger.Error.Fatalln("Failed loading config: ", err)
	}
	//ctx.Telegram.SendTelegramMsg("HI")
	for _, chain := range cfg.Chains {
		db, err := config.ConnectDatabase(chain.Database)
		if err != nil {
			logger.Error.Fatalln("Failed config.ConnectDatabase: ", err)
		}
		height := int64(20106032)
		var encCfg config.EncCfg
		encCfg.Cosmos = gaia.RegisterEncodingConfig()
		ctx := config.Context{
			Chain:  chain,
			DB:     db,
			Height: height,
			EncCfg: encCfg,
		}
		_ = ctx

		//err = service.DoBlockTxEvents(ctx)
		//if err != nil {
		//	logger.Error.Fatalln("Failed DoBlockTxEvents: ", err)
		//}
		//
		//err = service.DoBlockResultEvents(ctx)
		//if err != nil {
		//	logger.Error.Fatalln("Failed DoBlockResultEvents: ", err)
		//}

		err = service.DoValidators(ctx)
		if err != nil {
			logger.Error.Fatalln("Failed DoValidators: ", err)
		}

	}
}
