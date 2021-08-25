package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/axelarnetwork/axelar-core/x/tss/exported"
	voting "github.com/axelarnetwork/axelar-core/x/vote/exported"
	tssTypes "github.com/axelarnetwork/axelar-core/x/tss/types"

	"github.com/axelarnetwork/axelar-core/x/bitcoin/types"
)

// Query paths
const (
	QuerySigStatus            = "sig-status"
	QueryKeyStatus            = "key-status"
	QueryRecovery             = "recovery"
	QueryKeyID                = "key-id"
	QueryKeySharesByKeyID     = "key-share-id"
	QueryKeySharesByValidator = "key-share-validator"
	QueryDeactivated          = "deactivated"
)

// NewQuerier returns a new querier for the TSS module
func NewQuerier(k tssTypes.TSSKeeper, v tssTypes.Voter, s tssTypes.Snapshotter, staking tssTypes.StakingKeeper, n tssTypes.Nexus) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var res []byte
		var err error
		switch path[0] {
		case QuerySigStatus:
			res, err = querySigStatus(ctx, k, v, path[1])
		case QueryKeyStatus:
			res, err = queryKeygenStatus(ctx, k, v, path[1])
		case QueryRecovery:
			res, err = queryRecovery(ctx, k, s, path[1])
		case QueryKeyID:
			res, err = queryKeyID(ctx, k, n, path[1], path[2])
		case QueryKeySharesByKeyID:
			res, err = queryKeySharesByKeyID(ctx, k, s, path[1])
		case QueryKeySharesByValidator:
			res, err = queryKeySharesByValidator(ctx, k, n, s, path[1])
		case QueryDeactivated:
			res, err = queryDeactivatedOperator(ctx, k, s, staking)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("unknown tss query endpoint: %s", path[0]))
		}

		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrBitcoin, err.Error())
		}
		return res, nil
	}
}

func queryRecovery(ctx sdk.Context, k tssTypes.TSSKeeper, s tssTypes.Snapshotter, keyID string) ([]byte, error) {
	counter, ok := k.GetSnapshotCounterForKeyID(ctx, keyID)
	if !ok {
		return nil, fmt.Errorf("could not obtain snapshot counter for key ID %s", keyID)
	}

	snapshot, ok := s.GetSnapshot(ctx, counter)
	if !ok {
		return nil, fmt.Errorf("could not obtain snapshot for counter %d", counter)
	}

	participants := make([]string, 0, len(snapshot.Validators))
	participantShareCounts := make([]uint32, 0, len(snapshot.Validators))
	for _, validator := range snapshot.Validators {
		participants = append(participants, validator.GetSDKValidator().GetOperator().String())
		participantShareCounts = append(participantShareCounts, uint32(validator.ShareCount))
	}

	infos := k.GetAllRecoveryInfos(ctx, keyID)

	resp := tssTypes.QueryRecoveryResponse{
		Threshold:          int32(snapshot.CorruptionThreshold),
		PartyUids:          participants,
		PartyShareCounts:   participantShareCounts,
		ShareRecoveryInfos: infos,
	}

	return resp.Marshal()
}

func querySigStatus(ctx sdk.Context, k tssTypes.TSSKeeper, v tssTypes.Voter, sigID string) ([]byte, error) {
	var resp tssTypes.QuerySigResponse
	if sig, status := k.GetSig(ctx, sigID); status == exported.SigStatus_Signed {
		// poll was successful
		resp := tssTypes.QuerySigResponse{
			VoteStatus: tssTypes.VoteStatus_Decided,
			Signature: &tssTypes.Signature{
				R: sig.R.Bytes(),
				S: sig.S.Bytes(),
			},
		}
		return resp.Marshal()
	}

	pollMeta := voting.NewPollKey(tssTypes.ModuleName, sigID)
	poll := v.GetPoll(ctx, pollMeta)

	if poll == nil {
		// poll either never existed or has been closed
		resp.VoteStatus = tssTypes.VoteStatus_Unspecified
	} else {
		// poll still open, pending a decision
		resp.VoteStatus = tssTypes.VoteStatus_Pending
	}

	return resp.Marshal()
}

func queryKeygenStatus(ctx sdk.Context, k tssTypes.TSSKeeper, v tssTypes.Voter, keyID string) ([]byte, error) {
	var resp tssTypes.QueryKeyResponse

	if key, ok := k.GetKey(ctx, keyID); ok {
		// poll was successful
		resp = tssTypes.QueryKeyResponse{
			VoteStatus: tssTypes.VoteStatus_Decided,
			Role:       key.Role,
		}

		return resp.Marshal()
	}

	pollMeta := voting.NewPollKey(tssTypes.ModuleName, keyID)
	poll := v.GetPoll(ctx, pollMeta)
	if poll == nil {
		// poll either never existed or has been closed
		resp.VoteStatus = tssTypes.VoteStatus_Unspecified
	} else {
		// poll still open, pending a decision
		resp.VoteStatus = tssTypes.VoteStatus_Pending
	}

	return resp.Marshal()
}

// queryKeyID returns the keyID of the most recent key for a provided keyChain and keyRole
func queryKeyID(ctx sdk.Context, k tssTypes.TSSKeeper, n tssTypes.Nexus, keyChainStr string, keyRoleStr string) ([]byte, error) {
	keyChain, ok := n.GetChain(ctx, keyChainStr)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", keyChainStr)
	}

	keyRole, err := exported.KeyRoleFromSimpleStr(keyRoleStr)
	if err != nil {
		return nil, err
	}

	if keyRole == exported.ExternalKey {
		return nil, fmt.Errorf("use the chain specific query for %s to get external keyIDs", keyChainStr)
	}

	keyID, found := k.GetCurrentKeyID(ctx, keyChain, keyRole)
	if !found {
		return nil, fmt.Errorf("no key from chain %s role %s exists", keyChainStr, keyRoleStr)
	}

	return []byte(keyID), nil
}

func queryKeySharesByKeyID(ctx sdk.Context, k tssTypes.TSSKeeper, s tssTypes.Snapshotter, keyID string) ([]byte, error) {

	counter, ok := k.GetSnapshotCounterForKeyID(ctx, keyID)
	if !ok {
		return nil, fmt.Errorf("invalid keyID %s", keyID)
	}

	snapshot, ok := s.GetSnapshot(ctx, counter)
	if !ok {
		return nil, fmt.Errorf("no snapshot found for counter number %d", counter)
	}

	var allShareInfos []tssTypes.QueryKeyShareResponse_ShareInfo
	for _, validator := range snapshot.Validators {

		thisShareInfo := tssTypes.QueryKeyShareResponse_ShareInfo{
			KeyID:               keyID,
			SnapshotBlockNumber: snapshot.Height,
			ValidatorAddress:    validator.GetSDKValidator().GetOperator().String(),
			NumValidatorShares:  validator.ShareCount,
			NumTotalShares:      snapshot.TotalShareCount.Int64(),
		}

		allShareInfos = append(allShareInfos, thisShareInfo)
	}

	keyShareInfos := tssTypes.QueryKeyShareResponse{
		ShareInfos: allShareInfos,
	}

	return keyShareInfos.Marshal()
}

func queryKeySharesByValidator(ctx sdk.Context, k tssTypes.TSSKeeper, n tssTypes.Nexus, s tssTypes.Snapshotter, targetValidatorAddr string) ([]byte, error) {

	var allShareInfos []tssTypes.QueryKeyShareResponse_ShareInfo

	for _, chain := range n.GetChains(ctx) {
		for _, keyRole := range exported.GetKeyRoles() {

			keyID, found := k.GetCurrentKeyID(ctx, chain, keyRole)

			if !found {
				continue
			}

			counter, ok := k.GetSnapshotCounterForKeyID(ctx, keyID)
			if !ok {
				return nil, fmt.Errorf("could not get snapshot counter from keyID %s", keyID)
			}

			snapshot, ok := s.GetSnapshot(ctx, counter)
			if !ok {
				return nil, fmt.Errorf("no snapshot found for counter number %d", counter)
			}

			for _, validator := range snapshot.Validators {

				validatorAddr := validator.GetSDKValidator().GetOperator().String()
				if validatorAddr == targetValidatorAddr {

					thisShareInfo := tssTypes.QueryKeyShareResponse_ShareInfo{
						KeyID:               keyID,
						KeyChain:            chain.Name,
						KeyRole:             keyRole.String(),
						SnapshotBlockNumber: snapshot.Height,
						ValidatorAddress:    validator.GetSDKValidator().GetOperator().String(),
						NumValidatorShares:  validator.ShareCount,
						NumTotalShares:      snapshot.TotalShareCount.Int64(),
					}
					allShareInfos = append(allShareInfos, thisShareInfo)
					break
				}
			}
		}
	}

	keyShareInfos := tssTypes.QueryKeyShareResponse{
		ShareInfos: allShareInfos,
	}

	return keyShareInfos.Marshal()
}

func queryDeactivatedOperator(ctx sdk.Context, k tssTypes.TSSKeeper, s tssTypes.Snapshotter, staking tssTypes.StakingKeeper) ([]byte, error) {

	var deactivatedValidators []string
	validatorIter := func(_ int64, validator stakingtypes.ValidatorI) (stop bool) {

		// this explicit type cast is necessary, because we need to call UnpackInterfaces() on the validator
		// and it is not exposed in the ValidatorI interface
		v, ok := validator.(stakingtypes.Validator)
		if !ok {
			k.Logger(ctx).Error(fmt.Sprintf("unexpected validator type: expected %T, got %T", stakingtypes.Validator{}, validator))
			return false
		}

		_, active := s.GetProxy(ctx, v.GetOperator())
		if !active {
			deactivatedValidators = append(deactivatedValidators, v.GetOperator().String())
		}

		return false
	}
	// IterateBondedValidatorsByPower(https://github.com/cosmos/cosmos-sdk/blob/7fc7b3f6ff82eb5ede52881778114f6b38bd7dfa/x/staking/keeper/alias_functions.go#L33) iterates validators by power in descending order
	staking.IterateBondedValidatorsByPower(ctx, validatorIter)

	resp := tssTypes.QueryDeactivatedOperatorsResponse{
		OperatorAddresses: deactivatedValidators,
	}

	return types.ModuleCdc.MarshalBinaryLengthPrefixed(&resp)
}
