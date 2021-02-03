// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/axelarnetwork/axelar-core/x/bitcoin/types"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"sync"
)

// Ensure, that RPCClientMock does implement types.RPCClient.
// If this is not the case, regenerate this file with moq.
var _ types.RPCClient = &RPCClientMock{}

// RPCClientMock is a mock implementation of types.RPCClient.
//
// 	func TestSomethingThatUsesRPCClient(t *testing.T) {
//
// 		// make and configure a mocked types.RPCClient
// 		mockedRPCClient := &RPCClientMock{
// 			GetOutPointInfoFunc: func(out *wire.OutPoint) (types.OutPointInfo, error) {
// 				panic("mock out the GetOutPointInfo method")
// 			},
// 			ImportAddressRescanFunc: func(address string, account string, rescan bool) error {
// 				panic("mock out the ImportAddressRescan method")
// 			},
// 			NetworkFunc: func() types.Network {
// 				panic("mock out the Network method")
// 			},
// 			SendRawTransactionFunc: func(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error) {
// 				panic("mock out the SendRawTransaction method")
// 			},
// 		}
//
// 		// use mockedRPCClient in code that requires types.RPCClient
// 		// and then make assertions.
//
// 	}
type RPCClientMock struct {
	// GetOutPointInfoFunc mocks the GetOutPointInfo method.
	GetOutPointInfoFunc func(out *wire.OutPoint) (types.OutPointInfo, error)

	// ImportAddressRescanFunc mocks the ImportAddressRescan method.
	ImportAddressRescanFunc func(address string, account string, rescan bool) error

	// NetworkFunc mocks the Network method.
	NetworkFunc func() types.Network

	// SendRawTransactionFunc mocks the SendRawTransaction method.
	SendRawTransactionFunc func(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetOutPointInfo holds details about calls to the GetOutPointInfo method.
		GetOutPointInfo []struct {
			// Out is the out argument value.
			Out *wire.OutPoint
		}
		// ImportAddressRescan holds details about calls to the ImportAddressRescan method.
		ImportAddressRescan []struct {
			// Address is the address argument value.
			Address string
			// Account is the account argument value.
			Account string
			// Rescan is the rescan argument value.
			Rescan bool
		}
		// Network holds details about calls to the Network method.
		Network []struct {
		}
		// SendRawTransaction holds details about calls to the SendRawTransaction method.
		SendRawTransaction []struct {
			// Tx is the tx argument value.
			Tx *wire.MsgTx
			// AllowHighFees is the allowHighFees argument value.
			AllowHighFees bool
		}
	}
	lockGetOutPointInfo     sync.RWMutex
	lockImportAddressRescan sync.RWMutex
	lockNetwork             sync.RWMutex
	lockSendRawTransaction  sync.RWMutex
}

// GetOutPointInfo calls GetOutPointInfoFunc.
func (mock *RPCClientMock) GetOutPointInfo(out *wire.OutPoint) (types.OutPointInfo, error) {
	if mock.GetOutPointInfoFunc == nil {
		panic("RPCClientMock.GetOutPointInfoFunc: method is nil but RPCClient.GetOutPointInfo was just called")
	}
	callInfo := struct {
		Out *wire.OutPoint
	}{
		Out: out,
	}
	mock.lockGetOutPointInfo.Lock()
	mock.calls.GetOutPointInfo = append(mock.calls.GetOutPointInfo, callInfo)
	mock.lockGetOutPointInfo.Unlock()
	return mock.GetOutPointInfoFunc(out)
}

// GetOutPointInfoCalls gets all the calls that were made to GetOutPointInfo.
// Check the length with:
//     len(mockedRPCClient.GetOutPointInfoCalls())
func (mock *RPCClientMock) GetOutPointInfoCalls() []struct {
	Out *wire.OutPoint
} {
	var calls []struct {
		Out *wire.OutPoint
	}
	mock.lockGetOutPointInfo.RLock()
	calls = mock.calls.GetOutPointInfo
	mock.lockGetOutPointInfo.RUnlock()
	return calls
}

// ImportAddressRescan calls ImportAddressRescanFunc.
func (mock *RPCClientMock) ImportAddressRescan(address string, account string, rescan bool) error {
	if mock.ImportAddressRescanFunc == nil {
		panic("RPCClientMock.ImportAddressRescanFunc: method is nil but RPCClient.ImportAddressRescan was just called")
	}
	callInfo := struct {
		Address string
		Account string
		Rescan  bool
	}{
		Address: address,
		Account: account,
		Rescan:  rescan,
	}
	mock.lockImportAddressRescan.Lock()
	mock.calls.ImportAddressRescan = append(mock.calls.ImportAddressRescan, callInfo)
	mock.lockImportAddressRescan.Unlock()
	return mock.ImportAddressRescanFunc(address, account, rescan)
}

// ImportAddressRescanCalls gets all the calls that were made to ImportAddressRescan.
// Check the length with:
//     len(mockedRPCClient.ImportAddressRescanCalls())
func (mock *RPCClientMock) ImportAddressRescanCalls() []struct {
	Address string
	Account string
	Rescan  bool
} {
	var calls []struct {
		Address string
		Account string
		Rescan  bool
	}
	mock.lockImportAddressRescan.RLock()
	calls = mock.calls.ImportAddressRescan
	mock.lockImportAddressRescan.RUnlock()
	return calls
}

// Network calls NetworkFunc.
func (mock *RPCClientMock) Network() types.Network {
	if mock.NetworkFunc == nil {
		panic("RPCClientMock.NetworkFunc: method is nil but RPCClient.Network was just called")
	}
	callInfo := struct {
	}{}
	mock.lockNetwork.Lock()
	mock.calls.Network = append(mock.calls.Network, callInfo)
	mock.lockNetwork.Unlock()
	return mock.NetworkFunc()
}

// NetworkCalls gets all the calls that were made to Network.
// Check the length with:
//     len(mockedRPCClient.NetworkCalls())
func (mock *RPCClientMock) NetworkCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockNetwork.RLock()
	calls = mock.calls.Network
	mock.lockNetwork.RUnlock()
	return calls
}

// SendRawTransaction calls SendRawTransactionFunc.
func (mock *RPCClientMock) SendRawTransaction(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error) {
	if mock.SendRawTransactionFunc == nil {
		panic("RPCClientMock.SendRawTransactionFunc: method is nil but RPCClient.SendRawTransaction was just called")
	}
	callInfo := struct {
		Tx            *wire.MsgTx
		AllowHighFees bool
	}{
		Tx:            tx,
		AllowHighFees: allowHighFees,
	}
	mock.lockSendRawTransaction.Lock()
	mock.calls.SendRawTransaction = append(mock.calls.SendRawTransaction, callInfo)
	mock.lockSendRawTransaction.Unlock()
	return mock.SendRawTransactionFunc(tx, allowHighFees)
}

// SendRawTransactionCalls gets all the calls that were made to SendRawTransaction.
// Check the length with:
//     len(mockedRPCClient.SendRawTransactionCalls())
func (mock *RPCClientMock) SendRawTransactionCalls() []struct {
	Tx            *wire.MsgTx
	AllowHighFees bool
} {
	var calls []struct {
		Tx            *wire.MsgTx
		AllowHighFees bool
	}
	mock.lockSendRawTransaction.RLock()
	calls = mock.calls.SendRawTransaction
	mock.lockSendRawTransaction.RUnlock()
	return calls
}
