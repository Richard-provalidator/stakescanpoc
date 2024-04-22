package service

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
//				models.UpdateAddress(receiveAccount, receiveAmount.Add(amount).String())
//				models.UpdateAddress(sendAccount, sendAmount.Sub(amount).String())
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
//				models.UpdateAddress(sendAccount, sendAmount.Sub(amount).String())
//			}
//		}
//	}
//}

//func AmountFinder(account string) decimal.Decimal {
//	var amount decimal.Decimal
//	if account == models.Account.Address {
//		amountstr := models.Account.Amount
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
