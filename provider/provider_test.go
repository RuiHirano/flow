package provider

import (
	"testing"

	"github.com/RuiHirano/flow/api"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {

	
	t.Run("AgentProviderをfactoryで作成できるか", func (t *testing.T) {
		ptype := api.ProviderType_AGENT
		factory := NewProviderFactory()
		provider := factory.Create(ptype)
		assert.Equal(t, provider.Type(), ptype)
	})

	// Connect Synerex Server
	t.Run("Synerexが存在する時、接続に成功する", func (t *testing.T) {
		ptype := api.ProviderType_AGENT
		factory := NewProviderFactory()
		provider := factory.Create(ptype)
		err := provider.papi.ConnectSynerexServer()
		assert.Equal(t, err, nil)

	})

	t.Run("Synerexが存在しない時、エラーを発する", func (t *testing.T) {
	})
}
