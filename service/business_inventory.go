package service

import (
	"context"

	"github.com/project-inari/core-business-server/dto"
)

func (s *service) ListBusinessInventory(ctx context.Context, businessID int) (*dto.ListBusinessInventoryRes, error) {
	queryRes, err := s.databaseRepository.ListBusinessInventory(ctx, businessID)
	if err != nil {
		return nil, err
	}

	return &dto.ListBusinessInventoryRes{
		Inventory: queryRes,
	}, nil
}

func (s *service) InquiryProductInventory(ctx context.Context, variantID int) (*dto.BusinessInventoryModel, error) {
	queryRes, err := s.databaseRepository.InquiryProductInventory(ctx, variantID)
	if err != nil {
		return nil, err
	}

	return queryRes, nil
}
