package api

//go:generate protoc -I=. -I=$GOPATH/src --gogo_out=plugins=grpc:. deps.proto
//go:generate protoc -I=. -I=$GOPATH/src --gogo_out=plugins=grpc:. extractor.proto
