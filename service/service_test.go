package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"

	"github.com/project-inari/core-business-server/dto"
)

type mockDatabaseRepository struct {
	createNewBusinessRes *dto.BusinessEntity
	getBusinessRes       *dto.BusinessEntity
	err                  error
}

func (m *mockDatabaseRepository) CreateNewBusiness(_ context.Context, _ string, _ dto.BusinessEntity) (*dto.BusinessEntity, error) {
	return m.createNewBusinessRes, m.err
}

func (m *mockDatabaseRepository) GetBusiness(_ context.Context, _ string) (*dto.BusinessEntity, error) {
	return m.getBusinessRes, m.err
}

type mockCacheRepository struct {
	getRes *redis.StringCmd
	setRes *redis.StatusCmd
	err    error
}

func (m *mockCacheRepository) Get(_ context.Context, _ string) *redis.StringCmd {
	return m.getRes
}

func (m *mockCacheRepository) Set(_ context.Context, _ string, _ interface{}, _ time.Duration) *redis.StatusCmd {
	return m.setRes
}

func (m *mockCacheRepository) UpdateUserCacheNewBusinessJoined(_ context.Context, _ string, _ dto.BusinessInquiryRes) error {
	return m.err
}

const (
	mockName             = "mockName"
	mockIndustryType     = "mockIndustryType"
	mockBusinessType     = "mockBusinessType"
	mockDescription      = "mockDescription"
	mockPhoneNo          = "mockPhoneNo"
	mockOperatingHours   = `{"monday":{"open":true,"openTime":"09:00","closeTime":"17:00"},"tuesday":{"open":true,"openTime":"09:00","closeTime":"17:00"},"wednesday":{"open":true,"openTime":"09:00","closeTime":"17:00"},"thursday":{"open":true,"openTime":"09:00","closeTime":"17:00"},"friday":{"open":true,"openTime":"09:00","closeTime":"17:00"},"saturday":{"open":false},"sunday":{"open":false}}`
	mockAddress          = "mockAddress"
	mockBusinessImageURL = "mockBusinessImageURL"
	mockOwnerUsername    = "mockOwnerUsername"
)

var (
	mockOperatingHoursStruct = dto.OperatingHours{
		Monday: dto.OpenTime{
			Open:      true,
			OpenTime:  "09:00",
			CloseTime: "17:00",
		},
		Tuesday: dto.OpenTime{
			Open:      true,
			OpenTime:  "09:00",
			CloseTime: "17:00",
		},
		Wednesday: dto.OpenTime{
			Open:      true,
			OpenTime:  "09:00",
			CloseTime: "17:00",
		},
		Thursday: dto.OpenTime{
			Open:      true,
			OpenTime:  "09:00",
			CloseTime: "17:00",
		},
		Friday: dto.OpenTime{
			Open:      true,
			OpenTime:  "09:00",
			CloseTime: "17:00",
		},
		Saturday: dto.OpenTime{
			Open: false,
		},
		Sunday: dto.OpenTime{
			Open: false,
		},
	}
)

func TestCreateNewBusiness(t *testing.T) {
	ctx := context.Background()
	req := dto.CreateNewBusinessReq{
		Name:             mockName,
		IndustryType:     mockIndustryType,
		BusinessType:     mockBusinessType,
		Description:      mockDescription,
		PhoneNo:          mockPhoneNo,
		OperatingHours:   mockOperatingHoursStruct,
		Address:          mockAddress,
		BusinessImageURL: mockBusinessImageURL,
		OwnerUsername:    mockOwnerUsername,
	}

	expectedRes := &dto.CreateNewBusinessRes{
		BusinessID:   1,
		BusinessName: mockName,
		Success:      true,
	}

	t.Run("success", func(t *testing.T) {
		mockCacheRepository := &mockCacheRepository{
			err: nil,
		}
		mockDatabaseRepository := &mockDatabaseRepository{
			createNewBusinessRes: &dto.BusinessEntity{
				ID:               1,
				Name:             mockName,
				IndustryType:     mockIndustryType,
				BusinessType:     mockBusinessType,
				Description:      mockDescription,
				PhoneNo:          mockPhoneNo,
				OperatingHours:   mockOperatingHours,
				Address:          mockAddress,
				BusinessImageURL: mockBusinessImageURL,
			},
		}

		s := &service{
			databaseRepository: mockDatabaseRepository,
			cacheRepository:    mockCacheRepository,
		}

		res, err := s.CreateNewBusiness(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})

	t.Run("error - when database repo returned error", func(t *testing.T) {
		mockCacheRepository := &mockCacheRepository{}
		mockDatabaseRepository := &mockDatabaseRepository{
			err: errors.New("error"),
		}

		s := &service{
			databaseRepository: mockDatabaseRepository,
			cacheRepository:    mockCacheRepository,
		}

		res, err := s.CreateNewBusiness(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("error - when cache repo returned error", func(t *testing.T) {
		mockCacheRepository := &mockCacheRepository{
			err: errors.New("error"),
		}
		mockDatabaseRepository := &mockDatabaseRepository{
			createNewBusinessRes: &dto.BusinessEntity{
				ID:               1,
				Name:             mockName,
				IndustryType:     mockIndustryType,
				BusinessType:     mockBusinessType,
				Description:      mockDescription,
				PhoneNo:          mockPhoneNo,
				OperatingHours:   mockOperatingHours,
				Address:          mockAddress,
				BusinessImageURL: mockBusinessImageURL,
			},
		}

		s := &service{
			databaseRepository: mockDatabaseRepository,
			cacheRepository:    mockCacheRepository,
		}

		res, err := s.CreateNewBusiness(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestBusinessInquiry(t *testing.T) {
	ctx := context.Background()

	expectedRes := &dto.BusinessInquiryRes{
		ID:               1,
		Name:             mockName,
		IndustryType:     mockIndustryType,
		BusinessType:     mockBusinessType,
		Description:      mockDescription,
		PhoneNo:          mockPhoneNo,
		OperatingHours:   mockOperatingHoursStruct,
		Address:          mockAddress,
		BusinessImageURL: mockBusinessImageURL,
		CreatedAt:        "",
		UpdatedAt:        "",
	}

	t.Run("success", func(t *testing.T) {
		mockCacheRepository := &mockCacheRepository{}
		mockDatabaseRepository := &mockDatabaseRepository{
			getBusinessRes: &dto.BusinessEntity{
				ID:               1,
				Name:             mockName,
				IndustryType:     mockIndustryType,
				BusinessType:     mockBusinessType,
				Description:      mockDescription,
				PhoneNo:          mockPhoneNo,
				OperatingHours:   mockOperatingHours,
				Address:          mockAddress,
				BusinessImageURL: mockBusinessImageURL,
			},
		}

		s := &service{
			databaseRepository: mockDatabaseRepository,
			cacheRepository:    mockCacheRepository,
		}

		res, err := s.BusinessInquiry(ctx, mockName)

		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})

	t.Run("error - when database repo returned error", func(t *testing.T) {
		mockCacheRepository := &mockCacheRepository{}
		mockDatabaseRepository := &mockDatabaseRepository{
			err: errors.New("error"),
		}

		s := &service{
			databaseRepository: mockDatabaseRepository,
			cacheRepository:    mockCacheRepository,
		}

		res, err := s.BusinessInquiry(ctx, mockName)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
