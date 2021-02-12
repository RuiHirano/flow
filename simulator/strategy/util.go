package strategy


import (
	"math"
	"math/rand"

	"github.com/RuiHirano/flow/api"
	"github.com/google/uuid"
)

func GetMockAgents(agentNum uint64) []*api.Agent {
	agents := make([]*api.Agent, 0)
	minLon, maxLon, minLat, maxLat := 136.971626, 136.989379, 35.152210, 35.161499
	for i := 0; i < int(agentNum); i++ {
		position := &api.Vector{
			X: minLon + (maxLon-minLon)*rand.Float64(),
			Y:  minLat + (maxLat-minLat)*rand.Float64(),
		}
		transitPoint := &api.TransitPoint{
			Id: uuid.New().String(),
			Point:&api.Vector{
				X: minLon + (maxLon-minLon)*rand.Float64(),
				Y:  minLat + (maxLat-minLat)*rand.Float64(),
			},
		}
		transitPoints := []*api.TransitPoint{transitPoint}
		agents = append(agents, &api.Agent{
			Type: api.AgentType_PERSON,
			Id:   uuid.New().String(),
			Position:      position,
			Velocity:     &api.Vector{X: 10, Y: 10},
			Route: &api.Route{
				TransitPoints: transitPoints,
				NextTransit:   transitPoint,
			},
		})
	}
	return agents
}

type Agent2 struct{
	*api.Agent
}

// Point1からpoint2への方向を返す
func CalcDirection(point1 *api.Vector, point2 *api.Vector) *api.Vector{
	dx := point2.X - point1.X
	dy := point2.Y - point1.Y
	direction := &api.Vector{X:1, Y: dy/dx} // xを1で正規化
	return direction
}

// Point1からpoint2への距離を返す
func CalcDistance(point1 *api.Vector, point2 *api.Vector) float64{
	dx := point2.X - point1.X
	dy := point2.Y - point1.Y
	distance := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	return distance
}
