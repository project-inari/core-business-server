package service

import (
	"context"

	"github.com/project-inari/core-business-server/dto"
	"github.com/project-inari/core-business-server/pkg/utils"
)

const (
	roleOwner = "OWNER"
)

func (s *service) CreateNewBusiness(ctx context.Context, req dto.CreateNewBusinessReq) (*dto.CreateNewBusinessRes, error) {
	businessEntity := dto.BusinessEntity{
		Name:             req.Name,
		IndustryType:     req.IndustryType,
		BusinessType:     req.BusinessType,
		Description:      req.Description,
		PhoneNo:          req.PhoneNo,
		OperatingHours:   utils.EncodeJSONtoString(req.OperatingHours),
		Address:          req.Address,
		BusinessImageURL: req.BusinessImageURL,
	}

	businessRes, err := s.databaseRepository.CreateNewBusiness(ctx, req.OwnerUsername, businessEntity)
	if err != nil {
		return nil, err
	}

	if err := s.cacheRepository.UpdateUserCacheNewBusinessJoined(ctx, req.OwnerUsername, dto.BusinessCacheModel{
		ID:               businessRes.ID,
		Name:             businessRes.Name,
		IndustryType:     businessRes.IndustryType,
		BusinessType:     businessRes.BusinessType,
		Description:      businessRes.Description,
		PhoneNo:          businessRes.PhoneNo,
		OperatingHours:   *utils.DecodeJSONfromString[dto.OperatingHours](businessRes.OperatingHours),
		Address:          businessRes.Address,
		BusinessImageURL: businessRes.BusinessImageURL,
		CreatedAt:        businessRes.CreatedAt,
		UpdatedAt:        businessRes.UpdatedAt,
		UserRole:         roleOwner,
	}); err != nil {
		return nil, err
	}

	return &dto.CreateNewBusinessRes{
		BusinessID:   businessRes.ID,
		BusinessName: businessRes.Name,
		Success:      true,
	}, nil
}
