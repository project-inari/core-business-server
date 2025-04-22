package service

import (
	"context"

	"github.com/project-inari/core-business-server/dto"
	"github.com/project-inari/core-business-server/pkg/utils"
)

func (s *service) BusinessInquiry(ctx context.Context, businessID int) (*dto.BusinessInquiryRes, error) {
	queryRes, err := s.databaseRepository.GetBusiness(ctx, businessID)
	if err != nil {
		return nil, err
	}

	return constructBusinessInquiryRes(*queryRes), nil
}

func constructBusinessInquiryRes(entity dto.BusinessEntity) *dto.BusinessInquiryRes {
	operatingHours := utils.DecodeJSONfromString[dto.OperatingHours](entity.OperatingHours)

	return &dto.BusinessInquiryRes{
		ID:               entity.ID,
		Name:             entity.Name,
		IndustryType:     entity.IndustryType,
		BusinessType:     entity.BusinessType,
		Description:      entity.Description,
		PhoneNo:          entity.PhoneNo,
		OperatingHours:   *operatingHours,
		Address:          entity.Address,
		BusinessImageURL: entity.BusinessImageURL,
		CreatedAt:        entity.CreatedAt,
		UpdatedAt:        entity.UpdatedAt,
	}
}
