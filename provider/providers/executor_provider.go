package providers

type ExecutorProvider struct{
}

func NewExecutorProvider() *ExecutorProvider{
	return &ExecutorProvider{
	}
}

func (p *ExecutorProvider) State2() string{
	return "aa"
}