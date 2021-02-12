package provider

import (
	//"context"
	//"fmt"
	//"log"
	//"time"

	"github.com/RuiHirano/flow/api"
	"github.com/google/uuid"
	sxutil "github.com/synerex/synerex_sxutil"
)
type IProvider interface{
	//State() string // providerのStateを返す
	//State2() string
}

type Provider struct{
	id string   // providerID
	ptype api.ProviderType  // AGENT or GATEWAY or MASTER or EXECUTOR
	synerexAddress string
	nodeAddress string
	sclient *sxutil.SXServiceClient
	papi *api.FlowAPI
	IProvider
}

func NewProvider(ptype api.ProviderType, iprov IProvider) *Provider{
	//apiFactory := api.NewFlowAPIFactory()
	return &Provider{
		id: uuid.New().String(),
		ptype: ptype,
		IProvider: iprov,
		papi: apiFactory.create(ptype),
	}
}

func (p *Provider) Run() error{
	//p.synerexAddress = p.papi.ConnectNodeServer()
	//p.sclient = p.papi.ConnectSynerexServer(p.synerexAddress)
	return nil
}

func (p *Provider) Type() api.ProviderType{
	return p.ptype
}

func (p *Provider) ID() string{
	return p.id
}

func (p *Provider) SynerexAddress() string{
	return p.synerexAddress
}

func (p *Provider) NodeAddress() string{
	return p.nodeAddress
}

/*func (p *Provider) ConnectSynerexServer(addr string, callback func(*sxutil.SXServiceClient, *sxapi.MbusMsg)) error{
	p.synerexAddress = addr
	client := sxutil.GrpcConnectServer(addr) // if there is server address change, we should do it!
	argJson := fmt.Sprintf("{Client:%s}", p.ptype)
	chType := uint32(0)
	p.sclient = sxutil.NewSXServiceClient(client, chType, argJson)
	// SubscribeMbus
	go p.subscribeMbusLoop(p.sclient, callback)
	return nil
}

// MbusにSubscribeする
func (p *Provider) subscribeMbusLoop(sclient *sxutil.SXServiceClient, mbcb func(*sxutil.SXServiceClient, *sxapi.MbusMsg)) {
	//called as goroutine
	ctx := context.Background() // should check proper context
	sxutil.RegisterDeferFunction(func() {
		sclient.CloseMbus(ctx)
	})
	for {
		sclient.SubscribeMbus(ctx, mbcb)
		// comes here if channel closed
		time.Sleep(1 * time.Second)
	}
}

func (p *Provider) ConnectNodeServer(addr string) error{
	p.nodeAddress = addr
	chTypes := []uint32{uint32(0)}
	sxServerAddress , rerr := sxutil.RegisterNode(addr, "Provider", chTypes, nil)
	if rerr != nil {
		log.Fatal("Can't register node ", rerr)
	}
	log.Printf("Connecting SynerexServer at [%s]\n",sxServerAddress)
	return nil
}*/



