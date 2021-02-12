package providers

type GatewayProvider struct{
}

func NewGatewayProvider() *GatewayProvider{
	return &GatewayProvider{
	}
}

func (p *GatewayProvider) State2() string{
	return "aa"
}