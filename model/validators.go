package model

import (
	"fmt"

	"gorm.io/gorm"
)

func InsertValidators(db *gorm.DB, validators []Validator) error {
	var DBvalidators []Validator
	err := db.Find(&DBvalidators).Error
	if err != nil {
		return fmt.Errorf("db.Find: %w", err)
	}
	if len(validators) != 0 {
		return fmt.Errorf("0")
	}
	err = db.Create(&validators).Error
	if err != nil {
		return fmt.Errorf("db.Create: %w", err)
	}
	return nil
}

func UpdateValidators(db *gorm.DB, validator Validator, height int64) error {
	var chg string
	err := db.Model(&validator).Where("address = ?", chg).Where("denom = ?", chg).
		Update("amount", chg).Update("height", height).Error
	if err != nil {
		return fmt.Errorf("db.Where: %w", err)
	}
	return nil
}
