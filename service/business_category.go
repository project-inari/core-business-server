package service

import (
	"context"
	"github.com/project-inari/core-business-server/dto"
)

func (s *service) CreateNewCategory(ctx context.Context, req dto.CreateNewCategoryReq) (*dto.CreateNewCategoryRes, error) {
	parentCategoryID := new(int)
	if req.ParentCategoryID == 0 {
		parentCategoryID = nil
	} else {
		*parentCategoryID = req.ParentCategoryID
	}

	queryRes, err := s.databaseRepository.CreateNewCategory(ctx, dto.BusinessCategoryEntity{
		BusinessID:         req.BusinessID,
		CategoryName:       req.CategoryName,
		CategoryPictureURL: req.CategoryPictureURL,
		Description:        req.Description,
		ParentCategoryID:   parentCategoryID,
	})
	if err != nil {
		return nil, err
	}

	if err := s.databaseRepository.AddTagsToCategory(ctx, req.TagIDs, queryRes); err != nil {
		return nil, err
	}

	return &dto.CreateNewCategoryRes{
		CategoryID: queryRes,
		Success:    true,
	}, nil
}

func (s *service) ListBusinessCategories(ctx context.Context, businessID int) (*dto.ListBusinessCategoriesRes, error) {
	queryRes, err := s.databaseRepository.ListBusinessCategories(ctx, businessID)
	if err != nil {
		return nil, err
	}

	return s.constructListBusinessCategoriesRes(ctx, queryRes)
}

func (s *service) constructListBusinessCategoriesRes(ctx context.Context, entity []dto.BusinessCategoryEntity) (*dto.ListBusinessCategoriesRes, error) {
	var data []dto.BusinessCategoryModel
	for _, category := range entity {
		categoryTags, err := s.databaseRepository.ListCategoryTags(ctx, category.ID)
		if err != nil {
			return nil, err
		}

		var tags []dto.BusinessTagModel
		for _, tag := range categoryTags {
			tags = append(tags, dto.BusinessTagModel{
				ID:      tag.ID,
				TagName: tag.TagName,
				Color:   tag.Color,
			})
		}

		data = append(data, dto.BusinessCategoryModel{
			ID:                 category.ID,
			CategoryName:       category.CategoryName,
			CategoryPictureURL: category.CategoryPictureURL,
			Description:        category.Description,
			ParentCategoryID:   category.ParentCategoryID,
			Tags:               tags,
			CreatedAt:          category.CreatedAt,
			UpdatedAt:          category.UpdatedAt,
		})
	}

	res := dto.ListBusinessCategoriesRes{
		BusinessCategories: data,
	}
	return &res, nil
}
