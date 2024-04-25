package config

import (
	"flag"
	"fmt"
	"github.com/cosmos/gaia/v15/app/params"
	"os"
	"runtime"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Context struct {
	Chain  Chain
	DB     *gorm.DB
	Height int64
	EncCfg EncCfg
}

type EncCfg struct {
	Cosmos params.EncodingConfig
}

func GetRootPath() string {
	// home 옵션을 위한 환경 변수 값 가져오기
	home := os.Getenv("HOME")
	if home == "" {
		if runtime.GOOS == "windows" {
			home = "C:\\Users\\user\\GolandProjects\\stakescanpoc"
		} else {
			home = "%HOME/go/src/github.com/stakescanpoc"
		}
	}

	// home 옵션을 위한 플래그 정의
	homeFlag := flag.String("home", home, "path to home directory")

	// 플래그 파싱
	flag.Parse()

	// 사용자가 입력한 옵션 출력
	fmt.Println("Home directory:", *homeFlag)
	return home
}

func GetDirs(Chains []Chain, rootDir string) map[string]string {
	DirsMap := make(map[string]string)
	for _, chain := range Chains {
		lowerChainName := strings.ToLower(chain.ChainName)
		DirsMap[chain.ChainName+"txsDir"] = rootDir + "/txs/" + lowerChainName + "/"
		DirsMap["csvDir"] = rootDir + "/csv/"
		DirsMap["modulesDir"] = rootDir + "/modules/"
	}
	return DirsMap
}

func ConnectDatabase(d Database) (*gorm.DB, error) {
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

func (t Telegram) SendTelegramMsg(msgStr string) error {
	bot, err := tgbotapi.NewBotAPI(t.BotToken)
	if err != nil {
		return fmt.Errorf("tgbotapi.NewBotAPI: %w", err)
	}

	// Create a message config
	msg := tgbotapi.NewMessage(t.ChatID, msgStr)

	// Send the message
	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("bot.Send: %w", err)
	}
	return nil
}

// // 예전 GetRootPath(){
// // Get the absolute path of the current working directory.
// dir, err := os.Getwd()
// if err != nil {
// 	panic(err)
// }
// for {
// 	// Check if a go.mod file exists in the current directory.
// 	if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
// 		return dir
// 	}

// 	// Move up one directory.
// 	parentDir := filepath.Dir(dir)

// 	// If we've reached the root directory ("/"), exit the loop.
// 	if parentDir == dir {
// 		break
// 	}

// 	// Continue searching in the parent directory.
// 	dir = parentDir
// }

// return ""
// }
