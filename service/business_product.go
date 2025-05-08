package service

import (
	"context"

	"github.com/project-inari/core-business-server/dto"
)

func (s *service) CreateNewProduct(ctx context.Context, req dto.CreateNewProductReq) (*dto.CreateNewProductRes, error) {
	productEntity := dto.BusinessProductEntity{
		BusinessID: req.BusinessID,
		SupplierID: req.SupplierID,
		ItemName:   req.Name,
		Brand:      req.Brand,
		CategoryID: req.CategoryID,
	}

	var variantEntities []dto.BusinessProductVariantEntity
	for _, variant := range req.Variants {
		variantEntity := dto.BusinessProductVariantEntity{
			VariantName:       variant.VariantName,
			SKUNo:             variant.SKUNo,
			PictureURL:        variant.PictureURL,
			BaseSellingPrice:  variant.BaseSellingPrice,
			BasePurchasePrice: variant.BasePurchasePrice,
			Note:              variant.Note,
			TagIDs:            &variant.TagIds,
		}
		variantEntities = append(variantEntities, variantEntity)
	}

	productID, err := s.databaseRepository.CreateNewProduct(ctx, productEntity, variantEntities)
	if err != nil {
		return nil, err
	}

	return &dto.CreateNewProductRes{
		ProductID: productID,
	}, nil
}

func (s *service) ListBusinessProducts(ctx context.Context, businessID int) (*dto.ListBusinessProductsRes, error) {
	queryRes, err := s.databaseRepository.ListBusinessProducts(ctx, businessID)
	if err != nil {
		return nil, err
	}

	var products []dto.BusinessProductModel
	for _, product := range queryRes {
		variants, err := s.databaseRepository.ListBusinessProductVariants(ctx, product.ID)
		if err != nil {
			return nil, err
		}

		var variantModels []dto.BusinessVariantModel
		for _, variant := range variants {
			qtyInWarehouse, err := s.databaseRepository.ListProductVariantsInWarehouse(ctx, variant.ID)
			if err != nil {
				return nil, err
			}

			variantTagIds, err := s.databaseRepository.ListProductVariantTags(ctx, variant.ID)
			if err != nil {
				return nil, err
			}

			var tagIds []int
			for _, tag := range variantTagIds {
				tagIds = append(tagIds, tag.TagID)
			}

			variantModel := dto.BusinessVariantModel{
				ID:                variant.ID,
				ProductID:         variant.ProductID,
				VariantName:       variant.VariantName,
				SKUNo:             variant.SKUNo,
				PictureURL:        variant.PictureURL,
				BaseSellingPrice:  variant.BaseSellingPrice,
				BasePurchasePrice: variant.BasePurchasePrice,
				Note:              variant.Note,
				TagIds:            tagIds,
				QtyInWarehouse:    qtyInWarehouse,
				CreatedAt:         variant.CreatedAt,
				UpdatedAt:         variant.UpdatedAt,
			}
			variantModels = append(variantModels, variantModel)
		}

		productModel := dto.BusinessProductModel{
			ID:         product.ID,
			BusinessID: product.BusinessID,
			SupplierID: product.SupplierID,
			Name:       product.ItemName,
			Brand:      product.Brand,
			CategoryID: product.CategoryID,
			Variants:   variantModels,
			CreatedAt:  product.CreatedAt,
			UpdatedAt:  product.UpdatedAt,
		}
		products = append(products, productModel)
	}

	return &dto.ListBusinessProductsRes{
		Products: products,
	}, nil
}

func (s *service) InquiryBusinessProduct(ctx context.Context, productID int) (*dto.BusinessProductModel, error) {
	product, err := s.databaseRepository.InquiryBusinessProduct(ctx, productID)
	if err != nil {
		return nil, err
	}

	variants, err := s.databaseRepository.ListBusinessProductVariants(ctx, product.ID)
	if err != nil {
		return nil, err
	}

	var variantModels []dto.BusinessVariantModel
	for _, variant := range variants {
		qtyInWarehouse, err := s.databaseRepository.ListProductVariantsInWarehouse(ctx, variant.ID)
		if err != nil {
			return nil, err
		}

		variantTagIds, err := s.databaseRepository.ListProductVariantTags(ctx, variant.ID)
		if err != nil {
			return nil, err
		}

		var tagIds []int
		for _, tag := range variantTagIds {
			tagIds = append(tagIds, tag.TagID)
		}

		variantModel := dto.BusinessVariantModel{
			ID:                variant.ID,
			ProductID:         variant.ProductID,
			VariantName:       variant.VariantName,
			SKUNo:             variant.SKUNo,
			PictureURL:        variant.PictureURL,
			BaseSellingPrice:  variant.BaseSellingPrice,
			BasePurchasePrice: variant.BasePurchasePrice,
			Note:              variant.Note,
			TagIds:            tagIds,
			QtyInWarehouse:    qtyInWarehouse,
			CreatedAt:         variant.CreatedAt,
			UpdatedAt:         variant.UpdatedAt,
		}
		variantModels = append(variantModels, variantModel)
	}

	return &dto.BusinessProductModel{
		ID:         product.ID,
		Name:       product.ItemName,
		SupplierID: product.SupplierID,
		Brand:      product.Brand,
		CategoryID: product.CategoryID,
		Variants:   variantModels,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}, nil
}

func (s *service) InquiryBusinessProductVariant(ctx context.Context, variantID int) (*dto.BusinessVariantModel, error) {
	variant, err := s.databaseRepository.InquiryBusinessProductVariant(ctx, variantID)
	if err != nil {
		return nil, err
	}

	qtyInWarehouse, err := s.databaseRepository.ListProductVariantsInWarehouse(ctx, variant.ID)
	if err != nil {
		return nil, err
	}

	variantTagIds, err := s.databaseRepository.ListProductVariantTags(ctx, variant.ID)
	if err != nil {
		return nil, err
	}

	var tagIds []int
	for _, tag := range variantTagIds {
		tagIds = append(tagIds, tag.TagID)
	}

	return &dto.BusinessVariantModel{
		ID:                variant.ID,
		ProductID:         variant.ProductID,
		VariantName:       variant.VariantName,
		SKUNo:             variant.SKUNo,
		PictureURL:        variant.PictureURL,
		BaseSellingPrice:  variant.BaseSellingPrice,
		BasePurchasePrice: variant.BasePurchasePrice,
		Note:              variant.Note,
		TagIds:            tagIds,
		QtyInWarehouse:    qtyInWarehouse,
	}, nil
}
