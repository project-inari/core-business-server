package service

import (
	"context"

	"github.com/project-inari/core-business-server/dto"
	"github.com/project-inari/core-business-server/pkg/utils"
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

	return &dto.CreateNewBusinessRes{
		BusinessID:   businessRes.ID,
		BusinessName: businessRes.Name,
		Success:      true,
	}, nil
}
