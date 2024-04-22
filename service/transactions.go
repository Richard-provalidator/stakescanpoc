package service

import (
	"context"
	"encoding/json"
	"fmt"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	gaia "github.com/cosmos/gaia/v15/app"
	"github.com/stakescanpoc/config"
	"github.com/stakescanpoc/models"
)

type Txs struct {
	Tx     *sdktx.Tx
	TxHash string
	Res    *coretypes.ResultTx
}

func QueryTx(RPC string, result *coretypes.ResultBlock) ([]Txs, error) {
	var txs []Txs
	for _, txBytes := range result.Block.Txs {
		rpcClient, err := rpchttp.New(RPC, "/websocket")
		if err != nil {
			// log.Logger.Error.Println("connectRPC Failed : ", err)
			fmt.Println("connectRPC Failed : ", err)

		}
		res, err := rpcClient.Tx(context.Background(), txBytes.Hash(), false)
		if err != nil {
			return nil, fmt.Errorf("rpcClient.Tx: %w", err)
		}
		encCfg := gaia.RegisterEncodingConfig()

		tx, err := encCfg.TxConfig.TxDecoder()(res.Tx)
		if err != nil {
			return nil, fmt.Errorf("encCfg.TxConfig.TxDecoder: %w", err)
		}
		txJSON, err := encCfg.TxConfig.TxJSONEncoder()(tx)
		if err != nil {
			return nil, fmt.Errorf("encCfg.TxConfig.TxJSONEncoder: %w", err)
		}
		var JSON sdktx.MsgResponse
		err = json.Unmarshal(txJSON, &JSON)
		if err != nil {
			// log.Logger.Error.Println("JSON Unmarshal Failed : ", err)
			fmt.Println("JSON Unmarshal Failed : ", err)
			return nil, err
		}
		fmt.Println("JSON", JSON)
		//fmt.Println(string(txJSON)) // use it
		sdkTx := tx.(interface {
			GetProtoTx() *sdktx.Tx
		}).GetProtoTx()
		//sdkTx.Body.Messages
		txHash := fmt.Sprintf("%x", txBytes.Hash())
		txs = append(txs, Txs{
			Tx:     sdkTx,
			TxHash: txHash,
			Res:    res,
		})
	}
	return txs, nil
}

func InsertTxs(ctx config.Context, txs []Txs) error {
	for i, tx := range txs {
		transaction := models.Transactions{
			TxIDX:       i,
			Code:        tx.Res.TxResult.Code,
			TxHash:      fmt.Sprintf("%x", tx.Res.Hash),
			Height:      tx.Res.Height,
			Messages:    nil,
			MessageType: tx.Tx.Body.Messages[0].TypeUrl,
			Events:      tx.Res.TxResult.Events,
			GasWanted:   tx.Res.TxResult.GasWanted,
			GasUsed:     tx.Res.TxResult.GasUsed,
		}
		err := models.InsertTxs(ctx, transaction)
		if err != nil {
			return fmt.Errorf("models.InsertTxs: %w", err)
		}
	}
	return nil
}

//
//func MakeTxJSON(txCfg client.TxConfig, txBytes cmttypes.Tx) (map[string]interface{}, error) {
//	var JSON map[string]interface{}
//	tx, err := txCfg.TxDecoder()(txBytes)
//	if err != nil {
//		// log.Logger.Error.Println("TxDecoder Failed : ", err)
//		fmt.Println("TxDecoder Failed : ", err)
//		return nil, err
//	}
//	txJSON, err := txCfg.TxJSONEncoder()(tx)
//	if err != nil {
//		// log.Logger.Error.Println("TxJSONEncoder Failed : ", err)
//		fmt.Println("TxJSONEncoder Failed : ", err)
//		return nil, err
//	}
//	err = json.Unmarshal(txJSON, &JSON)
//	if err != nil {
//		// log.Logger.Error.Println("JSON Unmarshal Failed : ", err)
//		fmt.Println("JSON Unmarshal Failed : ", err)
//		return nil, err
//	}
//	return JSON, nil
//}

// func TxFinder(result *coretypes.ResultBlock) ([]map[string]interface{}, error) {
// 	encCfg := testutil.MakeTestEncodingConfig(
// 		auth.AppModuleBasic{},
// 		authz.AppModuleBasic{},
// 		bank.AppModuleBasic{},
// 		capability.AppModuleBasic{},
// 		consensus.AppModuleBasic{},
// 		crisis.AppModuleBasic{},
// 		distribution.AppModuleBasic{},
// 		evidence.AppModuleBasic{},
// 		// feegrant.AppModuleBasic{},
// 		genutil.AppModuleBasic{},
// 		gov.AppModuleBasic{},
// 		// group.AppModuleBasic{},
// 		mint.AppModuleBasic{},
// 		// nft.AppModuleBasic{},
// 		params.AppModuleBasic{},
// 		// simulation.AppModuleBasic{},
// 		slashing.AppModuleBasic{},
// 		staking.AppModuleBasic{},
// 		upgrade.AppModuleBasic{},
// 		// accounts.AppModuleBasic{},
// 		// circuit.AppModuleBasic{},
// 		// counter.AppModuleBasic{},
// 		// epochs.AppModuleBasic{},
// 		// protocolpool.AppModuleBasic{},
// 		// tx.AppModuleBasic{},
// 		// consensusparam.AppModuleBasic{},
// 		// pool.AppModuleBasic{},
// 		// Add more modules you'd like to
// 	)
// 	txCfg := encCfg.TxConfig
// 	var txs cmttypes.Txs
// 	var txsJSON []map[string]interface{}
// 	for _, txBytes := range result.Block.Data.Txs {
// 		txs = append(txs, txBytes)
// 		JSON, err := MakeTxJSON(txCfg, txBytes)
// 		if err != nil {
// 			// log.Logger.Error.Println("MakeTxJSON Failed : ", err)
// 			fmt.Println("MakeTxJSON Failed : ", err)
// 			return nil, err
// 		}
// 		txsJSON = append(txsJSON, JSON)
// 		// log.Logger.Trace.Println("JSON", JSON)
// 		fmt.Println("JSON", JSON)
// 		// sdkTx := tx.(interface {
// 		// 	GetProtoTx() *sdktx.Tx
// 		// }).GetProtoTx()
// 		// _ = sdkTx // use it
// 	}
// 	return txsJSON, nil
// }
