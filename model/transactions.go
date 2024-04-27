package model

import (
	"fmt"

	"gorm.io/gorm"
)

func InsertTx(db *gorm.DB, tx Transaction) error {
	var existTxs []Transaction
	err := db.Where("tx_hash = ? AND height = ?", tx.TxHash, tx.Height).Find(&existTxs).Error
	if err != nil {
		return fmt.Errorf("db.Where: %w", err)
	}
	if len(existTxs) != 0 {
		return fmt.Errorf("0")
	}
	err = db.Create(&tx).Error
	if err != nil {
		return fmt.Errorf("db.Create: %w", err)
	}
	return nil
}
