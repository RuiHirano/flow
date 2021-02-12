module providers

go 1.12

replace (
	github.com/RuiHirano/flow/api => ./../../api
	github.com/RuiHirano/flow/simulator => ./../../simulator
	github.com/RuiHirano/flow/provider => ./../../provider
)

require (
	github.com/RuiHirano/flow/api v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/flow/simulator v0.0.0-00010101000000-000000000000 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
)
