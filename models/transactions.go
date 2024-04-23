package models

import (
	"fmt"
	"gorm.io/gorm"
)

func InsertTxs(DB *gorm.DB, tx Transaction) error {
	var existTxs []Transaction
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
