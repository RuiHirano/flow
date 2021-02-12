package simulator

import (
	"github.com/RuiHirano/flow/api"
	"github.com/google/uuid"
)
type ISimulator interface{
	//State() string // providerのStateを返す
	CalcNextAgents(agents []*api.Agent) ([]*api.Agent, error)
}

type Simulator struct{
	id string   // simulatorID
	stype api.SimulatorType  // SUMO or RVO2
	ISimulator
}

func NewSimulator(stype api.SimulatorType, iprov ISimulator) *Simulator{
	return &Simulator{
		id: uuid.New().String(),
		stype: stype,
		ISimulator: iprov,
	}
}

func (p *Simulator) Type() api.SimulatorType{
	return p.stype
}



