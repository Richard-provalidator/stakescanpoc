package model

import (
	"fmt"

	"gorm.io/gorm"
)

func InsertBalance(db *gorm.DB, balance Balance) error {

	err := db.Create(&balance).Error
	if err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}

func FindBalance(db *gorm.DB, address []byte, denom string) (Balance, error) {
	var balance Balance
	err := db.Where("address = ? AND denom = ?", address, denom).First(&balance).Error
	if err != nil {
		return Balance{}, err
	}
	return balance, nil
}

func UpdateBalance(db *gorm.DB, balance Balance) error {
	var balances []Balance
	err := db.Model(&balances).Where("address = ?", balance.Address).Where("denom = ?", balance.Denom).
		Update("amount", balance.Amount).Update("height", balance.Height).Error
	if err != nil {
		return fmt.Errorf("DB.Where: %w", err)
	}
	return nil
}
