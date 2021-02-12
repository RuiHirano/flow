package strategy


import (
	"fmt"

	"github.com/RuiHirano/flow/api"
	rvo "github.com/RuiHirano/rvo2-go/src/rvosimulator"
)

type RVO2Parameter struct{
	timeStep float64
	neighborDist float64
	maxNeighbors int
	timeHorizon float64
	timeHorizonObst float64
	radius float64
	maxSpeed float64
	velocity *rvo.Vector2
}

type RVO2Strategy struct{
	sim   *rvo.RVOSimulator
	param  *RVO2Parameter
	agentMap map[*api.Agent]*rvo.Agent   // key: flow agent, value: rvo2 agent
}

func NewRVO2Strategy() *RVO2Strategy{
	return &RVO2Strategy{
		agentMap: map[*api.Agent]*rvo.Agent{},
	}
}

func (p *RVO2Strategy) SetParameter(param *RVO2Parameter){
	p.sim = rvo.NewRVOSimulator(
		param.timeStep, 
		param.neighborDist, 
		param.maxNeighbors,
		param.timeHorizon,
		param.timeHorizonObst,
		param.radius,
		param.maxSpeed,
		param.velocity,
	)
	p.param = param
}

func (p *RVO2Strategy) CalcNextAgents(agents []*api.Agent) ([]*api.Agent, error){
	if p.sim == nil{
		return nil, fmt.Errorf("parameter is not set")
	}
	p.setupScenario(agents)
	p.sim.DoStep()
	nextAgents := p.getRVOAgents()
	return nextAgents, nil
}

// SetupScenario: Scenarioを設定する関数
func (p *RVO2Strategy) setupScenario(agents []*api.Agent) {
	// Set Agent
	for _, agent := range agents {

		position := &rvo.Vector2{X: agent.Position.X, Y: agent.Position.Y}
		goal := &rvo.Vector2{X: agent.Route.NextTransit.Point.X, Y: agent.Route.NextTransit.Point.Y}

		// Agentを追加
		id, _ := p.sim.AddDefaultAgent(position)
		rvoAgent := p.sim.GetAgent(id)
		p.agentMap[agent] = rvoAgent
		

		// 目的地を設定
		p.sim.SetAgentGoal(id, goal)

		// エージェントの速度方向ベクトルを設定
		goalVector := p.sim.GetAgentGoalVector(id)
		p.sim.SetAgentPrefVelocity(id, goalVector)
		//sim.SetAgentMaxSpeed(id, float64(api.MaxSpeed))
	}
}

/*func (p *RVO2Strategy) getRVOAgent(id string) []*api.Agent{
	rvoID := p.idMap[id]
	rvoAgent := p.sim.GetAgent(rvoID)
	
	return agents
}

func (p *RVO2Strategy) convertRVOAgentToFlowAgent(rvoAgent) []*api.Agent{
	rvoID := p.idMap[id]
	rvoAgent := p.sim.GetAgent(rvoID)
	return agents
}*/

// SetupScenario: Scenarioを設定する関数
func (p *RVO2Strategy) getRVOAgents() []*api.Agent{
	agents := make([]*api.Agent, 0)
	currentAgents := make([]*api.Agent, 0, len(p.agentMap))
    for k := range p.agentMap {
        currentAgents = append(currentAgents, k)
    }
	for _, currentAgent := range currentAgents {
		rvoId := p.agentMap[currentAgent].ID
		// 管理エリア内のエージェントのみ抽出
		//position := agentInfo.Route.Position

		// rvoの位置情報を緯度経度に変換する
		rvoAgentPosition := p.sim.GetAgentPosition(rvoId)

		position := &api.Vector{
			Y:  rvoAgentPosition.Y,
			X: rvoAgentPosition.X,
		}

		// 現在の位置とゴールとの距離と角度を求める (度, m))
		//_, distance := rvo2route.CalcDirectionAndDistance(nextCoord, agentInfo.Route.NextTransit)
		// 次の経由地nextTransitを求める
		//nextTransit := rvo2route.DecideNextTransit(agentInfo.Route.NextTransit, agentInfo.Route.TransitPoints, distance, destination)
		//nextTransit := agentInfo.Route.NextTransit

		//nextTransit := rvo2route.DecideNextTransit(position, agentInfo.Route.NextTransit ,agentInfo.Route.TransitPoints)
		//goalVector := sim.GetAgentGoalVector(int(rvoId))
		//direction := math.Atan2(goalVector.Y, goalVector.X)
		//speed := agentInfo.Route.Speed

		/*nextRoute := &api.Route{
			Position:      nextCoord,
			Direction:     direction,
			Speed:         speed,
			TransitPoints: agentInfo.Route.TransitPoints,
			NextTransit:   nextTransit,
		}*/

		nextRoute := &api.Route{
			TransitPoints: currentAgent.Route.TransitPoints,
			NextTransit:   currentAgent.Route.NextTransit,
		}

		agent := &api.Agent{
			Id:    currentAgent.Id,
			Type:  currentAgent.Type,
			Position: position,
			Velocity: currentAgent.Velocity,
			Route: nextRoute,
			Data: currentAgent.Data,
		}

		agents = append(agents, agent)
	
	}
	return agents
}