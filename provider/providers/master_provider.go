package providers

type MasterProvider struct{
}

func NewMasterProvider() *MasterProvider{
	return &MasterProvider{
	}
}

func (p *MasterProvider) State2() string{
	return "aa"
}