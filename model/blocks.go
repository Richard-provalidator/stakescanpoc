package model

import (
	"fmt"

	"gorm.io/gorm"
)

func InsertBlocks(db *gorm.DB, block Block) error {
	var existBlock []Block
	err := db.Where("height = ?", block.Height).Find(&existBlock).Error
	if err != nil {
		return fmt.Errorf("db.Where: %w", err)
	}
	if len(existBlock) != 0 {
		return fmt.Errorf("0") // 에러가 0 이면 height+1 해주기
	}
	err = db.Create(&block).Error
	if err != nil {
		return fmt.Errorf("db.Create: %w", err)
	}
	return nil
}
