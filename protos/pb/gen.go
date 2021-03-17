package pb

//go:generate protoc -I. -I/usr/local/include --go_out=plugins=grpc:. --go_opt=paths=source_relative cltypes.proto clservice.proto
