package config

import (
	"fmt"
	"os"
	"runtime"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Telegram Telegram `yaml:"TELEGRAM"`
	Chains   []Chain  `yaml:"CHAINS"`
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

type Chain struct {
	ChainName        string   `yaml:"CHAIN_NAME"`
	RPC              string   `yaml:"RPC"`
	ValidatorAddress string   `yaml:"VALIDATOR_ADDRESS"`
	EX               string   `yaml:"EX"`
	Denom            string   `yaml:"DENOM"`
	Decimal          float64  `yaml:"DECIMAL"`
	Database         Database `yaml:"DATABASE"`
}

func (d Database) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		d.MysqlUserID, d.MysqlUserPW, d.MysqlServerURL, d.MysqlServerPort, d.MysqlSelectDBName)
}

func LoadYaml(rootPath string) (Config, error) {
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
