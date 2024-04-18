package config

import (
	"fmt"
	"os"
	"runtime"

	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Telegram   Telegram    `yaml:"TELEGRAM"`
	Database   Database    `yaml:"DATABASE"`
	ChainInfos []ChainInfo `yaml:"CHAIN_INFOS"`
}

type Telegram struct {
	BotName     string `yaml:"BOT_NAME"`
	BotToken    string `yaml:"BOT_TOKEN"`
	ChatID      int64  `yaml:"CHAT_ID"`
	ChatIDAdmin int    `yaml:"CHAT_ID_ADMIN"`
}

type Database struct {
	MysqlServerURL    string `yaml:"MYSQL_SERVER_URL"`
	MysqlServerPort   string `yaml:"MYSQL_SERVER_PORT"`
	MysqlUserID       string `yaml:"MYSQL_USER_ID"`
	MysqlUserPW       string `yaml:"MYSQL_USER_PW"`
	MysqlSelectDBName string `yaml:"MYSQL_SELECT_DB_NAME"`
}

type ChainInfo struct {
	ChainName        string  `yaml:"CHAIN_NAME"`
	RPC              string  `yaml:"RPC"`
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

func (d Database) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		d.MysqlUserID, d.MysqlUserPW, d.MysqlServerURL, d.MysqlServerPort, d.MysqlSelectDBName)
}

func loadYaml(rootPath string) (Config, error) {
	var cfg Config
	var fileName string
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		fileName = rootPath + "/config/local.yaml"
	} else {
		fileName = rootPath + "/config/prod.yaml"
	}
	f, err := os.ReadFile(fileName)
	if err != nil {
		return cfg, fmt.Errorf("os.ReadFile :%w", err)
	}
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return cfg, fmt.Errorf("yaml.Unmarshal: %w", err)
	}
	return cfg, nil
}
