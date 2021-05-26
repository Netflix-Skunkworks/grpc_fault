package service

import "google.golang.org/grpc"

type fakeGRPCServer struct {
	*grpc.Server
}



