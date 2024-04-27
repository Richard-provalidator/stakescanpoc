package model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var ErrAlreadyExists = errors.New("already exists")

func InsertProposal(db *gorm.DB, proposal Proposal) error {
	var existProposal []Proposal
	err := db.Where("proposal_id = ?", proposal.ProposalID).Find(&existProposal).Error
	if err != nil {
		return fmt.Errorf("db.Where: %w", err)
	}
	if len(existProposal) != 0 {
		return ErrAlreadyExists
	}
	err = db.Create(&proposal).Error
	if err != nil {
		return fmt.Errorf("db.Create: %w", err)
	}
	return nil
}

func UpdateProposal(db *gorm.DB, proposal Proposal) error {
	err := db.Where("proposal_id = ?", proposal.ProposalID).Find(&proposal).Error
	if err != nil {
		return fmt.Errorf("db.Where: %w", err)
	}
	err = db.Save(&proposal).Error
	if err != nil {
		return fmt.Errorf("db.Save: %w", err)
	}
	return nil
}
