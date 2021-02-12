package simulator

import (
	"fmt"

	"github.com/RuiHirano/flow/api"
	"github.com/RuiHirano/flow/simulator/strategy"
)

type SimulatorFactory struct{

}

func NewSimulatorFactory() *SimulatorFactory{
	return &SimulatorFactory{}
}

func (p *SimulatorFactory) Create(ptype api.SimulatorType) *Simulator{
	switch ptype {
		case api.SimulatorType_RVO2:
			return NewSimulator(ptype, strategy.NewRVO2Strategy())
		case api.SimulatorType_SUMO:
			return NewSimulator(ptype, strategy.NewSUMOStrategy())
		default:
			panic(fmt.Errorf("unknown simulator type"))
	}
}