package models

import (
	"fmt"
	"github.com/stakescanpoc/config"
)

func InsertBalance(ctx config.Context, balance Balance) error {

	err := ctx.DB.Create(&balance).Error
	if err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}

func FindBalance(ctx config.Context, address []byte, denom string) (Balance, error) {
	var balances []Balance
	err := ctx.DB.Where("address = ?", address).Where("denom = ?", denom).Find(&balances).Error
	if err != nil {
		return Balance{}, fmt.Errorf("DB.Where: %w", err)
	}
	if len(balances) == 0 {
		newBalance := Balance{
			Address: address,
			Denom:   denom,
			Amount:  "0",
			Height:  ctx.Height,
		}
		//err := InsertBalance(ctx, newBalance)
		//if err != nil {
		//	return Balance{}, fmt.Errorf("InsertBalance: %w", err)
		//}
		return newBalance, nil
	}
	return balances[0], nil
}

func UpdateBalance(ctx config.Context, balance Balance) error {
	var balances []Balance
	err := ctx.DB.Model(&balances).Where("address = ?", balance.Address).Where("denom = ?", balance.Denom).
		Update("amount", balance.Amount).Update("height", ctx.Height).Error
	if err != nil {
		return fmt.Errorf("DB.Where: %w", err)
	}
	return nil
}
