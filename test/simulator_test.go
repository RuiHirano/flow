package test

import (
	"testing"

	"github.com/RuiHirano/flow/simulator"

	//"github.com/RuiHirano/flow/providers"
	"github.com/RuiHirano/flow/api"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSimulator(t *testing.T) {

	t.Run("Create RVO2 Simulator", func (t *testing.T) {
		stype := api.SimulatorType_RVO2
		factory := simulator.NewSimulatorFactory()
		sim := factory.Create(stype)
		assert.Equal(t, sim.Type(), api.SimulatorType_RVO2)
	})

	t.Run("Calculate Next Agents", func (t *testing.T) {
		mockAgents := GetMockAgents(10)
		stype := api.SimulatorType_RVO2
		factory := simulator.NewSimulatorFactory()
		sim := factory.Create(stype)
		sim.SetParameter(&simulator.RVO2Parameter{
			timeStep: 0.1,
	neighborDist: 0.1,
	maxNeighbors: 0.1,
	timeHorizon: 0.1,
	timeHorizonObst: 0.1,
	radius: 0.1,
	maxSpeed: 0.1,
		})
		agents := sim.CalcNextAgents(mockAgents)
		assert.Equal(t, len(agents), 10)
	})
}


func GetMockAgents(agentNum uint64) []*api.Agent {
	uid, _ := uuid.NewRandom()
	agents := make([]*api.Agent, 0)
	for i := 0; i < int(agentNum); i++ {
		agents = append(agents, &api.Agent{
			Id:   uint64(uid.ID()),
		})
	}
	return agents
}