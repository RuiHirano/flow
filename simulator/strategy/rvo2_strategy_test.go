package strategy

import (
	"testing"

	//"github.com/RuiHirano/flow/providers"
	rvo "github.com/RuiHirano/rvo2-go/src/rvosimulator"
	"github.com/stretchr/testify/assert"
)

func TestRVOStrategy(t *testing.T) {

	defaultParam := &RVO2Parameter{
		timeStep: 0.01,
		neighborDist: 0.01,
		maxNeighbors: 10,
		timeHorizon: 0.01,
		timeHorizonObst: 0.01,
		radius: 0.01,
		maxSpeed: 0.01,
		velocity: &rvo.Vector2{X: 0, Y: 0},
	}

	t.Run("SetParamでtimeStepがセットできている", func (t *testing.T) {
		st := NewRVO2Strategy()
		st.SetParameter(defaultParam)
		assert.Equal(t, st.param.timeStep, defaultParam.timeStep)
	})


	t.Run("setupScenarioでエージェントがセットされる", func (t *testing.T) {
		mockAgents := GetMockAgents(1)
		st := NewRVO2Strategy()
		st.SetParameter(defaultParam)
		st.setupScenario(mockAgents)
		assert.Equal(t, st.agentMap[mockAgents[0]] ==nil, false)
	})

	t.Run("getRVOAgentsでエージェントが取得される", func (t *testing.T) {
		mockAgents := GetMockAgents(1)
		st := NewRVO2Strategy()
		st.SetParameter(defaultParam)
		st.setupScenario(mockAgents)
		agents := st.getRVOAgents()
		//t.Log("agentMap", agents, mockAgents)
		assert.Equal(t, agents, mockAgents)
	})

	// CalcNextAgents
	t.Run("CalcNextAgentsでParameterがセットできていなければerrを返す", func (t *testing.T) {
		mockAgents := GetMockAgents(10)
		st := NewRVO2Strategy()
		_, err := st.CalcNextAgents(mockAgents)
		assert.NotEqual(t, err, nil)
	})

	t.Run("CalcNextAgentsでParameterがセットできてればagentsを返す", func (t *testing.T) {
		mockAgents := GetMockAgents(10)
		st := NewRVO2Strategy()
		st.SetParameter(defaultParam)
		agents, _ := st.CalcNextAgents(mockAgents)
		assert.NotEqual(t, agents, nil)
	})

	t.Run("CalcNextAgentsで計算結果によってPositionが変化する", func (t *testing.T) {
		mockAgents := GetMockAgents(1)
		st := NewRVO2Strategy()
		st.SetParameter(defaultParam)
		agents, _ := st.CalcNextAgents(mockAgents)
		//t.Log("agent", agents[0].Route.Position, mockAgents[0].Route.Position)
		assert.NotEqual(t, agents[0].Position, mockAgents[0].Position)
	})


}