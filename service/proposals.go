package service

import (
	"context"
	"errors"
	"fmt"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/stakescanpoc/config"
	"github.com/stakescanpoc/models"
)

func GetProposalsFromRPC(ctx config.Context) ([]*v1.Proposal, error) {
	rpcClient, err := rpchttp.New(ctx.Chain.RPC, "/websocket")
	if err != nil {
		return nil, fmt.Errorf("rpchttp.New: %w", err)
	}
	c := NewClient(rpcClient)
	qc := v1.NewQueryClient(c)
	res, err := qc.Proposals(context.Background(), &v1.QueryProposalsRequest{})
	if err != nil {
		return nil, fmt.Errorf("qc.Proposals: %w", err)
	}
	fmt.Println(res.Proposals)
	return res.Proposals, nil
}

func UpdateProposals(ctx config.Context, proposal models.Proposal) error {
	err := models.InsertProposal(ctx.DB, proposal)
	if errors.Is(err, fmt.Errorf("0")) {
		err := models.UpdateProposal(ctx.DB, proposal)
		if err != nil {
			return fmt.Errorf("models.UpdateProposal: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("models.InsertProposal: %w", err)
	}

	return nil
}
