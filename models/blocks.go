package models

import (
	"fmt"
	"github.com/stakescanpoc/config"
)

func InsertBlocks(ctx config.Context, block Blocks) error {
	DB := ctx.DB
	var existBlock []Blocks
	err := DB.Where("height = ?", block.Height).Find(&existBlock).Error
	if err != nil {
		return fmt.Errorf("DB.Where: %w", err)
	}
	if len(existBlock) != 0 {
		return fmt.Errorf("0") // 에러가 0 이면 height+1 해주기
	}
	err = DB.Create(&block).Error
	if err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}
