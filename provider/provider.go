package provider

import (
	//"context"
	//"fmt"
	//"log"
	//"time"

	"github.com/RuiHirano/flow/api"
	"github.com/google/uuid"
)
type IProvider interface{
	//State() string // providerのStateを返す
	//State2() string
}

type Provider struct{
	id string   // providerID
	ptype api.ProviderType  // AGENT or GATEWAY or MASTER or EXECUTOR
	papi *api.FlowAPI
	IProvider
}

func NewProvider(ptype api.ProviderType, iprov IProvider) *Provider{
	flowapi := api.NewFlowAPI()
	return &Provider{
		id: uuid.New().String(),
		ptype: ptype,
		IProvider: iprov,
		papi: flowapi,
	}
}

func (p *Provider) Run(nodeAddr string) error{
	err := p.papi.ConnectServer(nodeAddr)
	if err != nil {
		log.Fatal("Can't Run Provider ", err)
	}
	return nil
}

func (p *Provider) Type() api.ProviderType{
	return p.ptype
}

func (p *Provider) ID() string{
	return p.id
}


