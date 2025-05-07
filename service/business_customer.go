package service

import (
	"context"

	"github.com/project-inari/core-business-server/dto"
)

func (s *service) CreateNewCustomer(ctx context.Context, req dto.CreateNewCustomerReq) (*dto.CreateNewCustomerRes, error) {
	customerEntity := dto.BusinessCustomerEntity{
		BusinessID: req.BusinessID,
		Name:       req.Name,
		Type:       req.Type,
		Address:    req.Address,
		PhoneNo:    req.PhoneNo,
	}

	customerID, err := s.databaseRepository.CreateNewCustomer(ctx, customerEntity)
	if err != nil {
		return nil, err
	}

	return &dto.CreateNewCustomerRes{
		CustomerID: customerID,
		Success:    true,
	}, nil
}

func (s *service) ListBusinessCustomers(ctx context.Context, businessID int) (*dto.ListBusinessCustomersRes, error) {
	queryRes, err := s.databaseRepository.ListBusinessCustomers(ctx, businessID)
	if err != nil {
		return nil, err
	}

	var customers []dto.BusinessCustomerModel
	for _, customer := range queryRes {
		customers = append(customers, dto.BusinessCustomerModel{
			ID:        customer.ID,
			Name:      customer.Name,
			Type:      customer.Type,
			Address:   customer.Address,
			PhoneNo:   customer.PhoneNo,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
		})
	}

	return &dto.ListBusinessCustomersRes{
		Customers: customers,
	}, nil
}
