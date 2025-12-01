package service

import ()

type Service struct {
}

type GrpcDeps struct {
	OrderService *orderService
}

func NewService() *Service {
	return &Service{}
}
