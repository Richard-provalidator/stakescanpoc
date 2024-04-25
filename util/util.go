package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/bech32"
	"net/url"
)

func EncodeBase64(src []byte) string {
	encoded := base64.StdEncoding.EncodeToString(src)
	return encoded
}

func EncodeAddress(src []byte) string {
	return fmt.Sprintf("%x", src)
}

func EncodeValPrefixAddress(src string) (string, error) {
	//lowerChainName := strings.ToLower(chainName)
	valoperPrefix := "valoper"
	hrp, decode, err := bech32.Decode(src)
	newHrp := hrp[:len(hrp)-len(valoperPrefix)]
	if err != nil {
		return "", fmt.Errorf("bech32.Decode: %w", err)
	}
	fmt.Println(fmt.Sprintf("%x", decode))
	newAddr, err := bech32.Encode(newHrp, decode)
	if err != nil {
		return "", fmt.Errorf("bech32.Encode: %w", err)
	}
	return newAddr, nil
}

func MakeBech32Address(src string) ([]byte, error) {
	_, decode, err := bech32.Decode(src)
	if err != nil {
		return nil, fmt.Errorf("bech32.Decode: %w", err)
	}
	return decode, nil
}

func EdncodeJSON(src any) map[string]interface{} {
	valJSON := make(map[string]interface{})
	marshal, err := json.MarshalIndent(src, "", "    ")
	if err != nil {
		return valJSON
	}
	err = json.Unmarshal(marshal, &valJSON)
	if err != nil {
		return valJSON
	}

	//JSON := make(map[string]interface{})
	//err = json.Unmarshal(marshal, &JSON)
	//if err != nil {
	//	return JSON
	//}
	return valJSON

	//// 파싱된 데이터를 사용하여 작업
	//moniker := parsedData["description"].(map[string]interface{})["moniker"].(string)
	//fmt.Println("Moniker:", moniker)
	//
	//operatorAddress := parsedData["operator_address"].(string)
	//fmt.Println("Operator Address:", operatorAddress)
	//
	//tokens := parsedData["tokens"].(string)
	//fmt.Println("Tokens:", tokens)
}

// Make paramstring
// URLValues 맵을 넣으면 query string으로 반환
func MakeQueryString(paramPairs url.Values) string {
	queryParams := paramPairs.Encode()
	return queryParams
}

//type Context struct {
//	ChainInfos []ChainInfo
//	Telegram   Telegram
//	DB         *gorm.DB
//	DirsMap    map[string]string
//}

// /*
// 	URL 호출
// 	ex)
// 	CallURL(URL) : 2sec (default)
// 	CallURL(URL,  10) : 10sec time out

// *
// */
// func CallURL(URL string, sec ...int) (string, error) {
// 	timeout := 2 * time.Second

// 	if len(sec) > 0 {
// 		timeout = time.Duration(sec[0]) * time.Second
// 	}
// 	client := resty.New()
// 	client.SetHeader("Accept", "application/json")
// 	client.SetTimeout(timeout) // timeout check
// 	client.SetHeaders(map[string]string{
// 		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
// 	})
// 	req, err := client.R().EnableTrace().Get(URL)
// 	if err != nil {
// 		log.Logger.Error.Println("Call URL Failed: ", err)
// 		return "", err
// 	}

// 	return string(req.Body()), nil
// }

// func SetConstant(a []models.ApiChainInfo, w []models.Wallet, p []models.Proposals) {
// 	apiChainUrls = a
// 	walletInfos = w
// 	proposalsInfos = p
// }

// // 응답 시간 측정
// func LatencyCheck(URL string) (time.Duration, string) {
// 	var out bytes.Buffer
// 	start := time.Now()

// 	res, err := CallURL(URL, 10) // timeout 3 sec

// 	if err != nil {
// 		// Time out
// 		if strings.Contains(err.Error(), "Client.Timeout exceeded") {
// 			return TIMEOUT_LATENCY, ""
// 		}
// 		// Error
// 		return BAD_LATENCY, ""
// 	}

// 	json.Indent(&out, []byte(res), "", "    ")

// 	latency := time.Since(start)
// 	// log.Logger.Trace.Println("latency : ", latency)

// 	return latency, out.String()
// }

// func SaveProposal(chainName string, proposal models.Proposal) {
// 	// models.DB.Begin()
// 	// voteOption := GetVoteCheck(chainName, proposal.ProposalID)
// 	// log.Logger.Trace.Println(proposal)
// 	sql := `
// 	INSERT INTO proposals(chain_name, proposal_id, type, title, description, status, voting_start_time, voting_end_time, update_date)
// 	VALUES(?, ?, ?, ?, ?, ?, ?, ?, NOW())
// 	ON DUPLICATE KEY UPDATE status = ?, update_date = NOW()
// 	`
// 	stmt, _ := models.DB.DB().Prepare(sql)
// 	_, err := stmt.Exec(chainName, proposal.ProposalID, proposal.Content.TypeURL, proposal.Content.Title, proposal.Content.Description, proposal.Status, proposal.VotingStartTime, proposal.VotingEndTime, proposal.Status)

// 	if err != nil {
// 		models.DB.Rollback()
// 		log.Logger.Error.Println(err)
// 		return
// 	}
// 	// log.Logger.Info.Println(chain, proposal.ProposalID, proposal.Content.TypeURL, proposal.Content.Title, proposal.Content.Description, proposal.Status, proposal.VotingStartTime, proposal.VotingEndTime, voteOption, proposal.Status, voteOption)
// 	models.DB.Begin().Commit()
// }

// func GetProposalJson(wg *sync.WaitGroup, URL string, chainName string, proposalsJson *models.ProposalsJson) {
// 	res, err := CallURL(URL, 10) // timeout 10 sec
// 	if err != nil {
// 		log.Logger.Error.Println("GetProposalJson error:", err)
// 		wg.Done()
// 		return
// 	}
// 	// log.Logger.Info.Println(res)
// 	json.Unmarshal([]byte(res), &proposalsJson)

// 	log.Logger.Trace.Println("["+chainName+"] proposalsJsonLen : ", len(proposalsJson.Proposals))

// 	// api 호출 후 해당 코인 마다 최대 10개씩 가장 최근의 프로포절을 들고온다.
// 	// 프로포절 10개를 DB에 업데이트
// 	for _, proposal := range proposalsJson.Proposals {
// 		chainName := models.GetChainNameFromDenom(proposal.TotalDeposit[0].Denom)
// 		go SaveProposal(chainName, proposal)
// 	}
// 	defer wg.Done()
// }

// // 프로포절 db에 업데이트
// func ProposalUpdate() {
// 	var wg sync.WaitGroup

// 	// DB에서 가져온 코인 api 정보 리스트 수 대로 루프( 현재 20개)
// 	for _, apiChainUrl := range apiChainUrls {
// 		var proposalsJson models.ProposalsJson
// 		proposalUrl := apiChainUrl.Url + "/cosmos/gov/v1beta1/proposals?pagination.limit=10&pagination.reverse=true"
// 		log.Logger.Trace.Println(proposalUrl)
// 		wg.Add(1)
// 		// api 호출
// 		go GetProposalJson(&wg, proposalUrl, apiChainUrl.ChainName, &proposalsJson)
// 		wg.Wait()
// 	}
// }

// // 체인의 URL 가져오기
// func GetChainUrl(chainName string) string {
// 	chainUrl := ""
// 	for _, apiChainUrl := range apiChainUrls {
// 		if apiChainUrl.ChainName == chainName {
// 			chainUrl = apiChainUrl.Url
// 			break
// 		}
// 	}
// 	return chainUrl
// }

// func SetVoteOption() {
// 	var wg sync.WaitGroup

// 	// DB에서 가져온 코인 api 정보 리스트 수 대로 루프( 현재 20개)
// 	for _, proposalsInfo := range proposalsInfos {
// 		wg.Add(1)
// 		// api 호출
// 		go GetVoteCheck(&wg, proposalsInfo.ChainName, proposalsInfo.ProposalID, proposalsInfo.Status)
// 	}
// 	wg.Wait()
// }

// // api를 통해서 보팅 여부를 db에 기록
// func GetVoteCheck(wg *sync.WaitGroup, chainName string, proposalNum string, status string) {
// 	var voteCheckJson models.VoteCheckJson
// 	voteOption := "VOTE_OPTION_UNKNOWN"
// 	voteCheckUrl := ""

// 	if status == "PROPOSAL_STATUS_VOTING_PERIOD" {
// 		var lcdVoteCheckJson models.LcdVoteCheckJson
// 		lcdUrl := GetChainUrl(chainName)
// 		validatorAddr := GetValidatorAddress(chainName)
// 		voteCheckUrl = lcdUrl + "/cosmos/gov/v1beta1/proposals/" + proposalNum + "/votes/" + validatorAddr

// 		res, err := CallURL(voteCheckUrl, 10) // timeout 10 sec
// 		if err != nil {
// 			log.Logger.Error.Println("GetVoteCheck1 error:", err)
// 			wg.Done()
// 			return
// 		}
// 		json.Unmarshal([]byte(res), &lcdVoteCheckJson)
// 		log.Logger.Trace.Println(voteCheckUrl)
// 		log.Logger.Trace.Println(lcdVoteCheckJson)

// 		// 보팅한 내역이 있으면 해당 투표 옵션을 넣어준다.
// 		if len(lcdVoteCheckJson.Vote.Options) > 0 {
// 			voteOption = lcdVoteCheckJson.Vote.Options[0].Option
// 		}
// 	} else {
// 		voteCheckUrl = "https://" + chainName + ".api.explorers.guru/api/proposals/" + proposalNum + "/votes/"
// 		res, err := CallURL(voteCheckUrl, 10) // timeout 10 sec
// 		if err != nil {
// 			log.Logger.Error.Println("GetVoteCheck2 error:", err)
// 			wg.Done()
// 			return
// 		}
// 		json.Unmarshal([]byte(res), &voteCheckJson)

// 		// 해당 주소가 보팅을 했는지 체크
// 		voteOption = GetVoteOption(chainName, voteCheckJson)
// 	}

// 	// SQL update
// 	log.Logger.Info.Println(voteCheckUrl)
// 	log.Logger.Info.Println("voteOption["+chainName+"_"+proposalNum+"]", voteOption)

// 	voteUpdateSql := `UPDATE proposals SET vote_option = ? WHERE chain_name = ? AND proposal_id = ?`
// 	stmt, _ := models.DB.DB().Prepare(voteUpdateSql)
// 	_, err2 := stmt.Exec(voteOption, chainName, proposalNum)
// 	models.DB.Begin().Commit()

// 	if err2 != nil {
// 		log.Logger.Error.Println("GetVoteCheck3 error:", err2)
// 		models.DB.Rollback()
// 		wg.Done()
// 		return
// 	}
// 	//
// 	wg.Done()
// }

// func GetVoteOptionStr(voteOption int) string {
// 	voteOptionStr := "VOTE_OPTION_UNKNOWN"
// 	//1 : Yes, 2: Abstain, 3: No, 4:NoWithVeto
// 	switch voteOption {
// 	case 1:
// 		voteOptionStr = "VOTE_OPTION_YES"
// 	case 2:
// 		voteOptionStr = "VOTE_OPTION_ABSTAIN"
// 	case 3:
// 		voteOptionStr = "VOTE_OPTION_NO"
// 	case 4:
// 		voteOptionStr = "VOTE_OPTION_NO_WITH_VETO"
// 	}
// 	return voteOptionStr
// }

// // json으로부터 해당 주소가 보팅을 했는지 체크
// func GetVoteOption(chainName string, voteCheckJson models.VoteCheckJson) string {
// 	voteOption := "VOTE_OPTION_UNKNOWN"
// 	for _, vote := range voteCheckJson {
// 		if vote.Voter == GetValidatorAddress(chainName) {
// 			voteOption = GetVoteOptionStr(vote.Options)
// 			break
// 		}
// 	}
// 	return voteOption
// }

// // 월렛 테이블에서 해당 체인의 밸리데이터 주소 가져오기
// func GetValidatorAddress(chainName string) string {
// 	address := "unknown"
// 	for _, w := range walletInfos {
// 		if w.ChainName == chainName {
// 			return w.Address
// 		}
// 	}
// 	return address
// }

// // 지연율 db에 업데이트
// func ApiLatencyUpdate(wg *sync.WaitGroup, apiChainInfo models.ApiChainInfo) {
// 	apiCheckUrl := apiChainInfo.Url + "/cosmos/base/tendermint/v1beta1/blocks/latest"
// 	latency, height, result := CosmosApiLatencyCheck(apiChainInfo.ChainName, apiCheckUrl)
// 	models.DB.Exec("UPDATE api_URLs SET latency = ?, height = ?, update_date = NOW(), result = ? WHERE URL = ?", latency.Milliseconds(), height, result, apiChainInfo.Url)
// 	wg.Done()
// }

// func ApiLatencyCheck() {
// 	var wg sync.WaitGroup

// 	// DB로 부터 api URL들을 가져온다.
// 	apiChainInfos := models.GetApiChainInfos()
// 	// log.Logger.Trace.Println("apiChainInfos", apiChainInfos)

// 	// Cosmos 계열 지연율 체크
// 	for _, apiChainInfo := range apiChainInfos {
// 		wg.Add(1)
// 		go ApiLatencyUpdate(&wg, apiChainInfo)
// 	}
// 	wg.Wait()

// }

// // Cosmos Api Latency Check
// func CosmosApiLatencyCheck(chain string, URL string) (time.Duration, int, string) {
// 	latency, jsonRes := LatencyCheck(URL)
// 	var apiBlockInfoNew models.ApiBlockInfoNew
// 	var err error
// 	var height int
// 	err = json.Unmarshal([]byte(jsonRes), &apiBlockInfoNew)
// 	height, _ = strconv.Atoi(apiBlockInfoNew.Block.Header.Height) //block height

// 	if err != nil {
// 		if latency == BAD_LATENCY {
// 			return BAD_LATENCY, 0, INVALID_RESPONSE
// 		} else {
// 			return TIMEOUT_LATENCY, 0, TIMEOUT_RESPONSE
// 		}
// 	} else {
// 		return latency, height, GOOD_RESPONSE
// 	}
// }
