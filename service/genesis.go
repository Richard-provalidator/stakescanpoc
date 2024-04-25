package service

import (
	"context"
	"encoding/json"
	"fmt"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cometbft/cometbft/types"
	"github.com/stakescanpoc/models"
	"github.com/stakescanpoc/util"
	"gorm.io/gorm"
	"time"
)

type AppState struct {
	Accounts []struct {
		AccountNumber string `json:"account_number"`
		Address       string `json:"address"`
		Coins         []struct {
			Amount string `json:"amount"`
			Denom  string `json:"denom"`
		} `json:"coins"`
		DelegatedFree     []interface{} `json:"delegated_free"`
		DelegatedVesting  []interface{} `json:"delegated_vesting"`
		EndTime           string        `json:"end_time"`
		ModuleName        string        `json:"module_name"`
		ModulePermissions []interface{} `json:"module_permissions"`
		OriginalVesting   []interface{} `json:"original_vesting"`
		SequenceNumber    string        `json:"sequence_number"`
		StartTime         string        `json:"start_time"`
	} `json:"accounts"`
	Auth struct {
		Params struct {
			MaxMemoCharacters      string `json:"max_memo_characters"`
			SigVerifyCostEd25519   string `json:"sig_verify_cost_ed25519"`
			SigVerifyCostSecp256K1 string `json:"sig_verify_cost_secp256k1"`
			TxSigLimit             string `json:"tx_sig_limit"`
			TxSizeCostPerByte      string `json:"tx_size_cost_per_byte"`
		} `json:"params"`
	} `json:"auth"`
	Bank struct {
		SendEnabled bool `json:"send_enabled"`
	} `json:"bank"`
	Crisis struct {
		ConstantFee struct {
			Amount string `json:"amount"`
			Denom  string `json:"denom"`
		} `json:"constant_fee"`
	} `json:"crisis"`
	Distribution struct {
		BaseProposerReward     string `json:"base_proposer_reward"`
		BonusProposerReward    string `json:"bonus_proposer_reward"`
		CommunityTax           string `json:"community_tax"`
		DelegatorStartingInfos []struct {
			DelegatorAddress string `json:"delegator_address"`
			StartingInfo     struct {
				Height         string `json:"height"`
				PreviousPeriod string `json:"previous_period"`
				Stake          string `json:"stake"`
			} `json:"starting_info"`
			ValidatorAddress string `json:"validator_address"`
		} `json:"delegator_starting_infos"`
		DelegatorWithdrawInfos []struct {
			DelegatorAddress string `json:"delegator_address"`
			WithdrawAddress  string `json:"withdraw_address"`
		} `json:"delegator_withdraw_infos"`
		FeePool struct {
			CommunityPool []struct {
				Amount string `json:"amount"`
				Denom  string `json:"denom"`
			} `json:"community_pool"`
		} `json:"fee_pool"`
		OutstandingRewards []struct {
			OutstandingRewards interface{} `json:"outstanding_rewards"`
			ValidatorAddress   string      `json:"validator_address"`
		} `json:"outstanding_rewards"`
		PreviousProposer                string `json:"previous_proposer"`
		ValidatorAccumulatedCommissions []struct {
			Accumulated      interface{} `json:"accumulated"`
			ValidatorAddress string      `json:"validator_address"`
		} `json:"validator_accumulated_commissions"`
		ValidatorCurrentRewards []struct {
			Rewards struct {
				Period  string      `json:"period"`
				Rewards interface{} `json:"rewards"`
			} `json:"rewards"`
			ValidatorAddress string `json:"validator_address"`
		} `json:"validator_current_rewards"`
		ValidatorHistoricalRewards []struct {
			Period  string `json:"period"`
			Rewards struct {
				CumulativeRewardRatio interface{} `json:"cumulative_reward_ratio"`
				ReferenceCount        int         `json:"reference_count"`
			} `json:"rewards"`
			ValidatorAddress string `json:"validator_address"`
		} `json:"validator_historical_rewards"`
		ValidatorSlashEvents []interface{} `json:"validator_slash_events"`
		WithdrawAddrEnabled  bool          `json:"withdraw_addr_enabled"`
	} `json:"distribution"`
	Gentxs interface{} `json:"gentxs"`
	Gov    struct {
		DepositParams struct {
			MaxDepositPeriod string `json:"max_deposit_period"`
			MinDeposit       []struct {
				Amount string `json:"amount"`
				Denom  string `json:"denom"`
			} `json:"min_deposit"`
		} `json:"deposit_params"`
		Deposits []struct {
			Amount []struct {
				Amount string `json:"amount"`
				Denom  string `json:"denom"`
			} `json:"amount"`
			Depositor  string `json:"depositor"`
			ProposalID string `json:"proposal_id"`
		} `json:"deposits"`
		Proposals []struct {
			Content struct {
				Type  string `json:"type"`
				Value struct {
					Description string `json:"description"`
					Title       string `json:"title"`
				} `json:"value"`
			} `json:"content"`
			DepositEndTime   time.Time `json:"deposit_end_time"`
			FinalTallyResult struct {
				Abstain    string `json:"abstain"`
				No         string `json:"no"`
				NoWithVeto string `json:"no_with_veto"`
				Yes        string `json:"yes"`
			} `json:"final_tally_result"`
			ID             string    `json:"id"`
			ProposalStatus string    `json:"proposal_status"`
			SubmitTime     time.Time `json:"submit_time"`
			TotalDeposit   []struct {
				Amount string `json:"amount"`
				Denom  string `json:"denom"`
			} `json:"total_deposit"`
			VotingEndTime   time.Time `json:"voting_end_time"`
			VotingStartTime time.Time `json:"voting_start_time"`
		} `json:"proposals"`
		StartingProposalID string `json:"starting_proposal_id"`
		TallyParams        struct {
			Quorum    string `json:"quorum"`
			Threshold string `json:"threshold"`
			Veto      string `json:"veto"`
		} `json:"tally_params"`
		Votes        []interface{} `json:"votes"`
		VotingParams struct {
			VotingPeriod string `json:"voting_period"`
		} `json:"voting_params"`
	} `json:"gov"`
	Mint struct {
		Minter struct {
			AnnualProvisions string `json:"annual_provisions"`
			Inflation        string `json:"inflation"`
		} `json:"minter"`
		Params struct {
			BlocksPerYear       string `json:"blocks_per_year"`
			GoalBonded          string `json:"goal_bonded"`
			InflationMax        string `json:"inflation_max"`
			InflationMin        string `json:"inflation_min"`
			InflationRateChange string `json:"inflation_rate_change"`
			MintDenom           string `json:"mint_denom"`
		} `json:"params"`
	} `json:"mint"`
	Slashing struct {
		MissedBlocks struct {
			Cosmosvalcons10Fqk678Xexhu7Y2Ef89Qjtnphnvahzsc30Vjap []struct {
				Index  string `json:"index"`
				Missed bool   `json:"missed"`
			} `json:"cosmosvalcons10fqk678xexhu7y2ef89qjtnphnvahzsc30vjap"`
		} `json:"missed_blocks"`
		Params struct {
			DowntimeJailDuration    string `json:"downtime_jail_duration"`
			MaxEvidenceAge          string `json:"max_evidence_age"`
			MinSignedPerWindow      string `json:"min_signed_per_window"`
			SignedBlocksWindow      string `json:"signed_blocks_window"`
			SlashFractionDoubleSign string `json:"slash_fraction_double_sign"`
			SlashFractionDowntime   string `json:"slash_fraction_downtime"`
		} `json:"params"`
		SigningInfos struct {
			Cosmosvalcons10Fqk678Xexhu7Y2Ef89Qjtnphnvahzsc30Vjap struct {
				IndexOffset         string    `json:"index_offset"`
				JailedUntil         time.Time `json:"jailed_until"`
				MissedBlocksCounter string    `json:"missed_blocks_counter"`
				StartHeight         string    `json:"start_height"`
				Tombstoned          bool      `json:"tombstoned"`
			} `json:"cosmosvalcons10fqk678xexhu7y2ef89qjtnphnvahzsc30vjap"`
		} `json:"signing_infos"`
	} `json:"slashing"`
	Staking struct {
		Delegations []struct {
			DelegatorAddress string `json:"delegator_address"`
			Shares           string `json:"shares"`
			ValidatorAddress string `json:"validator_address"`
		} `json:"delegations"`
		Exported            bool   `json:"exported"`
		LastTotalPower      string `json:"last_total_power"`
		LastValidatorPowers []struct {
			Address string `json:"Address"`
			Power   string `json:"Power"`
		} `json:"last_validator_powers"`
		Params struct {
			BondDenom     string `json:"bond_denom"`
			MaxEntries    int    `json:"max_entries"`
			MaxValidators int    `json:"max_validators"`
			UnbondingTime string `json:"unbonding_time"`
		} `json:"params"`
		Redelegations []struct {
			DelegatorAddress string `json:"delegator_address"`
			Entries          []struct {
				CompletionTime time.Time `json:"completion_time"`
				CreationHeight string    `json:"creation_height"`
				InitialBalance string    `json:"initial_balance"`
				SharesDst      string    `json:"shares_dst"`
			} `json:"entries"`
			ValidatorDstAddress string `json:"validator_dst_address"`
			ValidatorSrcAddress string `json:"validator_src_address"`
		} `json:"redelegations"`
		UnbondingDelegations []struct {
			DelegatorAddress string `json:"delegator_address"`
			Entries          []struct {
				Balance        string    `json:"balance"`
				CompletionTime time.Time `json:"completion_time"`
				CreationHeight string    `json:"creation_height"`
				InitialBalance string    `json:"initial_balance"`
			} `json:"entries"`
			ValidatorAddress string `json:"validator_address"`
		} `json:"unbonding_delegations"`
		Validators []struct {
			Commission struct {
				CommissionRates struct {
					MaxChangeRate string `json:"max_change_rate"`
					MaxRate       string `json:"max_rate"`
					Rate          string `json:"rate"`
				} `json:"commission_rates"`
				UpdateTime time.Time `json:"update_time"`
			} `json:"commission"`
			ConsensusPubkey string `json:"consensus_pubkey"`
			DelegatorShares string `json:"delegator_shares"`
			Description     struct {
				Details  string `json:"details"`
				Identity string `json:"identity"`
				Moniker  string `json:"moniker"`
				Website  string `json:"website"`
			} `json:"description"`
			Jailed            bool      `json:"jailed"`
			MinSelfDelegation string    `json:"min_self_delegation"`
			OperatorAddress   string    `json:"operator_address"`
			Status            int       `json:"status"`
			Tokens            string    `json:"tokens"`
			UnbondingHeight   string    `json:"unbonding_height"`
			UnbondingTime     time.Time `json:"unbonding_time"`
		} `json:"validators"`
	} `json:"staking"`
	Supply struct {
		Supply []interface{} `json:"supply"`
	} `json:"supply"`
}

func GetGenesis(RPC string, height int64) (*coretypes.ResultBlock, error) {
	rpcClient, err := rpchttp.New(RPC, "/websocket")
	if err != nil {
		// log.Logger.Error.Println("connectRPC Failed : ", err)
		fmt.Println("connectRPC Failed : ", err)
		return nil, err
	}
	results, err := rpcClient.Block(context.Background(), &height)
	if err != nil {
		// log.Logger.Error.Println("GetBlockByHeight Failed : ", err)
		fmt.Println("GetBlockByHeight Failed : ", err)
		return nil, err
	}

	return results, nil
}

func readGenesis(rootPath string) (*types.GenesisDoc, error) {
	filePath := rootPath + "/config/genesis.json"
	genesisDoc, err := types.GenesisDocFromFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("types.GenesisDocFromFile : %w", err)
	}
	return genesisDoc, nil
}

func InsertGenesisBlock(DB *gorm.DB, rootPath string) error {
	genesis, err := readGenesis(rootPath)
	if err != nil {
		return fmt.Errorf("readGenesis : %w", err)
	}
	dbblock := models.Block{
		Height:                   genesis.InitialHeight,
		ProposerConsensusAddress: "",
		Block:                    nil,
		Timestamp:                genesis.GenesisTime,
	}
	err = InsertBlock(DB, dbblock)
	if err != nil {
		return fmt.Errorf("InsertBlock : %w", err)
	}

	var validators []models.Validator
	for _, genval := range genesis.Validators {
		bech32ConAddr, err := util.MakeBech32Address(genval.PubKey.Address().String())
		if err != nil {
			return fmt.Errorf("util.MakeBech32Address(genval.PubKey.Address().String()): %w", err)
		}
		bech32Addr, err := util.MakeBech32Address(genval.Address.String())
		if err != nil {
			return fmt.Errorf("util.MakeBech32Address(genval.Address.String()): %w", err)
		}
		val := models.Validator{
			ConsensusAddress: bech32ConAddr,
			Address:          bech32Addr,
			Moniker:          genval.Name,
			Jailed:           false,
			Status:           "",
			VotingPower:      fmt.Sprint(genval.Power),
			Tokens:           fmt.Sprint(genval.Power),
			CommissionRate:   "",
			Info:             nil,
		}
		validators = append(validators, val)
	}
	err = InsertValidators(DB, validators)
	if err != nil {
		return fmt.Errorf("InsertValidators : %w", err)
	}
	appState := make(map[string]interface{})
	err = json.Unmarshal(genesis.AppState, &appState)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}
	// 파싱된 데이터를 사용하여 작업
	moniker := appState["description"].(map[string]interface{})["moniker"].(string)
	fmt.Println("Moniker:", moniker)

	operatorAddress := appState["operator_address"].(string)
	fmt.Println("Operator Address:", operatorAddress)

	tokens := appState["tokens"].(string)
	fmt.Println("Tokens:", tokens)
	return nil
}
