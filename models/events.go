package models

import (
	"fmt"
	"gorm.io/gorm"
)

func InsertBalance(DB *gorm.DB, address, denom string, height int64) error {
	balance := Balance{
		Address: address,
		Denom:   denom,
		Amount:  "0",
		Height:  height,
	}
	err := DB.Create(&balance).Error
	if err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}

func FindBalance(DB *gorm.DB, address, denom string, height int64) (Balance, error) {
	var balances []Balance
	err := DB.Where("address = ?", address).Where("denom = ?", denom).Find(&balances).Error
	if err != nil {
		return Balance{}, fmt.Errorf("DB.Where: %w", err)
	}
	if len(balances) == 0 {
		err := InsertBalance(DB, address, denom, height)
		if err != nil {
			return Balance{}, fmt.Errorf("InsertBalance: %w", err)
		}
		newBalance := Balance{
			Address: address,
			Denom:   denom,
			Amount:  "0",
			Height:  height,
		}
		return newBalance, fmt.Errorf("nothing has been inserted")
	}
	return balances[0], nil
}

func UpdateBalance(DB *gorm.DB, balance Balance, height int64) error {
	var balances []Balance
	err := DB.Model(&balances).Where("address = ?", balance.Address).Where("denom = ?", balance.Denom).
		Update("amount", balance.Amount).Update("height", height).Error
	if err != nil {
		return fmt.Errorf("DB.Where: %w", err)
	}
	return nil
}
