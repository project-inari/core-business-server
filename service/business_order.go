package service

import (
	"context"

	"github.com/project-inari/core-business-server/dto"
)

func (s *service) CreateNewCustomerOrder(ctx context.Context, req dto.CreateNewCustomerOrderReq) (*dto.CreateNewCustomerOrderRes, error) {
	customerOrderID, err := s.databaseRepository.CreateNewCustomerOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	return &dto.CreateNewCustomerOrderRes{
		CustomerOrderID: customerOrderID,
		Success:         true,
	}, nil
}

func (s *service) ListCustomerOrders(ctx context.Context, businessID int) (*dto.ListCustomerOrdersRes, error) {
	queryRes, err := s.databaseRepository.ListCustomerOrders(ctx, businessID)
	if err != nil {
		return nil, err
	}

	return &dto.ListCustomerOrdersRes{
		CustomerOrders: queryRes,
	}, nil
}
