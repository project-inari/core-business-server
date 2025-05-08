package service

import (
	"context"

	"github.com/project-inari/core-business-server/dto"
)

func (s *service) CreateNewTag(ctx context.Context, req dto.CreateNewTagReq) (*dto.CreateNewTagRes, error) {
	tagEntity := dto.BusinessTagEntity{
		BusinessID:  req.BusinessID,
		TagName:     req.TagName,
		Color:       req.Color,
		Description: req.Description,
	}

	tagID, err := s.databaseRepository.CreateNewTag(ctx, tagEntity)
	if err != nil {
		return nil, err
	}

	return &dto.CreateNewTagRes{
		TagID:   tagID,
		Success: true,
	}, nil
}

func (s *service) ListBusinessTags(ctx context.Context, businessID int) (*dto.ListBusinessTagsRes, error) {
	queryRes, err := s.databaseRepository.ListBusinessTags(ctx, businessID)
	if err != nil {
		return nil, err
	}

	var tags []dto.BusinessTagModel
	for _, tag := range queryRes {
		tags = append(tags, dto.BusinessTagModel{
			ID:          tag.ID,
			TagName:     tag.TagName,
			Color:       tag.Color,
			Description: tag.Description,
			CreatedAt:   tag.CreatedAt,
			UpdatedAt:   tag.UpdatedAt,
		})
	}

	return &dto.ListBusinessTagsRes{
		BusinessTags: tags,
	}, nil
}
