package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	gaia "github.com/cosmos/gaia/v15/app"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/provalidator/stakescan-indexer/config"
	"github.com/provalidator/stakescan-indexer/log"
	"github.com/provalidator/stakescan-indexer/model"
	"github.com/provalidator/stakescan-indexer/service"
)

func defaultHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homeDir, ".stakescan")
}

func newRootCmd() *cobra.Command {
	const flagHome = "home"
	cmd := &cobra.Command{
		Use: "indexer",
		RunE: func(cmd *cobra.Command, args []string) error {
			home, _ := cmd.Flags().GetString(flagHome)

			// Initialize logger
			logFilename := filepath.Join(home, "log")
			logger, logFile, err := log.NewLogger(logFilename, logrus.DebugLevel)
			if err != nil {
				return fmt.Errorf("new logger: %w", err)
			}
			defer logFile.Close()

			// Load config
			cfgFilename := filepath.Join(home, "config", "config.yaml")
			cfg, err := config.Load(cfgFilename)
			if err != nil {
				return fmt.Errorf("load config: %w", err)
			}

			db, err := model.ConnectDB(cfg.DB)
			if err != nil {
				return fmt.Errorf("connect to db: %w", err)
			}

			height := int64(20106032)

			encCfg := gaia.RegisterEncodingConfig()
			ctx := service.Context{
				Context: context.TODO(),
				Logger:  logger,
				Chain:   cfg.Chain,
				DB:      db,
				Height:  height,
				EncCfg:  encCfg,
			}

			if err = service.IndexBlock(ctx); err != nil {
				return fmt.Errorf("index block: %w", err)
			}

			return nil
		},
	}
	cmd.Flags().String("home", defaultHomeDir(), "home directory")
	return cmd
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
