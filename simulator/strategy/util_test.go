package strategy

import (
	"testing"

	"github.com/RuiHirano/flow/api"
	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	// GetMockAgent
	t.Run("GetMockAgentで指定した数のエージェントが帰ってくる", func (t *testing.T) {
		agents := GetMockAgents(10)
		assert.Equal(t, len(agents), 10)
	})

	t.Run("GetMockAgentで全ての要素がnilではないエージェントが返ってくる", func (t *testing.T) {
		agents := GetMockAgents(1)
		isNotNil := agents[0].Id != "" && agents[0].Route != nil
		assert.Equal(t, isNotNil, true)
	})

	// CalcDirection
	t.Run("CalcDirectionで2点間の方向を返す", func (t *testing.T) {
		point1 := &api.Vector{X: 10, Y: 20}
		point2 := &api.Vector{X: 30, Y: 60}
		direction := CalcDirection(point1, point2)
		assert.Equal(t, direction, &api.Vector{X: 1, Y: 2})
	})

	// CalcDistance
	t.Run("CalcDistanceで2点間の距離を返す", func (t *testing.T) {
		point1 := &api.Vector{X: 10, Y: 20}
		point2 := &api.Vector{X: 40, Y: 60}
		distance := CalcDistance(point1, point2)
		assert.Equal(t, distance, float64(50))
	})

	// DecideNextTransit
	t.Run("DecideNextTransitで経由地についている場合、次のtransitが帰ってくる", func (t *testing.T) {
		transitPoints := []*api.Point{&}
		position := &api.Point{X: 10, Y: 10}
		nextTransitPoint := DecideNextTransitPoint(transitPoints, position, )
		assert.Equal(t, distance, float64(50))
	})
}