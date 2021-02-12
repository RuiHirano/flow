package strategy

import (
	"github.com/RuiHirano/flow/api"
)

type SUMOStrategy struct{
}

func NewSUMOStrategy() *SUMOStrategy{
	return &SUMOStrategy{
	}
}

func (p *SUMOStrategy) CalcNextAgents(agents []*api.Agent) ([]*api.Agent, error){
	return []*api.Agent{}, nil
}