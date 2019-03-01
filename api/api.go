package api

//go:generate protoc -I=. -I=$GOPATH/src --gogo_out=:. deps.proto
//go:generate protoc -I=. -I=$GOPATH/src --gogo_out=:. extractor.proto
