// Package repository provides the repository interfaces for the domain
package repository

import (
	"context"
	"time"

	"github.com/project-inari/core-business-server/dto"

	"github.com/redis/go-redis/v9"
)

// DatabaseRepository represents the repository layer functions of database repository
type DatabaseRepository interface {
	CreateNewBusiness(ctx context.Context, username string, entity dto.BusinessEntity) (*dto.BusinessEntity, error)
	GetBusiness(ctx context.Context, businessID int) (*dto.BusinessEntity, error)
	CreateNewCategory(ctx context.Context, entity dto.BusinessCategoryEntity) (int, error)
	ListBusinessCategories(ctx context.Context, businessID int) ([]dto.BusinessCategoryEntity, error)
	CreateNewTag(ctx context.Context, entity dto.BusinessTagEntity) (int, error)
	ListBusinessTags(ctx context.Context, businessID int) ([]dto.BusinessTagEntity, error)
	AddTagsToCategory(ctx context.Context, tagIds []int, categoryID int) error
	ListCategoryTags(ctx context.Context, categoryID int) ([]dto.BusinessTagEntity, error)
}

// CacheRepository represents the repository layer functions of cache repository
type CacheRepository interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) *redis.StatusCmd
	UpdateUserCacheNewBusinessJoined(ctx context.Context, ownerUsername string, business dto.BusinessCacheModel) error
}
