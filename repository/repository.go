// Package repository provides the repository interfaces for the domain
package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/project-inari/core-business-server/dto"
)

// DatabaseRepository represents the repository layer functions of database repository
type DatabaseRepository interface {
	QueryTest() (*[]dto.TestEntity, error)
}

// CacheRepository represents the repository layer functions of cache repository
type CacheRepository interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) *redis.StatusCmd
}
