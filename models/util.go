package models

import "google.golang.org/grpc"

var Config config

type config struct {
	Telegram struct {
		BotName     string `yaml:"BOT_NAME"`
		BotToken    string `yaml:"BOT_TOKEN"`
		ChatID      int    `yaml:"CHAT_ID"`
		ChatIDAdmin int    `yaml:"CHAT_ID_ADMIN"`
	} `yaml:"TELEGRAM"`
	Database struct {
		MysqlServerURL    string `yaml:"MYSQL_SERVER_URL"`
		MysqlServerPort   string `yaml:"MYSQL_SERVER_PORT"`
		MysqlUserID       string `yaml:"MYSQL_USER_ID"`
		MysqlUserPW       string `yaml:"MYSQL_USER_PW"`
		MysqlSelectDBName string `yaml:"MYSQL_SELECT_DB_NAME"`
	} `yaml:"DATABASE"`
	ChainInfo []ChainInfo `yaml:"CHAIN_INFO"`
}

type ChainInfo struct {
	ChainName        string  `yaml:"CHAIN_NAME"`
	RPC              string  `yaml:"GRPC"`
	LCD              string  `yaml:"LCD"`
	ValidatorAddress string  `yaml:"VALIDATOR_ADDRESS"`
	PrivKey          string  `yaml:"PRIV_KEY"`
	EX               string  `yaml:"EX"`
	Denom            string  `yaml:"DENOM"`
	LeastAmount      float64 `yaml:"LEAST_AMOUNT"`
	Decimal          float64 `yaml:"DECIMAL"`
	Rate             float64 `yaml:"RATE"`
	Conn             *grpc.ClientConn
}
