package test

import (
	"testing"

	"github.com/RuiHirano/flow/simulator"

	//"github.com/RuiHirano/flow/providers"
	"github.com/RuiHirano/flow/api"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

/*func TestProvider(t *testing.T) {

	t.Run("Connect Synerex Server", func (t *testing.T) {
		prov := provider.NewProvider()
		res := prov.ConnectSynerexServer()
		assert.Equal(t, res, "SUCCESS")
	})

}*/

/*func TestFlowAPI(t *testing.T) {

	t.Run("Connect Synerex Server", func (t *testing.T) {
		api := flowapi.NewFlowAPI()
		res := api.ConnectSynerexServer()
		assert.Equal(t, res, "SUCCESS")
	})

}*/


func TestAgentProvider(t *testing.T) {

	/*t.Run("Create Agent Provider", func (t *testing.T) {
		ptype := api.ProviderType_AGENT
		factory := provider.NewProviderFactory()
		prov := factory.Create(ptype)
		assert.Equal(t, prov.Type(), api.ProviderType_AGENT)
	})*/

	/*t.Run("Connect Provider to Synerex Server", func (t *testing.T) {
		ptype := api.ProviderType_AGENT
		factory := provider.NewProviderFactory()
		provider := factory.Create(ptype)
		serverAddr := "127.0.0.1:10000"
		_, err := provider.ConnectSynerexServer(serverAddr)
		assert.Equal(t, err, nil)
	})*/
}


func Test(t *testing.T) {

	/*t.Run("Create Agent Provider", func (t *testing.T) {
		// Simulator
		stype := api.SimulatorType_RVO2
		factory := simulator.NewProviderFactory()
		sim := factory.Create(stype)
		// Provider Interface
		ptype := api.ProviderType_AGENT
		provFactory := provider.NewProviderFactory()
		iProv := provFactory.Create(ptype)
		// API Interface
		apiFactory := provider.NewAPIFactory()
		iApi := apiFactory.Create(ptype)

		prov := provider.NewProvider(stype, iProv, iApi)
		assert.Equal(t, sim.Type(), prov.type)
	})*/
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