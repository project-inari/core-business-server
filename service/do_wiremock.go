package service

import (
	"context"

	"github.com/project-inari/core-business-server/dto"
)

func (s *service) DoWiremock(ctx context.Context) (*dto.WiremockGetTestResponse, error) {
	res, err := s.wiremockAPIRepository.GetTest(ctx, dto.WiremockGetTestHeader{
		ContentType: "application/json",
		RequestID:   "123",
	})
	if err != nil {
		return nil, err
	}

	return &res.Response, nil
}
