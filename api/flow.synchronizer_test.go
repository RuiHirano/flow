package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSynchronizer(t *testing.T) {

	t.Run("isFinishSyncで揃っていればtrueを返す", func (t *testing.T) {
		syn := NewSynchronizer()
		targets := []uint64{1}
		msgId := uint64(0)
		flowMsg := &FlowMsg{
			MsgId: uint64(0),
			SenderId: uint64(1),
			Type: MsgType_SET_AGENT_REQUEST,
		}
		syn.syncChMap[msgId] = make(chan *FlowMsg, 30)
		syn.msgMap[msgId] = []*FlowMsg{flowMsg}
		isFinish := syn.isFinishSync(msgId, targets)
		assert.Equal(t, isFinish, true)
	})

	t.Run("isFinishSyncで揃っていなければfalseを返す", func (t *testing.T) {
		syn := NewSynchronizer()
		targets := []uint64{1}
		msgId := uint64(0)
		flowMsg := &FlowMsg{
			MsgId: uint64(0),
			SenderId: uint64(2),
			Type: MsgType_SET_AGENT_REQUEST,
		}
		syn.syncChMap[msgId] = make(chan *FlowMsg, 30)
		syn.msgMap[msgId] = []*FlowMsg{flowMsg}
		isFinish := syn.isFinishSync(msgId, targets)
		assert.Equal(t, isFinish, false)
	})

	t.Run("PublishMsgでchがなければerrorを返す", func (t *testing.T) {
		syn := NewSynchronizer()
		flowMsg := &FlowMsg{
			MsgId: uint64(0),
			SenderId: uint64(1),
			Type: MsgType_SET_AGENT_REQUEST,
		}
		err := syn.PublishMsg(flowMsg)
		assert.Equal(t, err, fmt.Errorf("Error: not found syncCh"))
	})

	/*t.Run("SyncMsgsでtargetsがくるまで待機する", func (t *testing.T) {
		syn := NewSynchronizer()
		targets := []uint64{1}
		msgId := uint64(0)
		flowMsg := &FlowMsg{
			MsgId: uint64(0),
			SenderId: uint64(1),
			Type: MsgType_SET_AGENT_REQUEST,
		}
		syn.PublishMsg(flowMsg)
		msgs, _ := syn.SyncMsgs(msgId, targets, 1000)
		assert.Equal(t, msgs[0], flowMsg)
	})*/

}
