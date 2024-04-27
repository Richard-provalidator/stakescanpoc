package service

import (
	"context"
	"errors"
	"fmt"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	"github.com/provalidator/stakescan-indexer/model"
)

func GetProposalsFromRPC(ctx Context) ([]*v1.Proposal, error) {
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

func UpdateProposals(ctx Context, proposal model.Proposal) error {
	err := model.InsertProposal(ctx.DB, proposal)
	if errors.Is(err, model.ErrAlreadyExists) {
		err := model.UpdateProposal(ctx.DB, proposal)
		if err != nil {
			return fmt.Errorf("update proposal: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("insert proposal: %w", err)
	}

	return nil
}
