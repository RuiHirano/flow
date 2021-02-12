module provider

go 1.12

replace (
	github.com/RuiHirano/flow/api => ./../api
	github.com/RuiHirano/flow/provider/providers => ./../provider/providers
	github.com/RuiHirano/flow/simulator => ./../simulator
	github.com/RuiHirano/flow/simulator/strategy => ./../simulator/strategy
)

require (
	github.com/RuiHirano/flow/api v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/flow/provider/providers v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/flow/simulator v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/flow/simulator/strategy v0.0.0-00010101000000-000000000000 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/uuid v1.2.0
	github.com/stretchr/testify v1.7.0
	github.com/synerex/synerex_api v0.4.2
	github.com/synerex/synerex_sxutil v0.6.2
)
