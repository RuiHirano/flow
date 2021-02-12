module test

go 1.12

require (
	github.com/RuiHirano/flow/api v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/flow/provider v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/flow/simulator v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/uuid v1.2.0
	github.com/stretchr/testify v1.7.0
)

replace (
	github.com/RuiHirano/flow/api => ./../api
	github.com/RuiHirano/flow/provider => ./../provider
	github.com/RuiHirano/flow/provider/providers => ./../provider/providers
	github.com/RuiHirano/flow/simulator => ./../simulator
)
