package models

import (
	"fmt"
	"gorm.io/gorm"
)

func InsertValidators(DB *gorm.DB, validators []Validator) error {
	var DBvalidators []Validator
	err := DB.Find(&DBvalidators).Error
	if err != nil {
		return fmt.Errorf("DB.Find: %w", err)
	}
	if len(validators) != 0 {
		return fmt.Errorf("0")
	}
	err = DB.Create(&validators).Error
	if err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}

func UpdateValidators(DB *gorm.DB, validator Validator, height int64) error {
	var chg string
	err := DB.Model(&validator).Where("address = ?", chg).Where("denom = ?", chg).
		Update("amount", chg).Update("height", height).Error
	if err != nil {
		return fmt.Errorf("DB.Where: %w", err)
	}
	return nil
}
