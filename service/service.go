package service

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/stakescanpoc/config"
	"github.com/stakescanpoc/models"
)

func DoBlockTxEvents(ctx config.Context) error {
	block, err := GetBlockByHeightFromRPC(ctx)
	if err != nil {
		return fmt.Errorf("GetBlockByHeightFromRPC: %w", err)
	}
	proto, err := block.Block.ToProto()
	if err != nil {
		return fmt.Errorf("block.Block.ToProto: %w", err)
	}
	var blockJSON []byte
	switch ctx.Chain.ChainName {
	case "Cosmos":
		blockJSON = ctx.EncCfg.Cosmos.Marshaler.MustMarshalJSON(proto)
	default:
		return fmt.Errorf("chain is unsupported")
	}

	modelBlock := models.Block{
		Height:                   block.Block.Height,
		ProposerConsensusAddress: block.Block.ProposerAddress.String(),
		Block:                    blockJSON,
		Timestamp:                block.Block.Time,
	}
	_ = modelBlock
	//err = InsertBlock(ctx.DB, modelBlock)
	//if err != nil {
	//	return fmt.Errorf("InsertBlock: %w", err)
	//}

	txsContext, err := QueryTx(ctx, block)
	if err != nil {
		return fmt.Errorf("QueryTx: %w", err)
	}
	_ = txsContext
	//err = InsertTxs(ctx.DB, txsContext)
	//if err != nil {
	//	return fmt.Errorf("InsertTxs: %w", err)
	//}
	return nil
}

func DoBlockResultEvents(ctx config.Context) error {
	blockResults, err := GetBlockResultsFromRPC(ctx)
	if err != nil {
		return fmt.Errorf("GetBlockResultsFromRPC: %w", err)
	}
	beginEvents, endEvents := FindBlockEvents(blockResults)

	err = ChangeBalance(ctx, beginEvents)
	if err != nil {
		return fmt.Errorf("ChangeBalance(ctx, beginEvents): %w", err)
	}
	err = ChangeBalance(ctx, endEvents)
	if err != nil {
		return fmt.Errorf("ChangeBalance(ctx, endEvents): %w", err)
	}
	return nil
}

func DoValidators(ctx config.Context) error {
	validators, err := GetValidatorsByHeightFromRPC(ctx.Chain.RPC)
	if err != nil {
		return fmt.Errorf("GetValidatorsByHeightFromRPC: %w", err)
	}
	//fmt.Println(util.EncodeAddress(validators[0].ConsensusPubkey.Value))
	hexEncoded := hex.EncodeToString(validators[0].ConsensusPubkey.Value)
	fmt.Println("Hexadecimal 인코딩 결과:", hexEncoded)
	hexBytes, err := hex.DecodeString(hexEncoded)
	fmt.Println("Hexadecimal 디코딩 결과:", hexBytes, err)
	base64Encoded := base64.StdEncoding.EncodeToString(hexBytes)
	fmt.Println("base64 인코딩 결과:", base64Encoded)
	base64Decoded, err := base64.StdEncoding.DecodeString(hexEncoded)
	fmt.Println("base64 디코딩 결과:", base64Decoded, err)
	hexEncoded = hex.EncodeToString(base64Decoded)
	fmt.Println("Hexadecimal 인코딩 결과:", hexEncoded)
	hexBytes, err = hex.DecodeString(string(base64Decoded))
	fmt.Println("Hexadecimal 디코딩 결과:", hexBytes, err)

	//fmt.Println(util.EncodeBase64(validators[0].ConsensusPubkey.Value))
	//fmt.Println(validators[0].ConsensusPubkey.Value)
	//bech32ConAddr, err := util.MakeBech32Address(util.EncodeBase64(validators[0].ConsensusPubkey.Value))
	//if err != nil {
	//	return fmt.Errorf("util.MakeBech32Address(util.EncodeBase64(val.ConsensusPubkey.Value)): %w", err)
	//}
	//fmt.Println(bech32ConAddr)
	//bech32Addr, err := util.MakeBech32Address(validators[0].OperatorAddress)
	//if err != nil {
	//	return fmt.Errorf("util.MakeBech32Address(val.OperatorAddress): %w", err)
	//}
	//fmt.Println(validators[0].OperatorAddress)
	//fmt.Println(bech32Addr)

	//var modelValidators []models.Validator
	//for _, val := range validators {
	//	fmt.Println(util.EncodeBase64(val.ConsensusPubkey.Value))
	//	bech32ConAddr, err := util.MakeBech32Address(util.EncodeBase64(val.ConsensusPubkey.Value))
	//	if err != nil {
	//		return fmt.Errorf("util.MakeBech32Address(util.EncodeBase64(val.ConsensusPubkey.Value)): %w", err)
	//	}
	//	fmt.Println(bech32ConAddr)
	//	bech32Addr, err := util.MakeBech32Address(val.OperatorAddress)
	//	if err != nil {
	//		return fmt.Errorf("util.MakeBech32Address(val.OperatorAddress): %w", err)
	//	}
	//	modelValidator := models.Validator{
	//		ConsensusAddress: nil,
	//		Address:          bech32Addr,
	//		Moniker:          val.GetMoniker(),
	//		Jailed:           val.Jailed,
	//		Status:           val.Status.String(),
	//		VotingPower:      val.Tokens.String(),
	//		Tokens:           val.Tokens.String(),
	//		CommissionRate:   val.Commission.Rate.String(),
	//		Info:             nil,
	//	}
	//	modelValidators = append(modelValidators, modelValidator)
	//}
	//ctx.EncCfg.Cosmos.Marshaler.MustMarshalJSON(&validators[0])
	return nil
}

func DoProposal(ctx config.Context) error {
	proposals, err := GetProposalsFromRPC(ctx)
	if err != nil {
		return fmt.Errorf("GetProposalsFromRPC(ctx) %w", err)
	}
	for _, proposal := range proposals {
		var proposalJSON []byte
		switch ctx.Chain.ChainName {
		case "Cosmos":
			proposalJSON = ctx.EncCfg.Cosmos.Marshaler.MustMarshalJSON(proposal)
		default:
			return fmt.Errorf("chain is unsupported")
		}
		modelProposal := models.Proposal{
			ProposalID: proposal.Id,
			Proposal:   proposalJSON,
		}
		_ = modelProposal
		err := UpdateProposals(ctx, modelProposal)
		if err != nil {
			return fmt.Errorf("UpdateProposals: %w", err)
		}

	}
	return nil
}
