package model

import (
	"fmt"

	"gorm.io/gorm"
)

func InsertAccounts(db *gorm.DB, account Account) error {
	var accounts []Account
	err := db.Where("address = ?", account.Address).Find(&accounts).Error
	if err != nil {
		return fmt.Errorf("db.Where: %w", err)
	}
	if len(accounts) != 0 {
		return fmt.Errorf("0")
	}
	err = db.Create(&account).Error
	if err != nil {
		return fmt.Errorf("db.Create: %w", err)
	}
	return nil
}

func FindTxID(db *gorm.DB, txHash string) (int64, error) {
	var txs []Transaction
	err := db.Where("tx_hash = ?", txHash).Find(&txs).Error
	if err != nil {
		return 0, fmt.Errorf("db.Where: %w", err)
	}
	if len(txs) == 0 {
		return 0, fmt.Errorf("nothing has been inserted")
	}
	return txs[0].TxID, nil
}

func FindAccID(db *gorm.DB, addr string) (int64, error) {
	var accounts []Account
	err := db.Where("address = ?", addr).Find(&accounts).Error
	if err != nil {
		return 0, fmt.Errorf("db.Where: %w", err)
	}
	if len(accounts) == 0 {
		return 0, fmt.Errorf("nothing has been inserted")
	}
	return accounts[0].AccID, nil
}

func InsertMapTxsAddr(db *gorm.DB, txID, accID int64) error {
	var mapTransactionAddress []MapTransactionAddress
	err := db.Where("tx_id = ?", txID).Where("acc_id = ?", accID).Find(&mapTransactionAddress).Error
	if err != nil {
		return fmt.Errorf("db.Where: %w", err)
	}
	if len(mapTransactionAddress) != 0 {
		return fmt.Errorf("0")
	}
	mapTxsAddr := MapTransactionAddress{
		TxID:  txID,
		AccID: accID,
	}
	err = db.Create(&mapTxsAddr).Error
	if err != nil {
		return fmt.Errorf("db.Create: %w", err)
	}
	return nil
}
