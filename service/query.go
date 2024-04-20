package service

import (
	"context"
	"encoding/json"
	"fmt"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	cmttypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/shopspring/decimal"

	gaia "github.com/cosmos/gaia/v15/app"

	"github.com/stakescanpoc/models"
)

// func connectHTTP(RPC, connectURL string) (*resty.Response, error) {
// 	client := resty.New()
// 	res, err := client.R().Get(RPC + connectURL)
// 	log.Logger.Trace.Println("RPC", RPC)

//		if err != nil {
//			log.Logger.Error.Println("connectRPC Failed : ", err)
//			return nil, err
//		}
//		return res, nil
//	}

func GetBlockByHeightFromRPC(RPC string, height int64) (*coretypes.ResultBlock, error) {
	rpcClient, err := rpchttp.New(RPC, "/websocket")
	if err != nil {
		// log.Logger.Error.Println("connectRPC Failed : ", err)
		fmt.Println("connectRPC Failed : ", err)
		return nil, err
	}
	results, err := rpcClient.Block(context.Background(), &height)
	if err != nil {
		// log.Logger.Error.Println("GetBlockByHeight Failed : ", err)
		fmt.Println("GetBlockByHeight Failed : ", err)
		return nil, err
	}

	return results, nil
}

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
// 		JSON, err := txJSONMaker(txCfg, txBytes)
// 		if err != nil {
// 			// log.Logger.Error.Println("txJSONMaker Failed : ", err)
// 			fmt.Println("txJSONMaker Failed : ", err)
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

func FindTx(result *coretypes.ResultBlock) ([]map[string]interface{}, error) {
	encCfg := gaia.RegisterEncodingConfig()
	txCfg := encCfg.TxConfig
	var txs cmttypes.Txs
	var txsJSON []map[string]interface{}
	for _, txBytes := range result.Block.Data.Txs {
		txs = append(txs, txBytes)
		JSON, err := txJSONMaker(txCfg, txBytes)
		if err != nil {
			// log.Logger.Error.Println("txJSONMaker Failed : ", err)
			fmt.Println("txJSONMaker Failed : ", err)
			return nil, err
		}
		txsJSON = append(txsJSON, JSON)
		// log.Logger.Trace.Println("JSON", JSON)
		fmt.Println("JSON", JSON)
		// sdkTx := tx.(interface {
		// 	GetProtoTx() *sdktx.Tx
		// }).GetProtoTx()
		// _ = sdkTx // use it
	}
	return txsJSON, nil
}

func txJSONMaker(txCfg client.TxConfig, txBytes cmttypes.Tx) (map[string]interface{}, error) {
	var JSON map[string]interface{}
	tx, err := txCfg.TxDecoder()(txBytes)
	if err != nil {
		// log.Logger.Error.Println("TxDecoder Failed : ", err)
		fmt.Println("TxDecoder Failed : ", err)
		return nil, err
	}
	txJSON, err := txCfg.TxJSONEncoder()(tx)
	if err != nil {
		// log.Logger.Error.Println("TxJSONEncoder Failed : ", err)
		fmt.Println("TxJSONEncoder Failed : ", err)
		return nil, err
	}
	err = json.Unmarshal(txJSON, &JSON)
	if err != nil {
		// log.Logger.Error.Println("JSON Unmarshal Failed : ", err)
		fmt.Println("JSON Unmarshal Failed : ", err)
		return nil, err
	}
	return JSON, nil
}

func GetBlockResultsFromRPC(RPC string, height int64) (*coretypes.ResultBlockResults, error) {
	rpcClient, err := rpchttp.New(RPC, "/websocket")
	if err != nil {
		// log.Logger.Error.Println("connectRPC Failed : ", err)
		fmt.Println("connectRPC Failed : ", err)
		return nil, err
	}
	results, err := rpcClient.BlockResults(context.Background(), &height)
	if err != nil {
		// log.Logger.Error.Println("GetBlockByHeight Failed : ", err)
		fmt.Println("GetBlockByHeight Failed : ", err)
		return nil, err
	}
	return results, nil
}

func EventFinder(results *coretypes.ResultBlockResults) ([]abcitypes.Event, []*abcitypes.ExecTxResult) {
	// log.Logger.Trace.Println("BlockResults", results)
	var events []abcitypes.Event
	var execTxResults []*abcitypes.ExecTxResult
	for _, finalizeBlockEvent := range results.FinalizeBlockEvents {
		events = append(events, finalizeBlockEvent)
	}
	for _, txsResult := range results.TxsResults {
		execTxResults = append(execTxResults, txsResult)
		for i, event := range txsResult.Events {
			// log.Logger.Trace.Println("txsResult.Event", i, event)
			fmt.Println("txsResult.Event", i, event)
		}
	}
	return events, execTxResults
}

func AddrMaper(events []abcitypes.Event) {
	for _, event := range events {
		switch event.Type {
		case "transfer":
			for _, attribute := range event.Attributes {
				var receiveAccount string
				var sendAccount string
				var amount decimal.Decimal
				switch attribute.Key {
				case "recipient":
					receiveAccount = attribute.Value
				case "sender":
					sendAccount = attribute.Value
				case "amount":
					amount, _ = decimal.NewFromString(attribute.Value)
				}
				receiveAmount := AmountFinder(receiveAccount)
				sendAmount := AmountFinder(sendAccount)
				receiveAmount.Add(amount)
				sendAmount.Sub(amount)
				models.UpdateAddress(receiveAccount, receiveAmount.Add(amount).String())
				models.UpdateAddress(sendAccount, sendAmount.Sub(amount).String())
			}
		case "tx":
			for _, attribute := range event.Attributes {
				var sendAccount string
				var amount decimal.Decimal
				switch attribute.Key {
				case "fee_payer":
					sendAccount = attribute.Value
				case "fee":
					amount, _ = decimal.NewFromString(attribute.Value)
				}
				sendAmount := AmountFinder(sendAccount)
				models.UpdateAddress(sendAccount, sendAmount.Sub(amount).String())
			}
		}
	}
}

func AmountFinder(account string) decimal.Decimal {
	var amount decimal.Decimal
	if account == models.Account.Address {
		amountstr := models.Account.Amount
		amount, _ = decimal.NewFromString(amountstr)
	}
	return amount
}

func GetBeginBlockEvent() {

}

func GetEndBlockEvent() {

}

func GetTXEvent() {

}

// func GetBlockByHeightByHTTP(chain models.ChainInfo, height int) (models.BlockJSON, error) {
// 	strheight := strconv.Itoa(height)
// 	var blockJSON models.BlockJSON
// 	// log.Logger.Trace.Println("RPC", chain.RPC)
// 	fmt.Println("RPC", chain.RPC)

// 	res, err := connectHTTP(chain.RPC, "/block?height="+strheight)
// 	if err != nil {
// 		// log.Logger.Error.Println("connectRPC Failed : ", err)
// 		fmt.Println("connectRPC Failed : ", err)
// 		return blockJSON, err
// 	}
// 	result := string(res.Body())
// 	// log.Logger.Trace.Println(chain.ChainName, "GetBlockByHeight ", result)
// 	fmt.Println(chain.ChainName, "GetBlockByHeight ", result)
// 	json.Unmarshal(res.Body(), &blockJSON)
// 	return blockJSON, nil
// }

// func GetBlockResultsByHttp(chain models.ChainInfo, height int) (models.BlockResultsJSON, error) {
// 	var blockResultsJSON models.BlockResultsJSON
// 	strheight := strconv.Itoa(height)
// 	res, err := connectHTTP(chain.RPC, "/block_results?height="+strheight)
// 	if err != nil {
// 		// log.Logger.Error.Println("connectRPC Failed : ", err)
// 		fmt.Println("connectRPC Failed : ", err)
// 		return blockResultsJSON, err
// 	}
// 	result := string(res.Body())
// 	json.Unmarshal(res.Body(), &blockResultsJSON)
// 	// log.Logger.Trace.Println("GetBlockResults ", result)
// 	fmt.Println("GetBlockResults ", result)
// 	return blockResultsJSON, nil
// }
