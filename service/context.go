package service

import (
	"context"

	"github.com/cosmos/gaia/v15/app/params"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/provalidator/stakescan-indexer/config"
)

type Context struct {
	context.Context
	Logger *logrus.Logger
	Chain  config.Chain
	DB     *gorm.DB
	Height int64
	EncCfg params.EncodingConfig
}
