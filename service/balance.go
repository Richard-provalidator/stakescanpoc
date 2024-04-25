package service

import (
	"cosmossdk.io/math"
	"fmt"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/stakescanpoc/config"
	"github.com/stakescanpoc/models"
	"github.com/stakescanpoc/util"
	"strings"
)

func ChangeBalance(ctx config.Context, events []abcitypes.Event) error {
	for _, event := range events {
		if event.Type == "coin_spent" || event.Type == "coin_received" {
			splitValue := strings.Split(event.Attributes[1].Value, ",") // Key == amount
			addr := event.Attributes[0].Value                           // Key == spender
			bech32Addr, err := util.MakeBech32Address(addr)
			if err != nil {
				return fmt.Errorf("util.MaMakeBech32Address(addr): %w", err)
			}
			for _, value := range splitValue {
				ibcValues := strings.Split(value, "ibc/")
				denomValues := strings.Split(value, ctx.Chain.Denom)

				var amount math.Int
				var denom string
				if len(ibcValues) == 2 {
					amount, _ = math.NewIntFromString(ibcValues[0])
					denom = "ibc/" + ibcValues[1]
				} else if len(denomValues) == 2 {
					amount, _ = math.NewIntFromString(denomValues[0])
					denom = ctx.Chain.Denom
				} else {
					continue // ibc/ 또는 uatom으로 분할된 두 부분으로 나누어지지 않는 경우 건너뜁니다.
				}
				balance, err := models.FindBalance(ctx, bech32Addr, denom)
				if err != nil {
					return fmt.Errorf("models.FindBalance: %w", err)
				}
				beforeAmount, _ := math.NewIntFromString(balance.Amount)
				fmt.Println("beforeAmount", beforeAmount)
				var newAmount string
				if event.Type == "coin_spent" {
					newAmount = beforeAmount.Sub(amount).String()
				} else if event.Type == "coin_received" {
					newAmount = beforeAmount.Add(amount).String()
				}
				fmt.Println("sub newAmount", newAmount)
				balance.Amount = newAmount
				//err = models.UpdateBalance(ctx, balance)
				//if err != nil {
				//	return fmt.Errorf("models.UpdateBalance: %w", err)
				//}
			}

		}
		for _, attr := range event.Attributes {
			if attr.Key == "spender" || attr.Key == "receiver" {
				//account := models.Account{Address: attr.Value}
				//err := models.InsertAccounts(ctx.DB, account)
				//if errors.Is(err, fmt.Errorf("0")) {
				//} else if err != nil {
				//	return fmt.Errorf("models.InsertAccounts: %w", err)
				//}
			}
		}
	}
	return nil
}
