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
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/provalidator/stakescan-indexer/model"
)

type TxContext struct {
	Tx     *sdktx.Tx
	TxMsgs datatypes.JSON
	Res    *coretypes.ResultTx
}
type TxMsg struct {
	TxMsg []byte
}

func QueryTx(ctx Context, result *coretypes.ResultBlock) ([]TxContext, error) {
	var txsContext []TxContext
	var txsEvents [][]abcitypes.Event
	for _, txBytes := range result.Block.Txs {
		rpcClient, err := rpchttp.New(ctx.Chain.RPC, "/websocket")
		if err != nil {
			// log.Logger.Error.Println("connectRPC Failed : ", err)
			fmt.Println("connectRPC Failed : ", err)

		}
		res, err := rpcClient.Tx(context.Background(), txBytes.Hash(), false)
		if err != nil {
			return nil, fmt.Errorf("rpcClient.Tx: %w", err)
		}
		tx, err := ctx.EncCfg.TxConfig.TxDecoder()(res.Tx)
		if err != nil {
			return nil, fmt.Errorf("encCfg.TxConfig.TxDecoder: %w", err)
		}
		//txJSON, err := encCfg.TxConfig.TxJSONEncoder()(tx)
		//if err != nil {
		//	return nil, fmt.Errorf("encCfg.TxConfig.TxJSONEncoder: %w", err)
		//}
		//fmt.Println(string(txJSON)) // use it
		var txMsgs []TxMsg
		msgs := tx.GetMsgs()
		for _, msg := range msgs {
			txMsg := ctx.EncCfg.Marshaler.MustMarshalJSON(msg)
			txMsgs = append(txMsgs, TxMsg{txMsg})
		}
		marshal, err := json.MarshalIndent(txMsgs, "", "    ")
		if err != nil {
			panic(err)
		}

		//fmt.Println(fmt.Sprintf("%x", res.Hash))
		sdkTx := tx.(interface {
			GetProtoTx() *sdktx.Tx
		}).GetProtoTx()

		if res.TxResult.Code == 0 {
			err = ChangeBalance(ctx, res.TxResult.Events)
			if err != nil {
				return nil, fmt.Errorf("ChangeBalance: %w", err)
			}
		}

		txsEvents = append(txsEvents, res.TxResult.Events)
		txsContext = append(txsContext, TxContext{
			Tx:     sdkTx,
			TxMsgs: marshal,
			Res:    res,
		})

		for _, event := range res.TxResult.Events {
			for _, attr := range event.Attributes {
				if attr.Key == "spender" || attr.Key == "receiver" {
					//account := model.Account{Address: attr.Value}
					//err := model.InsertAccounts(ctx.DB, account)
					//if errors.Is(err, fmt.Errorf("0")) {
					//} else if err != nil {
					//	return nil, fmt.Errorf("model.InsertAccounts: %w", err)
					//}
					//err = InsertMapTxsAddr(ctx.DB, fmt.Sprintf("%x", res.Hash), attr.Value)
					//if err != nil {
					//	return nil, err
					//}
				}
			}
		}
	}
	return txsContext, nil
}

func InsertTxs(DB *gorm.DB, txsContext []TxContext) error {
	for i, tx := range txsContext {
		transaction := model.Transaction{
			TxIdx:       i,
			Code:        tx.Res.TxResult.Code,
			TxHash:      fmt.Sprintf("%x", tx.Res.Hash),
			Height:      tx.Res.Height,
			Messages:    tx.TxMsgs,
			MessageType: tx.Tx.Body.Messages[0].TypeUrl,
			Events:      tx.Res.TxResult.Events,
			GasWanted:   tx.Res.TxResult.GasWanted,
			GasUsed:     tx.Res.TxResult.GasUsed,
		}
		err := model.InsertTx(DB, transaction)
		if errors.Is(err, fmt.Errorf("0")) {
		} else if err != nil {
			return fmt.Errorf("model.InsertTx: %w", err)
		}
	}
	return nil
}

func InsertMapTxsAddr(DB *gorm.DB, txHash, addr string) error {
	txID, err := model.FindTxID(DB, txHash)
	if err != nil {
		return fmt.Errorf("model.FindTxID: %w", err)
	}
	accID, err := model.FindAccID(DB, addr)
	if err != nil {
		return fmt.Errorf("model.FindAccID: %w", err)
	}
	err = model.InsertMapTxsAddr(DB, txID, accID)
	if errors.Is(err, fmt.Errorf("0")) {
	} else if err != nil {
		return fmt.Errorf("model.InsertMapTxsAddr: %w", err)
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
