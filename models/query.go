package models

import "time"

type BlockJSON struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockID struct {
			Hash  string `json:"hash"`
			Parts struct {
				Total int    `json:"total"`
				Hash  string `json:"hash"`
			} `json:"parts"`
		} `json:"block_id"`
		Block struct {
			Header struct {
				Version struct {
					Block string `json:"block"`
				} `json:"version"`
				ChainID     string    `json:"chain_id"`
				Height      string    `json:"height"`
				Time        time.Time `json:"time"`
				LastBlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Total int    `json:"total"`
						Hash  string `json:"hash"`
					} `json:"parts"`
				} `json:"last_block_id"`
				LastCommitHash     string `json:"last_commit_hash"`
				DataHash           string `json:"data_hash"`
				ValidatorsHash     string `json:"validators_hash"`
				NextValidatorsHash string `json:"next_validators_hash"`
				ConsensusHash      string `json:"consensus_hash"`
				AppHash            string `json:"app_hash"`
				LastResultsHash    string `json:"last_results_hash"`
				EvidenceHash       string `json:"evidence_hash"`
				ProposerAddress    string `json:"proposer_address"`
			} `json:"header"`
			Data struct {
				Txs []string `json:"txs"`
			} `json:"data"`
			Evidence struct {
				Evidence []interface{} `json:"evidence"`
			} `json:"evidence"`
			LastCommit struct {
				Height  string `json:"height"`
				Round   int    `json:"round"`
				BlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Total int    `json:"total"`
						Hash  string `json:"hash"`
					} `json:"parts"`
				} `json:"block_id"`
				Signatures []struct {
					BlockIDFlag      int       `json:"block_id_flag"`
					ValidatorAddress string    `json:"validator_address"`
					Timestamp        time.Time `json:"timestamp"`
					Signature        string    `json:"signature"`
				} `json:"signatures"`
			} `json:"last_commit"`
		} `json:"block"`
	} `json:"result"`
}

type Parts struct {
	Total int    `json:"total"`
	Hash  string `json:"hash"`
}
type BlockID struct {
	Hash  string `json:"hash"`
	Parts Parts  `json:"parts"`
}
type Version struct {
	Block string `json:"block"`
}
type LastBlockID struct {
	Hash  string `json:"hash"`
	Parts Parts  `json:"parts"`
}
type Header struct {
	Version            Version     `json:"version"`
	ChainID            string      `json:"chain_id"`
	Height             string      `json:"height"`
	Time               time.Time   `json:"time"`
	LastBlockID        LastBlockID `json:"last_block_id"`
	LastCommitHash     string      `json:"last_commit_hash"`
	DataHash           string      `json:"data_hash"`
	ValidatorsHash     string      `json:"validators_hash"`
	NextValidatorsHash string      `json:"next_validators_hash"`
	ConsensusHash      string      `json:"consensus_hash"`
	AppHash            string      `json:"app_hash"`
	LastResultsHash    string      `json:"last_results_hash"`
	EvidenceHash       string      `json:"evidence_hash"`
	ProposerAddress    string      `json:"proposer_address"`
}
type Data struct {
	Txs []string `json:"txs"`
}
type Evidence struct {
	Evidence []interface{} `json:"evidence"`
}
type Signatures struct {
	BlockIDFlag      int       `json:"block_id_flag"`
	ValidatorAddress string    `json:"validator_address"`
	Timestamp        time.Time `json:"timestamp"`
	Signature        string    `json:"signature"`
}
type LastCommit struct {
	Height     string       `json:"height"`
	Round      int          `json:"round"`
	BlockID    BlockID      `json:"block_id"`
	Signatures []Signatures `json:"signatures"`
}
type Block struct {
	Header     Header     `json:"header"`
	Data       Data       `json:"data"`
	Evidence   Evidence   `json:"evidence"`
	LastCommit LastCommit `json:"last_commit"`
}
type Result struct {
	BlockID BlockID `json:"block_id"`
	Block   Block   `json:"block"`
}

type BlockResultsJSON struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Result `json:"result"`
}

type Attributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Index bool   `json:"index"`
}
type Events struct {
	Type       string       `json:"type"`
	Attributes []Attributes `json:"attributes"`
}
type TxsResults struct {
	Code      int      `json:"code"`
	Data      string   `json:"data"`
	Log       string   `json:"log"`
	Info      string   `json:"info"`
	GasWanted string   `json:"gas_wanted"`
	GasUsed   string   `json:"gas_used"`
	Events    []Events `json:"events"`
	Codespace string   `json:"codespace"`
}
type BeginBlockEvents struct {
	Type       string       `json:"type"`
	Attributes []Attributes `json:"attributes"`
}
type EndBlockEvents struct {
	Type       string       `json:"type"`
	Attributes []Attributes `json:"attributes"`
}

// type Block struct {
// 	MaxBytes string `json:"max_bytes"`
// 	MaxGas   string `json:"max_gas"`
// }
// type Evidence struct {
// 	MaxAgeNumBlocks string `json:"max_age_num_blocks"`
// 	MaxAgeDuration  string `json:"max_age_duration"`
// 	MaxBytes        string `json:"max_bytes"`
// }
type Validator struct {
	PubKeyTypes []string `json:"pub_key_types"`
}
type ConsensusParamUpdates struct {
	Block     Block     `json:"block"`
	Evidence  Evidence  `json:"evidence"`
	Validator Validator `json:"validator"`
}

// type Result struct {
// 	Height                string                `json:"height"`
// 	TxsResults            []TxsResults          `json:"txs_results"`
// 	BeginBlockEvents      []BeginBlockEvents    `json:"begin_block_events"`
// 	EndBlockEvents        []EndBlockEvents      `json:"end_block_events"`
// 	ValidatorUpdates      interface{}           `json:"validator_updates"`
// 	ConsensusParamUpdates ConsensusParamUpdates `json:"consensus_param_updates"`
// }

type Event struct {
	Type       string           `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Attributes []EventAttribute `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

type EventAttribute struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Index bool   `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
}

type ExecTxResult struct {
	Code      uint32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data      []byte  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Log       string  `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info      string  `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	GasWanted int64   `protobuf:"varint,5,opt,name=gas_wanted,proto3" json:"gas_wanted,omitempty"`
	GasUsed   int64   `protobuf:"varint,6,opt,name=gas_used,proto3" json:"gas_used,omitempty"`
	Events    []Event `protobuf:"bytes,7,rep,name=events,proto3" json:"events,omitempty"`
	Codespace string  `protobuf:"bytes,8,opt,name=codespace,proto3" json:"codespace,omitempty"`
}
