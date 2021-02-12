package api

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
)

///////////////////////////////////////////
/////////////      Synchronizer      ////////////
//////////////////////////////////////////

type Synchronizer struct {
	syncChMap map[uint64]chan *FlowMsg
	msgMap    map[uint64][]*FlowMsg
}

func NewSynchronizer() *Synchronizer {
	w := &Synchronizer{
		syncChMap: make(map[uint64]chan *FlowMsg),
		msgMap:    make(map[uint64][]*FlowMsg),
	}
	return w
}

func (w *Synchronizer) SyncMsgs(msgId uint64, targets []uint64, timeout uint64) ([]*FlowMsg, error) {

	mu.Lock()
	bufSize := 30
	syncCh := make(chan *FlowMsg, bufSize)
	w.syncChMap[msgId] = syncCh
	w.msgMap[msgId] = make([]*FlowMsg, 0)
	mu.Unlock()

	wg := sync.WaitGroup{}
	wg.Add(1)
	var err error
	go func() {
		for {
			select {
			case flowMsg, _ := <-syncCh:
				mu.Lock()
				// spのidがidListに入っているか
				if flowMsg.GetMsgId() == msgId {
					w.msgMap[msgId] = append(w.msgMap[msgId], flowMsg)
					// 同期が終了したかどうか
					if w.isFinishSync(msgId, targets) {
						mu.Unlock()
						wg.Done()
						return
					}
				}
				mu.Unlock()
			}
		}
	}()
	wg.Wait()
	mu.Lock()
	msgs := w.msgMap[msgId]
	// delete key value
	delete(w.msgMap, msgId)
	delete(w.syncChMap, msgId)
	mu.Unlock()
	return msgs, err
}

func (w *Synchronizer) PublishMsg(flowMsg *FlowMsg) error{
	mu.Lock()
	syncCh, ok := w.syncChMap[flowMsg.GetMsgId()]
	mu.Unlock()
	if ok {
		syncCh <- flowMsg
		return nil
	}
	return fmt.Errorf("Error: not found syncCh")
}

func (w *Synchronizer) isFinishSync(msgId uint64, targets []uint64) bool {
	for _, target := range targets {
		//log.Printf("\nIsFini2\n")
		targetId := target
		isExist := false
		for _, flowMsg := range w.msgMap[msgId] {
			if targetId == flowMsg.GetSenderId() {
				isExist = true
			}
		}
		if isExist == false {
			return false
		}
	}
	return true
}