package models

import (
	"fmt"
	"gorm.io/gorm"
)

func InsertAccounts(DB *gorm.DB, account Account) error {
	var accounts []Account
	err := DB.Where("address = ?", account.Address).Find(&accounts).Error
	if err != nil {
		return fmt.Errorf("DB.Where: %w", err)
	}
	if len(accounts) != 0 {
		return fmt.Errorf("0")
	}
	err = DB.Create(&account).Error
	if err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}

func FindTxID(DB *gorm.DB, txHash string) (int64, error) {
	var txs []Transaction
	err := DB.Where("tx_hash = ?", txHash).Find(&txs).Error
	if err != nil {
		return 0, fmt.Errorf("DB.Where: %w", err)
	}
	if len(txs) == 0 {
		return 0, fmt.Errorf("nothing has been inserted")
	}
	return txs[0].TxID, nil
}

func FindAccID(DB *gorm.DB, addr string) (int64, error) {
	var accounts []Account
	err := DB.Where("address = ?", addr).Find(&accounts).Error
	if err != nil {
		return 0, fmt.Errorf("DB.Where: %w", err)
	}
	if len(accounts) == 0 {
		return 0, fmt.Errorf("nothing has been inserted")
	}
	return accounts[0].AccID, nil
}

func InsertMapTxsAddr(DB *gorm.DB, txID, accID int64) error {
	var mapTransactionAddress []MapTransactionAddress
	err := DB.Where("tx_id = ?", txID).Where("acc_id = ?", accID).Find(&mapTransactionAddress).Error
	if err != nil {
		return fmt.Errorf("DB.Where: %w", err)
	}
	if len(mapTransactionAddress) != 0 {
		return fmt.Errorf("0")
	}
	mapTxsAddr := MapTransactionAddress{
		TxID:  txID,
		AccID: accID,
	}
	err = DB.Create(&mapTxsAddr).Error
	if err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}
