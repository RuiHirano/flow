package providers

import (
	"github.com/RuiHirano/flow/api"
	sim "github.com/RuiHirano/flow/simulator"
)

type AgentProvider struct{
	agents []*api.Agent
	simulator *sim.Simulator
}

func NewAgentProvider() *AgentProvider{
	return &AgentProvider{
		agents: []*api.Agent{},
	}
}

func (ap *AgentProvider) SetSimulator(simulator *sim.Simulator) {
	ap.simulator = simulator
}

func (ap *AgentProvider) Agents() []*api.Agent{
	return ap.agents
}

func (ap *AgentProvider) ForwardStep() {
	sameAgents := ap.getSameAreaAgents()
	ap.calculateAgents(sameAgents)
	neighborAgents := ap.getNeighborAreaAgents()
	ap.updateAgents(neighborAgents)
}

// 同エリアの異種エージェントを取得
func (ap *AgentProvider) getSameAreaAgents() []*api.Agent{
	sameAgents := []*api.Agent{}
	return sameAgents
}

// 次ステップのAgentsを計算
func (ap *AgentProvider) calculateAgents(sameAgents []*api.Agent) {
	agents := append(ap.agents, sameAgents...)
	nextAgents, _ := ap.simulator.CalcNextAgents(agents)
	ap.agents = nextAgents
}

// 隣接エリアの同種エージェントを取得
func (ap *AgentProvider) getNeighborAreaAgents() []*api.Agent{
	neighborAgents := []*api.Agent{}
	return neighborAgents
}

// Agentsを更新
func (ap *AgentProvider) updateAgents(neighborAgents []*api.Agent) {
}
