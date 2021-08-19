package tss

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/axelarnetwork/axelar-core/app"
	broadcastMock "github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/broadcaster/types/mock"
	"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/tss/rpc/mock"
	"github.com/axelarnetwork/axelar-core/testutils"
	"github.com/axelarnetwork/axelar-core/testutils/rand"
)

func TestProcessNewBlockHeader(t *testing.T) {
	rpcClient := mock.ClientMock{}
	principalAddr := rand.Str(20)
	broadcaster := broadcastMock.BroadcasterMock{}
	sender := rand.Bytes(sdk.AddrLen)
	logger := log.TestingLogger()
	cdc := app.MakeEncodingConfig().Amino

	t.Run("should do nothing when the timeout queue is empty", testutils.Func(func(t *testing.T) {
		mgr := NewMgr(&rpcClient, time.Second, principalAddr, &broadcaster, sender, logger, cdc)

		mgr.ProcessNewBlockHeader(100)
	}))

	t.Run("should do nothing if first session in queue has not timed out yet", testutils.Func(func(t *testing.T) {
		mgr := NewMgr(&rpcClient, time.Second, principalAddr, &broadcaster, sender, logger, cdc)

		id := rand.Str(20)
		timeoutAt := int64(1234)

		mgr.timeoutQueue.Enqueue(id, timeoutAt)

		mgr.ProcessNewBlockHeader(timeoutAt - 1)
		assert.Len(t, mgr.timeoutQueue.queue, 1)
	}))

	t.Run("should signal every session in queue that has timed out", testutils.Func(func(t *testing.T) {
		mgr := NewMgr(&rpcClient, time.Second, principalAddr, &broadcaster, sender, logger, cdc)

		id1 := rand.Str(20)
		id2 := rand.Str(20)
		timeoutAt := int64(1234)

		session1 := mgr.timeoutQueue.Enqueue(id1, timeoutAt)
		session2 := mgr.timeoutQueue.Enqueue(id2, timeoutAt)
		mgr.timeoutQueue.Enqueue(rand.Str(20), timeoutAt+1)

		mgr.ProcessNewBlockHeader(timeoutAt)

		assert.Len(t, mgr.timeoutQueue.queue, 1)
		assert.Panics(t, func() { close(session1.timeout) })
		assert.Panics(t, func() { close(session2.timeout) })
	}))
}
