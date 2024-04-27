package service

import (
	"errors"
	"fmt"

	"cosmossdk.io/math"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"gorm.io/gorm"

	"github.com/provalidator/stakescan-indexer/model"
	"github.com/provalidator/stakescan-indexer/util"
)

func ChangeBalance(ctx Context, events []abcitypes.Event) error {
	for _, event := range events {
		if event.Type == banktypes.EventTypeCoinSpent || event.Type == banktypes.EventTypeCoinReceived {
			addr := event.Attributes[0].Value
			rawAddr, err := util.DecodeBech32Address(addr)
			if err != nil {
				return fmt.Errorf("decode bech32: %w", err)
			}

			coins, err := sdk.ParseCoinsNormalized(event.Attributes[1].Value)
			if err != nil {
				return fmt.Errorf("parse coins: %w", err)
			}

			for _, coin := range coins {
				balance, err := model.FindBalance(ctx.DB, rawAddr, coin.Denom)
				if err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return fmt.Errorf("find balance: %w", err)
					}
					balance = model.Balance{
						Address: rawAddr,
						Denom:   coin.Denom,
						Amount:  "0",
					}
				}
				beforeAmount, _ := math.NewIntFromString(balance.Amount)
				fmt.Println("beforeAmount", beforeAmount)
				var afterAmount string
				if event.Type == banktypes.EventTypeCoinSpent {
					afterAmount = beforeAmount.Sub(coin.Amount).String()
				} else {
					afterAmount = beforeAmount.Add(coin.Amount).String()
				}
				fmt.Println("afterAmount", afterAmount)
				balance.Amount = afterAmount
				balance.Height = ctx.Height
				//err = model.UpdateBalance(ctx, balance)
				//if err != nil {
				//	return fmt.Errorf("model.UpdateBalance: %w", err)
				//}
			}
		}

		for _, attr := range event.Attributes {
			if attr.Key == "spender" || attr.Key == "receiver" {
				//account := model.Account{Address: attr.Value}
				//err := model.InsertAccounts(ctx.DB, account)
				//if errors.Is(err, fmt.Errorf("0")) {
				//} else if err != nil {
				//	return fmt.Errorf("model.InsertAccounts: %w", err)
				//}
			}
		}
	}
	return nil
}
