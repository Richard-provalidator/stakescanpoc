package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Chain    Chain    `yaml:"chain"`
	DB       DB       `yaml:"db"`
	Telegram Telegram `yaml:"telegram"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (d DB) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		d.User, d.Password, d.Host, d.Port, d.Database)
}

type Chain struct {
	Name string `yaml:"name"`
	RPC  string `yaml:"rpc"`
}

type Telegram struct {
	BotName     string `yaml:"bot_name"`
	BotToken    string `yaml:"bot_token"`
	ChatID      int64  `yaml:"chat_id"`
	ChatIDAdmin int    `yaml:"chat_id_admin"`
}

func Load(filename string) (Config, error) {
	var cfg Config
	f, err := os.Open(filename)
	if err != nil {
		return cfg, err
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("unmarshal yaml: %w", err)
	}
	return cfg, nil
}
