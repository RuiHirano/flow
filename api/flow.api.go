package api

type FlowAPI struct{

}

func NewFlowAPI() *FlowAPI{
	return &FlowAPI{}
}

func (p *FlowAPI) ConnectSynerexServer() string{
	return "SUCCESS"
}