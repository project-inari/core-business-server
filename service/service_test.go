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
	err                  error
}

func (m *mockDatabaseRepository) CreateNewBusiness(_ context.Context, _ string, _ dto.BusinessEntity) (*dto.BusinessEntity, error) {
	return m.createNewBusinessRes, m.err
}

type mockCacheRepository struct{}

func (m *mockCacheRepository) Get(_ context.Context, _ string) *redis.StringCmd {
	return nil
}

func (m *mockCacheRepository) Set(_ context.Context, _ string, _ interface{}, _ time.Duration) *redis.StatusCmd {
	return nil
}

const (
	mockName             = "mockName"
	mockIndustryType     = "mockIndustryType"
	mockBusinessType     = "mockBusinessType"
	mockDescription      = "mockDescription"
	mockPhoneNo          = "mockPhoneNo"
	mockOperatingHours   = "mockOperatingHours"
	mockAddress          = "mockAddress"
	mockBusinessImageURL = "mockBusinessImageURL"
	mockOwnerUsername    = "mockOwnerUsername"
)

func TestCreateNewBusiness(t *testing.T) {
	ctx := context.Background()
	req := dto.CreateNewBusinessReq{
		Name:         mockName,
		IndustryType: mockIndustryType,
		BusinessType: mockBusinessType,
		Description:  mockDescription,
		PhoneNo:      mockPhoneNo,
		OperatingHours: dto.OperatingHours{
			Monday: dto.OpenTime{
				Open:      true,
				OpenTime:  "08:00",
				CloseTime: "17:00",
			},
			Tuesday: dto.OpenTime{
				Open:      true,
				OpenTime:  "08:00",
				CloseTime: "17:00",
			},
			Wednesday: dto.OpenTime{
				Open:      true,
				OpenTime:  "08:00",
				CloseTime: "17:00",
			},
			Thursday: dto.OpenTime{
				Open:      true,
				OpenTime:  "08:00",
				CloseTime: "17:00",
			},
			Friday: dto.OpenTime{
				Open:      true,
				OpenTime:  "08:00",
				CloseTime: "17:00",
			},
			Saturday: dto.OpenTime{
				Open: false,
			},
			Sunday: dto.OpenTime{
				Open: false,
			},
		},
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
		mockCacheRepository := &mockCacheRepository{}
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
}
