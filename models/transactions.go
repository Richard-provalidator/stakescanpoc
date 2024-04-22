package models

import (
	"fmt"
	"github.com/stakescanpoc/config"
)

func InsertTxs(ctx config.Context, tx Transactions) error {
	DB := ctx.DB
	var existTxs []Transactions
	err := DB.Where("height = ?", tx.Height).Find(&existTxs).Error
	if err != nil {
		return fmt.Errorf("DB.Where: %w", err)
	}
	if len(existTxs) != 0 {
		return fmt.Errorf("0")
	}
	err = DB.Create(&tx).Error
	if err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}
