package provider

import (
	"fmt"

	"github.com/RuiHirano/flow/api"
	"github.com/RuiHirano/flow/provider/providers"
)

type ProviderFactory struct{

}

func NewProviderFactory() *ProviderFactory{
	return &ProviderFactory{}
}

func (p *ProviderFactory) Create(ptype api.ProviderType) *Provider{
	switch ptype {
		case api.ProviderType_AGENT:
			return NewProvider(ptype, providers.NewAgentProvider())
		case api.ProviderType_MASTER:
			return NewProvider(ptype, providers.NewMasterProvider())
		case api.ProviderType_EXECUTOR:
			return NewProvider(ptype, providers.NewExecutorProvider())
		case api.ProviderType_GATEWAY:
			return NewProvider(ptype, providers.NewGatewayProvider())
		default:
			panic(fmt.Errorf("unknown provider type"))
	}
}