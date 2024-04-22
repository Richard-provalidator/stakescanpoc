package models

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"time"
)

type Blocks struct {
	Height                   int64
	ProposerConsensusAddress string
	Block                    types.Block
	Timestamp                time.Time
}

type Transactions struct {
	TxID        int64
	TxIDX       int
	Code        uint32
	TxHash      string
	Height      int64
	Messages    tx.MsgResponse
	MessageType string
	Events      []abcitypes.Event
	GasWanted   int64
	GasUsed     int64
}
type Messages struct {
	Type string
	Msg  interface{}
}
type MapTransactionAddress struct {
	TxID  int64
	AccID int64
}

type Accounts struct {
	AccID   int64
	Address string
}

type Balances struct {
	Address string
	Denom   string
	Amount  string
	Height  int64
}

type Assets struct {
	Denom     string
	Amount    string
	ChannelID string
}

type Validators struct {
	ConsensusAddress string
	Address          string
	Moniker          string
	Jailed           bool
	Status           string
	VotingPower      string
	Tokens           string
	CommissionRate   string
	Info             stakingtypes.Validator
}

type ValidatorsSign struct {
	ConsensusAddress string
	Height           int64
}

type Proposals struct {
	ProposalID int64
	Proposal   v1beta1.Proposal
}

type ValidatorsProposalVote struct {
	ConsensusAddress string
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
