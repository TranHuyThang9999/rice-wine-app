package interfaces

import (
	"context"
	"rice-wine-shop/core/generator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderClientService struct {
	*generator.UnimplementedOrderServiceServer
}

func NewOrderServerService() *OrderClientService {
	return &OrderClientService{}
}
func (u *OrderClientService) CreateOrder(ctx context.Context, req *generator.CreateOrderRequest) (*generator.CreateOrderResponse, error) {
	return &generator.CreateOrderResponse{
		Order: &generator.Order{
			Id:    88888888888,
			Name:  req.Name,
			Price: req.Price + 99,
		},
	}, nil
}

func (u *OrderClientService) GetOrder(ctx context.Context, req *generator.GetOrderRequest) (*generator.GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (u *OrderClientService) ListOrders(ctx context.Context, req *generator.ListOrdersRequest) (*generator.ListOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
