// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/axelarnetwork/axelar-core/x/vote/exported"
	"github.com/axelarnetwork/axelar-core/x/vote/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	"sync"
)

// Ensure, that StoreMock does implement types.Store.
// If this is not the case, regenerate this file with moq.
var _ types.Store = &StoreMock{}

// StoreMock is a mock implementation of types.Store.
//
// 	func TestSomethingThatUsesStore(t *testing.T) {
//
// 		// make and configure a mocked types.Store
// 		mockedStore := &StoreMock{
// 			DeletePollFunc: func()  {
// 				panic("mock out the DeletePoll method")
// 			},
// 			GetPollFunc: func(key exported.PollKey) exported.Poll {
// 				panic("mock out the GetPoll method")
// 			},
// 			GetShareCountFunc: func(voter github_com_cosmos_cosmos_sdk_types.ValAddress) (int64, bool) {
// 				panic("mock out the GetShareCount method")
// 			},
// 			GetTotalShareCountFunc: func() github_com_cosmos_cosmos_sdk_types.Int {
// 				panic("mock out the GetTotalShareCount method")
// 			},
// 			GetTotalVoterCountFunc: func() int64 {
// 				panic("mock out the GetTotalVoterCount method")
// 			},
// 			GetVoteFunc: func(hash string) (types.TalliedVote, bool) {
// 				panic("mock out the GetVote method")
// 			},
// 			GetVotesFunc: func() []types.TalliedVote {
// 				panic("mock out the GetVotes method")
// 			},
// 			HasVotedFunc: func(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool {
// 				panic("mock out the HasVoted method")
// 			},
// 			SetMetadataFunc: func(metadata exported.PollMetadata)  {
// 				panic("mock out the SetMetadata method")
// 			},
// 			SetVoteFunc: func(voter github_com_cosmos_cosmos_sdk_types.ValAddress, vote types.TalliedVote)  {
// 				panic("mock out the SetVote method")
// 			},
// 		}
//
// 		// use mockedStore in code that requires types.Store
// 		// and then make assertions.
//
// 	}
type StoreMock struct {
	// DeletePollFunc mocks the DeletePoll method.
	DeletePollFunc func()

	// GetPollFunc mocks the GetPoll method.
	GetPollFunc func(key exported.PollKey) exported.Poll

	// GetShareCountFunc mocks the GetShareCount method.
	GetShareCountFunc func(voter github_com_cosmos_cosmos_sdk_types.ValAddress) (int64, bool)

	// GetTotalShareCountFunc mocks the GetTotalShareCount method.
	GetTotalShareCountFunc func() github_com_cosmos_cosmos_sdk_types.Int

	// GetTotalVoterCountFunc mocks the GetTotalVoterCount method.
	GetTotalVoterCountFunc func() int64

	// GetVoteFunc mocks the GetVote method.
	GetVoteFunc func(hash string) (types.TalliedVote, bool)

	// GetVotesFunc mocks the GetVotes method.
	GetVotesFunc func() []types.TalliedVote

	// HasVotedFunc mocks the HasVoted method.
	HasVotedFunc func(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool

	// SetMetadataFunc mocks the SetMetadata method.
	SetMetadataFunc func(metadata exported.PollMetadata)

	// SetVoteFunc mocks the SetVote method.
	SetVoteFunc func(voter github_com_cosmos_cosmos_sdk_types.ValAddress, vote types.TalliedVote)

	// calls tracks calls to the methods.
	calls struct {
		// DeletePoll holds details about calls to the DeletePoll method.
		DeletePoll []struct {
		}
		// GetPoll holds details about calls to the GetPoll method.
		GetPoll []struct {
			// Key is the key argument value.
			Key exported.PollKey
		}
		// GetShareCount holds details about calls to the GetShareCount method.
		GetShareCount []struct {
			// Voter is the voter argument value.
			Voter github_com_cosmos_cosmos_sdk_types.ValAddress
		}
		// GetTotalShareCount holds details about calls to the GetTotalShareCount method.
		GetTotalShareCount []struct {
		}
		// GetTotalVoterCount holds details about calls to the GetTotalVoterCount method.
		GetTotalVoterCount []struct {
		}
		// GetVote holds details about calls to the GetVote method.
		GetVote []struct {
			// Hash is the hash argument value.
			Hash string
		}
		// GetVotes holds details about calls to the GetVotes method.
		GetVotes []struct {
		}
		// HasVoted holds details about calls to the HasVoted method.
		HasVoted []struct {
			// Voter is the voter argument value.
			Voter github_com_cosmos_cosmos_sdk_types.ValAddress
		}
		// SetMetadata holds details about calls to the SetMetadata method.
		SetMetadata []struct {
			// Metadata is the metadata argument value.
			Metadata exported.PollMetadata
		}
		// SetVote holds details about calls to the SetVote method.
		SetVote []struct {
			// Voter is the voter argument value.
			Voter github_com_cosmos_cosmos_sdk_types.ValAddress
			// Vote is the vote argument value.
			Vote types.TalliedVote
		}
	}
	lockDeletePoll         sync.RWMutex
	lockGetPoll            sync.RWMutex
	lockGetShareCount      sync.RWMutex
	lockGetTotalShareCount sync.RWMutex
	lockGetTotalVoterCount sync.RWMutex
	lockGetVote            sync.RWMutex
	lockGetVotes           sync.RWMutex
	lockHasVoted           sync.RWMutex
	lockSetMetadata        sync.RWMutex
	lockSetVote            sync.RWMutex
}

// DeletePoll calls DeletePollFunc.
func (mock *StoreMock) DeletePoll() {
	if mock.DeletePollFunc == nil {
		panic("StoreMock.DeletePollFunc: method is nil but Store.DeletePoll was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDeletePoll.Lock()
	mock.calls.DeletePoll = append(mock.calls.DeletePoll, callInfo)
	mock.lockDeletePoll.Unlock()
	mock.DeletePollFunc()
}

// DeletePollCalls gets all the calls that were made to DeletePoll.
// Check the length with:
//     len(mockedStore.DeletePollCalls())
func (mock *StoreMock) DeletePollCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDeletePoll.RLock()
	calls = mock.calls.DeletePoll
	mock.lockDeletePoll.RUnlock()
	return calls
}

// GetPoll calls GetPollFunc.
func (mock *StoreMock) GetPoll(key exported.PollKey) exported.Poll {
	if mock.GetPollFunc == nil {
		panic("StoreMock.GetPollFunc: method is nil but Store.GetPoll was just called")
	}
	callInfo := struct {
		Key exported.PollKey
	}{
		Key: key,
	}
	mock.lockGetPoll.Lock()
	mock.calls.GetPoll = append(mock.calls.GetPoll, callInfo)
	mock.lockGetPoll.Unlock()
	return mock.GetPollFunc(key)
}

// GetPollCalls gets all the calls that were made to GetPoll.
// Check the length with:
//     len(mockedStore.GetPollCalls())
func (mock *StoreMock) GetPollCalls() []struct {
	Key exported.PollKey
} {
	var calls []struct {
		Key exported.PollKey
	}
	mock.lockGetPoll.RLock()
	calls = mock.calls.GetPoll
	mock.lockGetPoll.RUnlock()
	return calls
}

// GetShareCount calls GetShareCountFunc.
func (mock *StoreMock) GetShareCount(voter github_com_cosmos_cosmos_sdk_types.ValAddress) (int64, bool) {
	if mock.GetShareCountFunc == nil {
		panic("StoreMock.GetShareCountFunc: method is nil but Store.GetShareCount was just called")
	}
	callInfo := struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
	}{
		Voter: voter,
	}
	mock.lockGetShareCount.Lock()
	mock.calls.GetShareCount = append(mock.calls.GetShareCount, callInfo)
	mock.lockGetShareCount.Unlock()
	return mock.GetShareCountFunc(voter)
}

// GetShareCountCalls gets all the calls that were made to GetShareCount.
// Check the length with:
//     len(mockedStore.GetShareCountCalls())
func (mock *StoreMock) GetShareCountCalls() []struct {
	Voter github_com_cosmos_cosmos_sdk_types.ValAddress
} {
	var calls []struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
	}
	mock.lockGetShareCount.RLock()
	calls = mock.calls.GetShareCount
	mock.lockGetShareCount.RUnlock()
	return calls
}

// GetTotalShareCount calls GetTotalShareCountFunc.
func (mock *StoreMock) GetTotalShareCount() github_com_cosmos_cosmos_sdk_types.Int {
	if mock.GetTotalShareCountFunc == nil {
		panic("StoreMock.GetTotalShareCountFunc: method is nil but Store.GetTotalShareCount was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetTotalShareCount.Lock()
	mock.calls.GetTotalShareCount = append(mock.calls.GetTotalShareCount, callInfo)
	mock.lockGetTotalShareCount.Unlock()
	return mock.GetTotalShareCountFunc()
}

// GetTotalShareCountCalls gets all the calls that were made to GetTotalShareCount.
// Check the length with:
//     len(mockedStore.GetTotalShareCountCalls())
func (mock *StoreMock) GetTotalShareCountCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetTotalShareCount.RLock()
	calls = mock.calls.GetTotalShareCount
	mock.lockGetTotalShareCount.RUnlock()
	return calls
}

// GetTotalVoterCount calls GetTotalVoterCountFunc.
func (mock *StoreMock) GetTotalVoterCount() int64 {
	if mock.GetTotalVoterCountFunc == nil {
		panic("StoreMock.GetTotalVoterCountFunc: method is nil but Store.GetTotalVoterCount was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetTotalVoterCount.Lock()
	mock.calls.GetTotalVoterCount = append(mock.calls.GetTotalVoterCount, callInfo)
	mock.lockGetTotalVoterCount.Unlock()
	return mock.GetTotalVoterCountFunc()
}

// GetTotalVoterCountCalls gets all the calls that were made to GetTotalVoterCount.
// Check the length with:
//     len(mockedStore.GetTotalVoterCountCalls())
func (mock *StoreMock) GetTotalVoterCountCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetTotalVoterCount.RLock()
	calls = mock.calls.GetTotalVoterCount
	mock.lockGetTotalVoterCount.RUnlock()
	return calls
}

// GetVote calls GetVoteFunc.
func (mock *StoreMock) GetVote(hash string) (types.TalliedVote, bool) {
	if mock.GetVoteFunc == nil {
		panic("StoreMock.GetVoteFunc: method is nil but Store.GetVote was just called")
	}
	callInfo := struct {
		Hash string
	}{
		Hash: hash,
	}
	mock.lockGetVote.Lock()
	mock.calls.GetVote = append(mock.calls.GetVote, callInfo)
	mock.lockGetVote.Unlock()
	return mock.GetVoteFunc(hash)
}

// GetVoteCalls gets all the calls that were made to GetVote.
// Check the length with:
//     len(mockedStore.GetVoteCalls())
func (mock *StoreMock) GetVoteCalls() []struct {
	Hash string
} {
	var calls []struct {
		Hash string
	}
	mock.lockGetVote.RLock()
	calls = mock.calls.GetVote
	mock.lockGetVote.RUnlock()
	return calls
}

// GetVotes calls GetVotesFunc.
func (mock *StoreMock) GetVotes() []types.TalliedVote {
	if mock.GetVotesFunc == nil {
		panic("StoreMock.GetVotesFunc: method is nil but Store.GetVotes was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetVotes.Lock()
	mock.calls.GetVotes = append(mock.calls.GetVotes, callInfo)
	mock.lockGetVotes.Unlock()
	return mock.GetVotesFunc()
}

// GetVotesCalls gets all the calls that were made to GetVotes.
// Check the length with:
//     len(mockedStore.GetVotesCalls())
func (mock *StoreMock) GetVotesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetVotes.RLock()
	calls = mock.calls.GetVotes
	mock.lockGetVotes.RUnlock()
	return calls
}

// HasVoted calls HasVotedFunc.
func (mock *StoreMock) HasVoted(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool {
	if mock.HasVotedFunc == nil {
		panic("StoreMock.HasVotedFunc: method is nil but Store.HasVoted was just called")
	}
	callInfo := struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
	}{
		Voter: voter,
	}
	mock.lockHasVoted.Lock()
	mock.calls.HasVoted = append(mock.calls.HasVoted, callInfo)
	mock.lockHasVoted.Unlock()
	return mock.HasVotedFunc(voter)
}

// HasVotedCalls gets all the calls that were made to HasVoted.
// Check the length with:
//     len(mockedStore.HasVotedCalls())
func (mock *StoreMock) HasVotedCalls() []struct {
	Voter github_com_cosmos_cosmos_sdk_types.ValAddress
} {
	var calls []struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
	}
	mock.lockHasVoted.RLock()
	calls = mock.calls.HasVoted
	mock.lockHasVoted.RUnlock()
	return calls
}

// SetMetadata calls SetMetadataFunc.
func (mock *StoreMock) SetMetadata(metadata exported.PollMetadata) {
	if mock.SetMetadataFunc == nil {
		panic("StoreMock.SetMetadataFunc: method is nil but Store.SetMetadata was just called")
	}
	callInfo := struct {
		Metadata exported.PollMetadata
	}{
		Metadata: metadata,
	}
	mock.lockSetMetadata.Lock()
	mock.calls.SetMetadata = append(mock.calls.SetMetadata, callInfo)
	mock.lockSetMetadata.Unlock()
	mock.SetMetadataFunc(metadata)
}

// SetMetadataCalls gets all the calls that were made to SetMetadata.
// Check the length with:
//     len(mockedStore.SetMetadataCalls())
func (mock *StoreMock) SetMetadataCalls() []struct {
	Metadata exported.PollMetadata
} {
	var calls []struct {
		Metadata exported.PollMetadata
	}
	mock.lockSetMetadata.RLock()
	calls = mock.calls.SetMetadata
	mock.lockSetMetadata.RUnlock()
	return calls
}

// SetVote calls SetVoteFunc.
func (mock *StoreMock) SetVote(voter github_com_cosmos_cosmos_sdk_types.ValAddress, vote types.TalliedVote) {
	if mock.SetVoteFunc == nil {
		panic("StoreMock.SetVoteFunc: method is nil but Store.SetVote was just called")
	}
	callInfo := struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
		Vote  types.TalliedVote
	}{
		Voter: voter,
		Vote:  vote,
	}
	mock.lockSetVote.Lock()
	mock.calls.SetVote = append(mock.calls.SetVote, callInfo)
	mock.lockSetVote.Unlock()
	mock.SetVoteFunc(voter, vote)
}

// SetVoteCalls gets all the calls that were made to SetVote.
// Check the length with:
//     len(mockedStore.SetVoteCalls())
func (mock *StoreMock) SetVoteCalls() []struct {
	Voter github_com_cosmos_cosmos_sdk_types.ValAddress
	Vote  types.TalliedVote
} {
	var calls []struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
		Vote  types.TalliedVote
	}
	mock.lockSetVote.RLock()
	calls = mock.calls.SetVote
	mock.lockSetVote.RUnlock()
	return calls
}
