package api

import (
	//"context"
	//"fmt"
	//"log"
	//"time"

	"log"

	sxapi "github.com/synerex/synerex_api"
	sxutil "github.com/synerex/synerex_sxutil"
)

type FlowAPI struct{
	sclient *sxutil.SXServiceClient
	synerexAddress string
	nodeAddress string
}

func NewFlowAPI() *FlowAPI{
	return &FlowAPI{}
}

func (api *FlowAPI) ConnectServer(nodeAddr string, iCallback ICallback) error{
	synerexAddress, _:= api.connectNodeServer(nodeAddr)
	err := api.connectSynerexServer(synerexAddress)
	if err != nil {
		log.Fatal("Error on Connect Server ", err)
		return err
	}
	return nil
}

// NodeServerに接続
func (api *FlowAPI) connectNodeServer(nodeAddr string) (string, error){
	api.nodeAddress = nodeAddr
	chTypes := []uint32{uint32(0)}
	sxServerAddress , rerr := sxutil.RegisterNode(nodeAddr, "Provider", chTypes, nil)
	if rerr != nil {
		log.Fatal("Can't register node ", rerr)
	}
	log.Printf("Connecting SynerexServer at [%s]\n",sxServerAddress)
	return sxServerAddress, nil
}

// SynerexServerに接続
func (api *FlowAPI) connectSynerexServer(synerexAddr string) error{
	/*api.synerexAddress = synerexAddr
	client := sxutil.GrpcConnectServer(synerexAddr) // if there is server address change, we should do it!
	argJson := fmt.Sprintf("{Client:%s}")
	chType := uint32(0)
	api.sclient = sxutil.NewSXServiceClient(client, chType, argJson)
	// SubscribeMbus
	go api.subscribeMbus(api.sclient, callback)*/
	return nil
}

// MbusCallbackを登録
func (api *FlowAPI) subscribeMbus(sclient *sxutil.SXServiceClient, mbcb func(*sxutil.SXServiceClient, *sxapi.MbusMsg)) {
	//called as goroutine
	/*ctx := context.Background() // should check proper context
	sxutil.RegisterDeferFunction(func() {
		sclient.CloseMbus(ctx)
	})
	for {
		sclient.SubscribeMbus(ctx, mbcb)
		// comes here if channel closed
		time.Sleep(1 * time.Second)
	}*/
}

