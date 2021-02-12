module strategy

go 1.12

replace github.com/RuiHirano/flow/api => ./../../api

require (
	github.com/RuiHirano/flow/api v0.0.0-00010101000000-000000000000
	github.com/RuiHirano/rvo2-go/src/rvosimulator v0.0.0-20200707091306-e572a9b06cee
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/uuid v1.2.0
	github.com/stretchr/testify v1.7.0
)
