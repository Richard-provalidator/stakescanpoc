package models

import (
	"fmt"
	"gorm.io/gorm"
)

func InsertProposal(DB *gorm.DB, proposal Proposal) error {
	var existProposal []Proposal
	err := DB.Where("proposal_id = ?", proposal.ProposalID).Find(&existProposal).Error
	if err != nil {
		return fmt.Errorf("DB.Where: %w", err)
	}
	if len(existProposal) != 0 {
		return fmt.Errorf("0") // 에러가 0 이면 height+1 해주기
	}
	err = DB.Create(&proposal).Error
	if err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}

func UpdateProposal(DB *gorm.DB, proposal Proposal) error {
	err := DB.Where("proposal_id = ?", proposal.ProposalID).Find(&proposal).Error
	if err != nil {
		return fmt.Errorf("DB.Where: %w", err)
	}
	err = DB.Save(&proposal).Error
	if err != nil {
		return fmt.Errorf("DB.Save: %w", err)
	}
	return nil
}
