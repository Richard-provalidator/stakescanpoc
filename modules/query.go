package modules

import (
	"context"
	"encoding/json"
	"strconv"

	"cosmossdk.io/x/circuit"
	"cosmossdk.io/x/evidence"
	"cosmossdk.io/x/upgrade"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authz "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/cosmos/cosmos-sdk/types/module/testutil"

	// gaia "github.com/cosmos/gaia/v15/app"
	"github.com/go-resty/resty/v2"

	"github.com/stakescanpoc/log"
	"github.com/stakescanpoc/models"
)

func connectHTTP(RPC, connectURL string) (*resty.Response, error) {
	client := resty.New()
	res, err := client.R().Get(RPC + connectURL)
	log.Logger.Trace.Println("RPC", RPC)

	if err != nil {
		log.Logger.Error.Println("connectRPC Failed : ", err)
		return nil, err
	}
	return res, nil
}
func GetBlockByHeightByRPC(RPC string, height int64) (*coretypes.ResultBlock, error) {
	rpcClient, err := rpchttp.New(RPC, "/websocket")
	if err != nil {
		log.Logger.Error.Println("connectRPC Failed : ", err)
		return nil, err
	}
	results, err := rpcClient.Block(context.Background(), &height)
	if err != nil {
		log.Logger.Error.Println("GetBlockByHeight Failed : ", err)
		return nil, err
	}

	return results, nil
}

func TxFinder(result *coretypes.ResultBlock) ([]map[string]interface{}, error) {
	encCfg := testutil.MakeTestEncodingConfig(
		staking.AppModuleBasic{},
		distribution.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		slashing.AppModuleBasic{},
		mint.AppModuleBasic{},
		gov.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		authz.AppModuleBasic{},
		evidence.AppModuleBasic{},
		circuit.AppModuleBasic{},
		// accounts.AppModuleBasic{},
		// feegrant.AppModuleBasic{},
		// group.AppModuleBasic{},
		// nft.AppModuleBasic{},
		// consensusparam.AppModuleBasic{},
		// pool.AppModuleBasic{},
		// epochs.AppModuleBasic{},
		// Add more modules you'd like to
	)
	txCfg := encCfg.TxConfig
	var txsJSON []map[string]interface{}
	for _, txBytes := range result.Block.Data.Txs {
		tx, err := txCfg.TxDecoder()(txBytes)
		if err != nil {
			log.Logger.Error.Println("TxDecoder Failed : ", err)
			return nil, err
		}
		txJSON, err := txCfg.TxJSONEncoder()(tx)
		if err != nil {
			log.Logger.Error.Println("TxJSONEncoder Failed : ", err)
			return nil, err
		}
		var JSON map[string]interface{}
		err = json.Unmarshal(txJSON, &JSON)
		if err != nil {
			log.Logger.Error.Println("JSON Unmarshal Failed : ", err)
			return nil, err
		}
		txsJSON = append(txsJSON, JSON)
		log.Logger.Trace.Println("JSON", JSON)
		// sdkTx := tx.(interface {
		// 	GetProtoTx() *sdktx.Tx
		// }).GetProtoTx()
		// _ = sdkTx // use it
	}
	return txsJSON, nil
}

func GetBlockResultsByRPC(RPC string, height int64) (*coretypes.ResultBlockResults, error) {
	rpcClient, err := rpchttp.New(RPC, "/websocket")
	if err != nil {
		log.Logger.Error.Println("connectRPC Failed : ", err)
		return nil, err
	}
	results, err := rpcClient.BlockResults(context.Background(), &height)
	if err != nil {
		log.Logger.Error.Println("GetBlockByHeight Failed : ", err)
		return nil, err
	}
	return results, nil
}

func EventFinder(results *coretypes.ResultBlockResults) {
	// log.Logger.Trace.Println("BlockResults", results)
	log.Logger.Trace.Println("results.FinalizeBlockEvents", results.FinalizeBlockEvents)
	log.Logger.Trace.Println("results.TxsResults", results.TxsResults)
}

func GetBlockByHeightByHTTP(chain models.ChainInfo, height int) (models.BlockJSON, error) {
	strheight := strconv.Itoa(height)
	var blockJSON models.BlockJSON
	log.Logger.Trace.Println("RPC", chain.RPC)

	res, err := connectHTTP(chain.RPC, "/block?height="+strheight)
	if err != nil {
		log.Logger.Error.Println("connectRPC Failed : ", err)
		return blockJSON, err
	}
	result := string(res.Body())
	log.Logger.Trace.Println(chain.ChainName, "GetBlockByHeight ", result)
	json.Unmarshal(res.Body(), &blockJSON)
	return blockJSON, nil
}

func GetBlockResultsByHttp(chain models.ChainInfo, height int) (models.BlockResultsJSON, error) {
	var blockResultsJSON models.BlockResultsJSON
	strheight := strconv.Itoa(height)
	res, err := connectHTTP(chain.RPC, "/block_results?height="+strheight)
	if err != nil {
		log.Logger.Error.Println("connectRPC Failed : ", err)
		return blockResultsJSON, err
	}
	result := string(res.Body())
	json.Unmarshal(res.Body(), &blockResultsJSON)
	log.Logger.Trace.Println("GetBlockResults ", result)
	return blockResultsJSON, nil
}

func GetBeginBlockEvent(chain models.ChainInfo) {

}

func GetEndBlockEvent() {

}

func GetTXEvent() {

}
