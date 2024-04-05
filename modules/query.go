package modules

import (
	"context"
	"encoding/json"
	"strconv"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"

	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	gaia "github.com/cosmos/gaia/v15/app"
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
func GetBlockByHeight(RPC string, height int64) (*coretypes.ResultBlock, error) {
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

func TxFinder(result *coretypes.ResultBlock) {
	encCfg := gaia.RegisterEncodingConfig()
	txCfg := encCfg.TxConfig
	for _, txBytes := range result.Block.Data.Txs {
		tx, err := txCfg.TxDecoder()(txBytes)
		if err != nil {
			log.Logger.Error.Println("TxDecoder Failed : ", err)
		}
		txJSON, err := txCfg.TxJSONEncoder()(tx)
		if err != nil {
			log.Logger.Error.Println("TxJSONEncoder Failed : ", err)
		}
		log.Logger.Trace.Println("txBytes", txBytes)
		log.Logger.Trace.Println("tx", tx)
		log.Logger.Trace.Println("txJSON", txJSON)
		sdkTx := tx.(interface {
			GetProtoTx() *sdktx.Tx
		}).GetProtoTx()
		_ = sdkTx // use it

		//// Base64 디코딩
		//decoded, err := base64.StdEncoding.DecodeString(strings.Split(sdkTx.String(), "Tx{")[0])
		//if err != nil {
		//	log.Logger.Error.Println("base64 디코딩 실패:", err)
		//}
		//
		//var decoder decode.Decoder
		//decodedTx, err := decoder.Decode(decoded)
		//if err != nil {
		//	log.Logger.Error.Println("decoder.Decode(txs) Failed : ", err)
		//}
		//log.Logger.Trace.Println("decodedTx", decodedTx)
	}

	//log.Logger.Trace.Println(len(txs))
	//
	//log.Logger.Trace.Println("string(txs)", string(txs))
	//log.Logger.Trace.Println("txs", txs)
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

func GetBlockResults(chain models.ChainInfo, height int) (models.BlockResultsJSON, error) {
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
