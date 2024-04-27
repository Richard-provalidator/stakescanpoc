package service

import (
	"context"
	"fmt"
	"strconv"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Client struct {
	RPC     *rpchttp.HTTP
	grpcCdc encoding.Codec
}

func NewClient(rpcClient *rpchttp.HTTP) *Client {
	ir := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(ir)
	return &Client{
		RPC:     rpcClient,
		grpcCdc: cdc.GRPCCodec(),
	}
}

func (c *Client) Invoke(ctx context.Context, method string, args, reply any, _ ...grpc.CallOption) error {
	req, err := c.grpcCdc.Marshal(args)
	if err != nil {
		return err
	}
	path := method
	data := req
	height, _ := BlockHeightFromOutgoingContext(ctx)
	res, err := c.RPC.ABCIQueryWithOptions(ctx, path, data, rpcclient.ABCIQueryOptions{Height: height})
	if err != nil {
		return fmt.Errorf("abci query: %w", err)
	}
	if !res.Response.IsOK() {
		return status.Error(codes.Unknown, res.Response.Log) // TODO: better status code?
	}
	if err := c.grpcCdc.Unmarshal(res.Response.Value, reply); err != nil {
		return err
	}
	return nil
}

func (c *Client) NewStream(context.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("not implemented")
}

func BlockHeightFromOutgoingContext(ctx context.Context) (int64, bool) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return 0, false
	}
	vs := md.Get(grpctypes.GRPCBlockHeightHeader)
	if len(vs) == 0 {
		return 0, false
	}
	height, err := strconv.ParseInt(vs[0], 10, 64)
	if err != nil {
		return 0, false
	}
	return height, true
}

func AppendBlockHeightToOutgoingContext(ctx context.Context, height int64) context.Context {
	return metadata.AppendToOutgoingContext(ctx, grpctypes.GRPCBlockHeightHeader, fmt.Sprint(height))
}

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

//func AddrMaper(events []abcitypes.Event) {
//	for _, event := range events {
//		switch event.Type {
//		case "transfer":
//			for _, attribute := range event.Attributes {
//				var receiveAccount string
//				var sendAccount string
//				var amount decimal.Decimal
//				switch attribute.Key {
//				case "recipient":
//					receiveAccount = attribute.Value
//				case "sender":
//					sendAccount = attribute.Value
//				case "amount":
//					amount, _ = decimal.NewFromString(attribute.Value)
//				}
//				receiveAmount := AmountFinder(receiveAccount)
//				sendAmount := AmountFinder(sendAccount)
//				receiveAmount.Add(amount)
//				sendAmount.Sub(amount)
//				model.UpdateAddress(receiveAccount, receiveAmount.Add(amount).String())
//				model.UpdateAddress(sendAccount, sendAmount.Sub(amount).String())
//			}
//		case "tx":
//			for _, attribute := range event.Attributes {
//				var sendAccount string
//				var amount decimal.Decimal
//				switch attribute.Key {
//				case "fee_payer":
//					sendAccount = attribute.Value
//				case "fee":
//					amount, _ = decimal.NewFromString(attribute.Value)
//				}
//				sendAmount := AmountFinder(sendAccount)
//				model.UpdateAddress(sendAccount, sendAmount.Sub(amount).String())
//			}
//		}
//	}
//}

//func AmountFinder(account string) decimal.Decimal {
//	var amount decimal.Decimal
//	if account == model.Account.Address {
//		amountstr := model.Account.Amount
//		amount, _ = decimal.NewFromString(amountstr)
//	}
//	return amount
//}

func GetBeginBlockEvent() {

}

func GetEndBlockEvent() {

}

func GetTXEvent() {

}

// func GetBlockByHeightByHTTP(chain model.ChainInfo, height int) (model.BlockJSON, error) {
// 	strheight := strconv.Itoa(height)
// 	var blockJSON model.BlockJSON
// 	// log.Logger.Trace.Println("RPC", chain.RPC)
// 	fmt.Println("RPC", chain.RPC)

// 	res, err := connectHTTP(chain.RPC, "/block?height="+strheight)
// 	if err != nil {
// 		// log.Logger.Error.Println("connectRPC Failed : ", err)
// 		fmt.Println("connectRPC Failed : ", err)
// 		return blockJSON, err
// 	}
// 	result := string(res.Body())
// 	// log.Logger.Trace.Println(chain.Name, "GetBlockByHeight ", result)
// 	fmt.Println(chain.Name, "GetBlockByHeight ", result)
// 	json.Unmarshal(res.Body(), &blockJSON)
// 	return blockJSON, nil
// }

// func GetBlockResultsByHttp(chain model.ChainInfo, height int) (model.BlockResultsJSON, error) {
// 	var blockResultsJSON model.BlockResultsJSON
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
