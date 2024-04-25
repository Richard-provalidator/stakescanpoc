package service

import (
	"context"
	"fmt"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/types/query"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stakescanpoc/models"
	"gorm.io/gorm"
)

//type Validator struct {
//	ValidatorName       string        `json:"validatorName"`
//	ValidatorImgURL     string        `json:"validatorImgUrl"`
//	ValidatorAddress    string        `json:"validatorAddress"`
//	OperatorAddress     string        `json:"operatorAddress"`
//	ValidatorDetail     string        `json:"validatorDetail"`
//	Rank                int           `json:"rank"`
//	DelegatorCount      int           `json:"delegatorCount"`
//	Commission          Commission    `json:"commission"`
//	BondedHeight        int           `json:"bondedHeight"`
//	SelfBonded          float64       `json:"selfBonded"`
//	SelfDelegation      float64       `json:"selfDelegation"`
//	ValidatorSinceBlock int           `json:"validatorSinceBlock"`
//	Participation       Participation `json:"participation"`
//	Status              string        `json:"status"`
//	UptimePercent       float64       `json:"uptimePercent"`
//	VotingPower         VotingPower   `json:"votingPower"`
//}
//type Commission struct {
//	CommissionPercent    int `json:"commissionPercent"`
//	MaxCommissionPercent int `json:"maxCommissionPercent"`
//	MaxChangePercent     int `json:"maxChangePercent"`
//}
//type Participation struct {
//	TotalParticipation   int `json:"totalParticipation"`
//	CurrentParticipation int `json:"currentParticipation"`
//}
//type VotingPower struct {
//	DelegatedAmount    int     `json:"delegatedAmount"`
//	StakedAmount       int     `json:"stakedAmount"`
//	VotingPowerPercent float64 `json:"votingPowerPercent"`
//}

func GetValidatorsByHeightFromRPC(RPC string) ([]stakingtypes.Validator, error) {
	rpcClient, err := rpchttp.New(RPC, "/websocket")
	if err != nil {
		return nil, fmt.Errorf("rpchttp.New: %w", err)
	}
	c := NewClient(rpcClient)
	qc := stakingtypes.NewQueryClient(c)
	res, err := qc.Validators(context.Background(), &stakingtypes.QueryValidatorsRequest{
		Pagination: &query.PageRequest{Limit: 500},
	})
	if err != nil {
		return nil, fmt.Errorf("qc.Validators: %w", err)
	}

	return res.Validators, nil
}

func InsertValidators(DB *gorm.DB, validators []models.Validator) error {

	return nil
}
