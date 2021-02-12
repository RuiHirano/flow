module simulator

go 1.12

replace (
	github.com/RuiHirano/flow/api => ./../api
	github.com/RuiHirano/flow/simulator/strategy => ./../simulator/strategy
	github.com/RuiHirano/flow/provider/providers => ./../provider/providers
)

require (
	github.com/RuiHirano/flow/api v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/rvo2-go/src/rvosimulator v0.0.0-20200707091306-e572a9b06cee // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/synerex/synerex_sxutil v0.6.2 // indirect
)
