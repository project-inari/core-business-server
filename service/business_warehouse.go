package service

import (
	"context"

	"github.com/project-inari/core-business-server/dto"
)

func (s *service) CreateNewWarehouse(ctx context.Context, req dto.CreateNewWarehouseReq) (*dto.CreateNewWarehouseRes, error) {
	warehouseEntity := dto.BusinessWarehouseEntity{
		BusinessID:          req.BusinessID,
		Name:                req.WarehouseName,
		Description:         req.Description,
		WarehousePictureURL: req.WarehousePictureURL,
	}

	warehouseID, err := s.databaseRepository.CreateNewWarehouse(ctx, warehouseEntity)
	if err != nil {
		return nil, err
	}

	return &dto.CreateNewWarehouseRes{
		WarehouseID: warehouseID,
		Success:     true,
	}, nil
}

func (s *service) ListBusinessWarehouses(ctx context.Context, businessID int) (*dto.ListBusinessWarehousesRes, error) {
	queryRes, err := s.databaseRepository.ListBusinessWarehouses(ctx, businessID)
	if err != nil {
		return nil, err
	}

	var warehouses []dto.BusinessWarehouseModel
	for _, warehouse := range queryRes {
		warehouses = append(warehouses, dto.BusinessWarehouseModel{
			ID:                  warehouse.ID,
			Name:                warehouse.Name,
			Description:         warehouse.Description,
			WarehousePictureURL: warehouse.WarehousePictureURL,
			CreatedAt:           warehouse.CreatedAt,
			UpdatedAt:           warehouse.UpdatedAt,
		})
	}

	return &dto.ListBusinessWarehousesRes{
		Warehouses: warehouses,
	}, nil
}
