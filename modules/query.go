package modules

import (
	"github.com/go-resty/resty/v2"
	"github.com/stakescanpoc/log"
)

func connectRPC(RPC, connectURL string) (*resty.Response, error) {
	client := resty.New()
	res, err := client.R().Get(RPC + connectURL)

	if err != nil {
		log.Logger.Error.Println("connectRPC Failed : ", err)
		return nil, err
	}
	return res, nil
}
