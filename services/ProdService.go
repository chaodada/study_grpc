package services

import (
	"context"
	"google.golang.org/grpc"
	"net/rpc"
)

type ProdService struct {
	
}

func(p *ProdService) GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error) {

return &ProdResponse{ProdStock:20},nil
}