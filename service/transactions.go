package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	gaia "github.com/cosmos/gaia/v15/app"
	"github.com/stakescanpoc/config"
	"github.com/stakescanpoc/models"
	"gorm.io/gorm"
)

type TxContext struct {
	Tx     *sdktx.Tx
	TxMsgs []sdktx.MsgResponse
	Res    *coretypes.ResultTx
}

func QueryTx(DB *gorm.DB, chain config.ChainInfo, result *coretypes.ResultBlock) ([]TxContext, error) {
	var txs []TxContext
	var txsEvents [][]abcitypes.Event
	for _, txBytes := range result.Block.Txs {
		rpcClient, err := rpchttp.New(chain.RPC, "/websocket")
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
		//txJSON, err := encCfg.TxConfig.TxJSONEncoder()(tx)
		//if err != nil {
		//	return nil, fmt.Errorf("encCfg.TxConfig.TxJSONEncoder: %w", err)
		//}
		//fmt.Println(string(txJSON)) // use it
		var txMsgs []sdktx.MsgResponse
		msgs := tx.GetMsgs()
		for _, msg := range msgs {
			var txMsg sdktx.MsgResponse
			err = json.Unmarshal(encCfg.Marshaler.MustMarshalJSON(msg), &txMsg)
			if err != nil {
				return nil, fmt.Errorf("json.Unmarshal: %w", err)
			}
			txMsgs = append(txMsgs, txMsg)
		}

		fmt.Println(fmt.Sprintf("%x", res.Hash))
		sdkTx := tx.(interface {
			GetProtoTx() *sdktx.Tx
		}).GetProtoTx()

		if res.TxResult.Code == 0 {
			err = ChangeBalance(DB, result.Block.Height, chain.Denom, res.TxResult.Events)
			if err != nil {
				return nil, fmt.Errorf("ChangeBalance: %w", err)
			}
		}

		txsEvents = append(txsEvents, res.TxResult.Events)
		txs = append(txs, TxContext{
			Tx:     sdkTx,
			TxMsgs: txMsgs,
			Res:    res,
		})

		for _, event := range res.TxResult.Events {
			for _, attr := range event.Attributes {
				if attr.Key == "spender" || attr.Key == "receiver" {
					account := models.Account{Address: attr.Value}
					err := models.InsertAccounts(DB, account)
					if errors.Is(err, fmt.Errorf("0")) {
					} else if err != nil {
						return nil, fmt.Errorf("models.InsertAccounts: %w", err)
					}
					err = InsertMapTxsAddr(DB, fmt.Sprintf("%x", res.Hash), attr.Value)
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}
	return txs, nil
}

func InsertTxs(DB *gorm.DB, txs []TxContext) error {
	for i, tx := range txs {
		transaction := models.Transaction{
			TxIDX:       i,
			Code:        tx.Res.TxResult.Code,
			TxHash:      fmt.Sprintf("%x", tx.Res.Hash),
			Height:      tx.Res.Height,
			Messages:    tx.TxMsgs,
			MessageType: tx.Tx.Body.Messages[0].TypeUrl,
			Events:      tx.Res.TxResult.Events,
			GasWanted:   tx.Res.TxResult.GasWanted,
			GasUsed:     tx.Res.TxResult.GasUsed,
		}
		err := models.InsertTxs(DB, transaction)
		if errors.Is(err, fmt.Errorf("0")) {
		} else if err != nil {
			return fmt.Errorf("models.InsertTxs: %w", err)
		}
	}
	return nil
}

func InsertMapTxsAddr(DB *gorm.DB, txHash, addr string) error {
	txID, err := models.FindTxID(DB, txHash)
	if err != nil {
		return fmt.Errorf("models.FindTxID: %w", err)
	}
	accID, err := models.FindAccID(DB, addr)
	if err != nil {
		return fmt.Errorf("models.FindAccID: %w", err)
	}
	err = models.InsertMapTxsAddr(DB, txID, accID)
	if errors.Is(err, fmt.Errorf("0")) {
	} else if err != nil {
		return fmt.Errorf("models.InsertMapTxsAddr: %w", err)
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
