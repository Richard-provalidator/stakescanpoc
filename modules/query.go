package modules

import (
	"context"
	"encoding/json"
	"strconv"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cometbft/cometbft/types"
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
func GetBlockByHeight(RPC string, height int64) (*types.Block, error) {
	rpcClient, err := rpchttp.New(RPC, "/websocket")
	if err != nil {
		log.Logger.Error.Println("connectRPC Failed : ", err)
		return nil, err
	}
	res, err := rpcClient.Block(context.Background(), &height)
	if err != nil {
		log.Logger.Error.Println("GetBlockByHeight Failed : ", err)
		return nil, err
	}

	blockID := res.BlockID
	// splitBlockID := strings.Split(res.BlockID.String(), ":")
	// var data map[string]interface{}

	// // Base64 디코딩
	// Block, err := base64.StdEncoding.DecodeString(res.Block.String())
	// if err != nil {
	// 	log.Logger.Error.Println("Block base64 디코딩 실패:", err)
	// 	return nil, err
	// }

	result := res.Block
	log.Logger.Trace.Println("ID ", blockID.String())
	// log.Logger.Trace.Println("Block ", Block)
	log.Logger.Trace.Println("GetBlockByHeight ", result)
	return res.Block, nil
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
