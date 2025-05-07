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
	CreateNewWarehouse(ctx context.Context, entity dto.BusinessWarehouseEntity) (int, error)
	ListBusinessWarehouses(ctx context.Context, businessID int) ([]dto.BusinessWarehouseEntity, error)
	CreateNewSupplier(ctx context.Context, entity dto.BusinessSupplierEntity) (int, error)
	CreateNewSupplierContact(ctx context.Context, entity dto.BusinessSupplierContactEntity) (int, error)
	ListBusinessSuppliers(ctx context.Context, businessID int) ([]dto.BusinessSupplierEntity, error)
	ListBusinessSupplierContacts(ctx context.Context, supplierID int) ([]dto.BusinessSupplierContactEntity, error)
	InquiryBusinessSupplier(ctx context.Context, supplierID int) (*dto.BusinessSupplierEntity, error)
	CreateNewProduct(ctx context.Context, product dto.BusinessProductEntity, variants []dto.BusinessProductVariantEntity) (int, error)
	ListBusinessProducts(ctx context.Context, businessID int) ([]dto.BusinessProductEntity, error)
	ListBusinessProductVariants(ctx context.Context, productID int) ([]dto.BusinessProductVariantEntity, error)
	InquiryBusinessProduct(ctx context.Context, productID int) (*dto.BusinessProductEntity, error)
	InquiryBusinessProductVariant(ctx context.Context, variantID int) (*dto.BusinessProductVariantEntity, error)
	ListProductVariantsInWarehouse(ctx context.Context, variantID int) ([]dto.WarehouseQty, error)
	ListProductVariantTags(ctx context.Context, variantID int) ([]dto.BusinessProductVariantTagEntity, error)
	CreateNewSupplierOrder(ctx context.Context, orderInfo dto.CreateNewSupplierOrderReq) (int, error)
	ListSupplierOrders(ctx context.Context, businessID int) ([]dto.SupplierOrderModel, error)
	ListBusinessInventory(ctx context.Context, businessID int) ([]dto.BusinessInventoryModel, error)
	InquiryProductInventory(ctx context.Context, variantID int) (*dto.BusinessInventoryModel, error)
	CreateNewCustomer(ctx context.Context, entity dto.BusinessCustomerEntity) (int, error)
	ListBusinessCustomers(ctx context.Context, businessID int) ([]dto.BusinessCustomerEntity, error)
	CreateNewCustomerOrder(ctx context.Context, orderInfo dto.CreateNewCustomerOrderReq) (int, error)
	ListCustomerOrders(ctx context.Context, businessID int) ([]dto.CustomerOrderModel, error)
}

// CacheRepository represents the repository layer functions of cache repository
type CacheRepository interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) *redis.StatusCmd
	UpdateUserCacheNewBusinessJoined(ctx context.Context, ownerUsername string, business dto.BusinessCacheModel) error
}
