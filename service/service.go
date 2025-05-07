// Package service provides the business logic service layer for the server
package service

import (
	"context"

	"github.com/project-inari/core-business-server/dto"
	"github.com/project-inari/core-business-server/repository"
)

// Port represents the service layer functions
type Port interface {
	CreateNewBusiness(ctx context.Context, req dto.CreateNewBusinessReq) (*dto.CreateNewBusinessRes, error)
	BusinessInquiry(ctx context.Context, businessID int) (*dto.BusinessInquiryRes, error)
	CreateNewCategory(ctx context.Context, req dto.CreateNewCategoryReq) (*dto.CreateNewCategoryRes, error)
	ListBusinessCategories(ctx context.Context, businessID int) (*dto.ListBusinessCategoriesRes, error)
	CreateNewTag(ctx context.Context, req dto.CreateNewTagReq) (*dto.CreateNewTagRes, error)
	ListBusinessTags(ctx context.Context, businessID int) (*dto.ListBusinessTagsRes, error)
	CreateNewWarehouse(ctx context.Context, req dto.CreateNewWarehouseReq) (*dto.CreateNewWarehouseRes, error)
	ListBusinessWarehouses(ctx context.Context, businessID int) (*dto.ListBusinessWarehousesRes, error)
	CreateNewSupplier(ctx context.Context, req dto.CreateNewSupplierReq) (*dto.CreateNewSupplierRes, error)
	CreateNewSupplierContact(ctx context.Context, req dto.CreateNewSupplierContactReq) (*dto.CreateNewSupplierContactRes, error)
	ListBusinessSuppliers(ctx context.Context, businessID int) (*dto.ListBusinessSuppliersRes, error)
	InquiryBusinessSupplier(ctx context.Context, supplierID int) (*dto.BusinessSupplierModel, error)
	CreateNewProduct(ctx context.Context, req dto.CreateNewProductReq) (*dto.CreateNewProductRes, error)
	ListBusinessProducts(ctx context.Context, businessID int) (*dto.ListBusinessProductsRes, error)
	InquiryBusinessProduct(ctx context.Context, productID int) (*dto.BusinessProductModel, error)
}

type service struct {
	databaseRepository repository.DatabaseRepository
	cacheRepository    repository.CacheRepository
}

// Dependencies represents the dependencies for the service
type Dependencies struct {
	DatabaseRepository repository.DatabaseRepository
	CacheRepository    repository.CacheRepository
}

// New creates a new service
func New(d Dependencies) Port {
	return &service{
		databaseRepository: d.DatabaseRepository,
		cacheRepository:    d.CacheRepository,
	}
}
