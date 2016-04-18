package idl

//go:generate protoc -I "../../../idl" --go_out=plugins=grpc:. ../../../idl/hwgrpc.proto
