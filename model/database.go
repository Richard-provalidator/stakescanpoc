package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/provalidator/stakescan-indexer/config"
)

// import (
// 	"github.com/provalidator/stakescan-indexer/log"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

func ConnectDB(d config.DB) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{
		DSN: d.DSN(),
		// default size for string fields
		DefaultStringSize: 256,
		// disable datetime precision, which not supported before MySQL 5.6
		DisableDatetimePrecision: true,
		// drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameIndex: true,
		// `change` when rename column, rename column not supported before MySQL 8, MariaDB
		DontSupportRenameColumn: true,
		// auto configure based on currently MySQL version
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
}

func GetAddress() {
	//현재 있는 모든 어드레스 가져옴
}

func UpdateAddress(address, amount string) {
	//addr 없으면 추가

	//addr 있으면 업데이트
}

// // 밸리데이터 정보를 db에 저장하는 펑션
// func InsertValidatorsInfo(validators Validators) error {
// 	var todayvalidators []Validators
// 	today := time.Now().Local().AddDate(0, 0, 0).Format("2006-01-02")
// 	err := DB.Where("DATE(reg_date) = ?", today).Find(&todayvalidators).Error
// 	if err != nil {
// 		log.Logger.Error.Println("Query Validators Info Create failed: ", err)
// 		return err
// 	}
// 	if len(todayvalidators) != 0 {
// 		log.Logger.Info.Println(validators.Name, " is already have validators in DB, so it doesn't update to DB")
// 		return nil
// 	}
// 	err = DB.Create(&validators).Error

// 	if err != nil {
// 		log.Logger.Error.Println("Insert Validators Info Create failed: ", err)
// 		return err
// 	}

// 	return nil
// }

// // 밸리데이터 정보를 db에 저장하는 펑션
// func InsertValidatorsInfo2(validators []Validators) error {
// 	err := DB.Create(&validators).Error

// 	if err != nil {
// 		log.Logger.Error.Println("Insert Validators Info Create failed: ", err)
// 		return err
// 	}

// 	return nil
// }

// // 딜리게이터 정보를 db에 저장하는 펑션
// func InsertDelegatorsInfo(delegators []Delegators) error {
// 	var todaydelegators []Delegators
// 	today := time.Now().Local().AddDate(0, 0, 0).Format("2006-01-02")
// 	err := DB.Where("DATE(reg_date) = ?", today).Find(&todaydelegators).Error
// 	if err != nil {
// 		log.Logger.Error.Println("Query Delegators Info Create failed: ", err)
// 		return err
// 	}
// 	if len(todaydelegators) != 0 {
// 		log.Logger.Info.Println(delegators[0].Name, " is already have validators in DB, so it doesn't update to DB")
// 		return nil
// 	}
// 	// 트랜잭션 시작
// 	tx := DB.Begin()

// 	if tx.Error != nil {
// 		log.Logger.Error.Fatal("Insert Delegators Info DB tx Begin failed: ", tx.Error)
// 	}

// 	// 인서트
// 	err = tx.CreateInBatches(&delegators, 200).Error

// 	// 인서트 도중 에러 발생시 롤백
// 	if err != nil {
// 		tx.Rollback()
// 		log.Logger.Error.Println("Insert Delegators Info Create In Batches failed, , Rollback DB: ", err)
// 		return err
// 	}

// 	tx.Commit()
// 	return nil

// }

// // 클레임 정보를 db에 저장하는 펑션
// func InsertClaim(claims Claims) {
// 	DB.Create(&claims)
// }

// func UpdateClaim(claims Claims) {

// }

// // send 테이블에 sum을 한 송금 정보를 db에 저장하는 펑션
// func InsertSends(delegators []Delegators) error {

// 	// 트랜잭션 시작
// 	tx := DB.Begin()

// 	if tx.Error != nil {
// 		log.Logger.Error.Fatal("Insert Sends DB tx Begin failed: ", tx.Error)

// 	}

// 	var sends []Sends
// 	for _, delegator := range delegators {
// 		sends = append(sends, Sends{Name: delegator.Name, Address: delegator.Address, PaybackAmount: delegator.PaybackAmount, Denom: delegator.Denom, RegDate: time.Now().Local()})
// 		log.Logger.Trace.Println("Delegator: ", delegator)
// 	}

// 	err := tx.CreateInBatches(&sends, 200).Error
// 	// 인서트 도중 에러 발생시 롤백
// 	if err != nil {
// 		tx.Rollback()
// 		log.Logger.Error.Println("Insert Sends Create In Batches failed, , Rollback DB: ", err)
// 		return err
// 	}

// 	tx.Commit()

// 	err2 := UpdateDelegatorsCalc()

// 	if err2 != nil {
// 		log.Logger.Error.Println("Insert Sends Update Delegators Calc failed: ", err2)
// 		return err2
// 	}
// 	return nil
// }

// func GetYesterdayValidator() ([]Validators, error) {
// 	var validators []Validators
// 	yesterday := time.Now().Local().AddDate(0, 0, -1).Format("2006-01-02")
// 	log.Logger.Info.Println("yesterday: ", yesterday)
// 	//DATE_FORMAT(reg_date, '%Y-%m-%d')='2023-11-03'
// 	err := DB.Where("DATE(reg_date) = ?", yesterday).Find(&validators).Error
// 	if err != nil {
// 		log.Logger.Error.Println("Get Yesterday Validators List failed: ", err)
// 		return nil, err
// 	}
// 	if len(validators) == 0 {
// 		return nil, nil
// 	}
// 	log.Logger.Info.Println("GetYesterdayValidator : ", validators)
// 	return validators, nil
// }

// // 딜리게이터에서 합쳐야 할 딜리게이터들의 주소를 가져오기
// func GetSumDelegators() []Delegators {
// 	var delegators []Delegators

// 	DB.Select("chain_name, address, denom, sum(payback_amount) AS payback_amount").Group("chain_name, address, denom").Where("calc = 0").Find(&delegators)
// 	return delegators
// }

// // 셀렉트 해서 샌드 테이블에 들어간 딜리게이터스들의 calc를 1로 바꾼다.
// func UpdateDelegatorsCalc() error {
// 	var delegators Delegators
// 	err := DB.Model(&delegators).Where("calc = ?", 0).Update("calc", 1).Error
// 	if err != nil {
// 		log.Logger.Error.Println("Get Update Delegators Calc Info failed: ", err)
// 		return err
// 	}
// 	return nil
// }

// // 모든 딜리게이터 정보를 가져오기
// func GetDelegators() {
// 	var delegators []Delegators
// 	DB.Find(&delegators)
// }

// // 모든 샌드리스트 정보를 가져오기
// func GetSendList() ([]Sends, error) {
// 	var sends []Sends
// 	err := DB.Where("txhash = ''").Find(&sends).Error
// 	if err != nil {
// 		log.Logger.Error.Println("Get Send List failed: ", err)
// 		return sends, err
// 	}
// 	return sends, nil
// }

// func GetResendList() ([]Sends, error) {
// 	var sends []Sends
// 	err := DB.Where("txhash IS NOT NULL AND is_tx_confirm = ?", 0).Find(&sends).Error
// 	if err != nil {
// 		log.Logger.Error.Println("Get Send List failed: ", err)
// 		return sends, err
// 	}
// 	return sends, nil
// }

// func GetTxResultList() ([]SendResults, []Claims, error) {
// 	var sends []Sends
// 	var claims []Claims
// 	var sendResults []SendResults
// 	err := DB.Select("chain_name, txhash, begin_height").Where("txhash IS NOT NULL AND is_tx_confirm = ?", 0).Group("chain_name, txhash, begin_height").Find(&sends).Error
// 	if err != nil {
// 		log.Logger.Error.Println("Get Send List failed: ", err)
// 		return sendResults, claims, err
// 	}
// 	for _, send := range sends {
// 		sendResults = append(sendResults, SendResults{Name: send.Name, Txhash: send.Txhash, BeginHeight: send.BeginHeight})
// 	}
// 	err = DB.Where("txhash is not null").Where("is_tx_confirm = ?", 0).Find(&claims).Error
// 	if err != nil {
// 		log.Logger.Error.Println("Get Send List failed: ", err)
// 		return sendResults, claims, err
// 	}
// 	return sendResults, claims, nil
// }

// func GetClaims() []Claims {
// 	var claims []Claims
// 	DB.Find(&claims)
// 	return claims
// }

// func UpdateClaimTx(txhash, amount string) error {
// 	var claims []Claims
// 	// 트랜잭션 시작
// 	tx := DB.Begin()

// 	if tx.Error != nil {
// 		log.Logger.Error.Println("Update ClaimTX DB tx Begin failed: ", tx.Error)
// 	}

// 	err := DB.Model(&claims).Where("txhash = ?", txhash).Update("is_tx_confirm", 1).Update("amount", amount).Error
// 	// 인서트 도중 에러 발생시 롤백
// 	if err != nil {
// 		tx.Rollback()
// 		log.Logger.Error.Println("Update ClaimTx Tx failed, Rollback DB: ", err)
// 		return err
// 	}

// 	tx.Commit()
// 	return nil
// }

// func UpdateReclaimTx(chainName, txhash string, height int64) error {
// 	var claims []Claims
// 	// 트랜잭션 시작
// 	tx := DB.Begin()

// 	if tx.Error != nil {
// 		log.Logger.Error.Println("Update ClaimTX DB tx Begin failed: ", tx.Error)
// 	}

// 	err := DB.Model(&claims).Where("chain_name = ?", chainName).Where("is_tx_confirm = ?", 0).Update("txhash", txhash).Update("begin_height", height).Error
// 	// 인서트 도중 에러 발생시 롤백
// 	if err != nil {
// 		tx.Rollback()
// 		log.Logger.Error.Println("Update ClaimTx Tx failed, Rollback DB: ", err)
// 		return err
// 	}

// 	tx.Commit()
// 	return nil
// }

// // 전송이 완료된 지갑들의 주소 배열을 가지고 와서 txhash값을 업데이트해줌.
// func UpdateSendTx(addrList []string, txhash string, height int64) error {
// 	var sends Sends
// 	var addrStr string
// 	for _, addr := range addrList {
// 		addrStr += "'" + addr + "',"
// 	}
// 	// 'str1','str2'이런식으로 나옴.
// 	addrStr = addrStr[:len(addrStr)-1]

// 	// is_tx_confirm 이 0이고, address가 addrList배열에 있는 값들만 is_tx_confirm 변화 시켜준다.
// 	//where := "address IN (" + addrStr + ")"

// 	// 트랜잭션 시작
// 	tx := DB.Begin()

// 	if tx.Error != nil {
// 		log.Logger.Error.Fatal("Update Send DB tx Begin failed: ", tx.Error)
// 	}

// 	// address가 addrList배열에 있는 값들만 txhash를 변화 시켜준다.
// 	err := DB.Model(&sends).Where("address IN ?", addrList).Where("is_tx_confirm = ?", 0).Update("txhash", txhash).Update("begin_height", height).Error
// 	//err := DB.Model(&sends).Where(where).Update("txhash", txhash).Update("begin_height", height).Error
// 	// 인서트 도중 에러 발생시 롤백
// 	if err != nil {
// 		tx.Rollback()
// 		log.Logger.Error.Println("Update Send Tx failed, Rollback DB: ", err)
// 		return err
// 	}

// 	tx.Commit()
// 	return nil
// }

// func UpdateSendTXConfirm(txhash string) error {
// 	var sends Sends

// 	// 트랜잭션 시작
// 	tx := DB.Begin()

// 	if tx.Error != nil {
// 		log.Logger.Error.Println("Update Send DB tx Begin failed: ", tx.Error)
// 		return tx.Error
// 	}

// 	err := DB.Model(&sends).Where("txhash = ?", txhash).Update("is_tx_confirm", 1).Error
// 	// 인서트 도중 에러 발생시 롤백
// 	if err != nil {
// 		tx.Rollback()
// 		log.Logger.Error.Println("Update Send Tx failed, Rollback DB: ", err)
// 		return err
// 	}

// 	tx.Commit()
// 	return nil

// }

// func UpdateClaimTXConfirm(txhash string) error {
// 	var claims Claims

// 	// 트랜잭션 시작
// 	tx := DB.Begin()

// 	if tx.Error != nil {
// 		log.Logger.Error.Println("Update Claim DB tx Begin failed: ", tx.Error)
// 		return tx.Error
// 	}

// 	err := DB.Model(&claims).Where("txhash = ?", txhash).Update("is_tx_confirm", 1).Error
// 	// 인서트 도중 에러 발생시 롤백
// 	if err != nil {
// 		tx.Rollback()
// 		log.Logger.Error.Println("Update Claim Tx failed, Rollback DB: ", err)
// 		return err
// 	}

// 	tx.Commit()
// 	return nil

// }

// // 실패했을 때 addr를 통해 다시 Send 불러오기
// func GetSendByAddr(addrList []string) ([]Sends, error) {
// 	var sends []Sends
// 	if err := DB.Where("address IN ?", addrList).Find(&sends).Error; err != nil {
// 		return nil, err
// 	}
// 	return sends, nil
// }

// // 체인 api 정보 가져오기
// func GetApiChainNames() []ApiChainName {
// 	var apiChainNames []ApiChainName
// 	sql := `SELECT chain_name FROM cosmos.api_urls GROUP BY chain_name`
// 	DB.Raw(sql).Scan(&apiChainNames)
// 	// log.Logger.Info.Println(apiChainNames)
// 	return apiChainNames
// }

// // 체인 proposals 정보 가져오기
// func GetProposalsInfos() []Proposal {
// 	var proposals []Proposal
// 	sql := `SELECT * FROM cosmos.proposals`
// 	DB.Raw(sql).Scan(&proposals)
// 	// log.Logger.Info.Println(proposals)
// 	return proposals
// }

// // wallet 정보 가져오기(밸리데이터 주소가져오는것)
// func GetWalletInfos() []Wallet {
// 	var wallet []Wallet
// 	sql := `SELECT * FROM cosmos.wallets WHERE active = 1 AND type = 'validator' AND owner_name = 'Provalidator'`
// 	DB.Raw(sql).Scan(&wallet)
// 	// log.Logger.Info.Println(wallet)
// 	return wallet
// }

// // good이거나, 아니면 업데이트 된지 1시간이 지난 경우만 가져오기
// func GetApiChainInfos() []ApiChainInfo {
// 	var apiChainInfos []ApiChainInfo
// 	sql := `SELECT chain_name, url, latency, height, update_date
// 	FROM cosmos.api_urls
// 	WHERE result = 'good' OR update_date <= DATE_ADD(NOW(), INTERVAL -1 SECOND)`
// 	DB.Raw(sql).Scan(&apiChainInfos)
// 	// log.Logger.Info.Println(apiChainInfos)
// 	return apiChainInfos
// }

// // 가장 좋은 lcd 서버 url 배열 들고 오기
// func GetApiChainUrl() []ApiChainInfo {
// 	var apiChainInfos []ApiChainInfo
// 	sql := `
// 	SELECT chain_name, url, latency , height , result , update_date
// 	FROM (
// 	  SELECT chain_name, url, latency , height , result , update_date , ROW_NUMBER() OVER (PARTITION BY chain_name ORDER BY height DESC ,latency ASC) AS row_num
// 	  FROM api_urls
// 	) ranked
// 	WHERE row_num = 1
// 	`
// 	DB.Raw(sql).Scan(&apiChainInfos)
// 	// log.Logger.Info.Println(apiChainInfos)
// 	return apiChainInfos
// }

// func GetApiUrls(){
// 	sql := ""
// }

// func WriteLog(sub string, token string, ip string, method string) {
// 	//log.Logger.Info.Println("WriteLog. Token:", token, ", IP:", ip, ", Method:", method)
// 	now := time.Now()
// 	date := now.Format("2006-01-02 15:04:05")
// 	input := Proposal{Sub: sub, Token: token, Ip: ip, Method: method, Date: date, Name: os.Getenv("CHAIN_NAME")}
// 	DB.Create(&input)

// 	// Usage count
// 	ym := now.Format("0601")
// 	uqKey := "tomochain" + ym + "-" + sub
// 	DB.Exec("INSERT INTO nodereal.api_usages_counts(uqkey,chain) values (?,?) ON DUPLICATE KEY UPDATE uqkey = ?, counts = counts+1;", uqKey, "tomochain", uqKey)

// }

// func Count(sub string) int {
// 	apiUsagesCount := Proposal{}
// 	now := time.Now()
// 	// Usage count
// 	ym := now.Format("0601")
// 	uqKey := "tomochain" + ym + "-" + sub
// 	DB.Raw("SELECT counts FROM nodereal.api_usages_counts WHERE uqkey = ?", uqKey).Scan(&apiUsagesCount)
// 	return apiUsagesCount.Counts
// }
