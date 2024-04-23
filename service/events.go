package service

import (
	"cosmossdk.io/math"
	"errors"
	"fmt"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/stakescanpoc/models"
	"gorm.io/gorm"
	"strings"
)

func FindBlockEvents(results *coretypes.ResultBlockResults) ([]abcitypes.Event, []abcitypes.Event) {
	var beginEvents []abcitypes.Event
	var endEvents []abcitypes.Event
	for _, BeginBlockEvents := range results.BeginBlockEvents {
		beginEvents = append(beginEvents, BeginBlockEvents)
	}
	for _, EndBlockEvents := range results.EndBlockEvents {
		endEvents = append(endEvents, EndBlockEvents)
	}
	return beginEvents, endEvents
}

func ChangeBalance(DB *gorm.DB, height int64, denom string, events []abcitypes.Event) error {
	for _, event := range events {
		if event.Type == "coin_spent" {
			addr, amount := SetBalance(event, denom)
			err := SubBalance(DB, addr, denom, amount, height)
			if err != nil {
				return fmt.Errorf("AddBalance: %w", err)
			}
		}
		if event.Type == "coin_received" {
			addr, amount := SetBalance(event, denom)
			err := AddBalance(DB, addr, denom, amount, height)
			if err != nil {
				return fmt.Errorf("AddBalance: %w", err)
			}
		}
		for _, attr := range event.Attributes {
			if attr.Key == "spender" || attr.Key == "receiver" {
				account := models.Account{Address: attr.Value}
				err := models.InsertAccounts(DB, account)
				if errors.Is(err, fmt.Errorf("0")) {
				} else if err != nil {
					return fmt.Errorf("models.InsertAccounts: %w", err)
				}
			}
		}
	}
	return nil
}

func SetBalance(events abcitypes.Event, denom string) (string, math.Int) {
	amountStr := strings.Split(events.Attributes[1].Value, denom)[0] // Key == amount
	amount, _ := math.NewIntFromString(amountStr)
	addr := events.Attributes[0].Value // Key == spender
	return addr, amount
}

func SubBalance(DB *gorm.DB, addr, denom string, amount math.Int, height int64) error {
	balance, err := models.FindBalance(DB, addr, denom, height)
	if err != nil {
		return fmt.Errorf("models.FindBalance: %w", err)
	}
	beforeAmount, _ := math.NewIntFromString(balance.Amount)
	newAmount := beforeAmount.Sub(amount).String()
	balance.Amount = newAmount
	err = models.UpdateBalance(DB, balance, height)
	if err != nil {
		return fmt.Errorf("models.UpdateBalance: %w", err)
	}
	return nil
}

func AddBalance(DB *gorm.DB, addr, denom string, amount math.Int, height int64) error {
	balance, err := models.FindBalance(DB, addr, denom, height)
	if err != nil {
		return fmt.Errorf("models.FindBalance: %w", err)
	}
	beforeAmount, _ := math.NewIntFromString(balance.Amount)
	newAmount := beforeAmount.Add(amount).String()
	balance.Amount = newAmount
	err = models.UpdateBalance(DB, balance, height)
	if err != nil {
		return fmt.Errorf("models.UpdateBalance: %w", err)
	}
	return nil
}
