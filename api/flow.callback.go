package api

import (
	sxapi "github.com/synerex/synerex_api"
	sxutil "github.com/synerex/synerex_sxutil"
)


///////////////////////////////
// callback
/////////////////////////////

type ICallback interface {
	SetAgentRequest(clt *sxutil.SXServiceClient, simMsg *sxapi.MbusMsg)
}

type Callback struct {
	//simapi *SimAPI
	ICallback
}

func NewCallback(cbif ICallback) *Callback {
	cb := &Callback{
		ICallback: cbif,
		//simapi:            simapi,
	}
	return cb
}

/*func (cb *Callback) Callback(clt *sxutil.SXServiceClient, msg *sxapi.MbusMsg) {
	go func() {
		simMsg := &SimMsg{}
		proto.Unmarshal(msg.GetCdata().GetEntity(), simMsg)
		//targets := []uint64{simMsg.GetSenderId()}
		//log.Printf("get Agent Callback %v\n", simMsg)
		msgId := simMsg.GetMsgId()
		switch simMsg.GetType() {
		case MsgType_SET_AGENT_REQUEST:
			cb.SetAgentRequest(clt, msg)
			// response
			targetId := msg.GetSenderId()
			cb.simapi.SetAgentResponse(clt, msgId, targetId)
		}
		return
	}()
}

// Agent
func (cb *Callback) SetAgentRequest(clt *sxutil.SXServiceClient, msg *sxapi.MbusMsg)  {}
*/