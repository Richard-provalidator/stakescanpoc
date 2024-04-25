package models

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"gorm.io/datatypes"
	"time"
)

type Block struct {
	Height                   int64
	ProposerConsensusAddress string
	Block                    datatypes.JSON
	Timestamp                time.Time
}

type Transaction struct {
	TxID        int64
	TxIDX       int
	Code        uint32
	TxHash      string
	Height      int64
	Messages    datatypes.JSON
	MessageType string
	Events      []abcitypes.Event
	GasWanted   int64
	GasUsed     int64
}

type Event struct {
	Type       string
	Attributes []EventAttribute
}
type EventAttribute struct {
	Key   string
	Value string
	Index bool
}

type MapTransactionAddress struct {
	TxID  int64
	AccID int64
}

type Account struct {
	AccID   int64
	Address []byte
}

type Balance struct {
	Address []byte
	Denom   string
	Amount  string
	Height  int64
}

type Asset struct {
	Denom     string
	Amount    string
	ChannelID string
}

type Validator struct {
	ConsensusAddress []byte
	Address          []byte
	Moniker          string
	Jailed           bool
	Status           string
	VotingPower      string
	Tokens           string
	CommissionRate   string
	Info             datatypes.JSON
}

type ValidatorsSign struct {
	ConsensusAddress []byte
	Height           int64
}

type Proposal struct {
	ProposalID uint64
	Proposal   datatypes.JSON
}

type ValidatorsProposalVote struct {
	ConsensusAddress []byte
	VoteOption       string
	TxID             int64
}

type Ibc struct {
	TxID   int64
	Status string
	Denom  string
	Amount string
}

type IbcChannel struct {
	ChannelID     string
	State         bool
	EscrowAddress string
	ChainID       string
}

type CoinGecko struct {
	Denom string
	Price int64
}
