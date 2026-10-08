package service

import (
	"context"
	"theater-backend/internal/models"
	"theater-backend/internal/repository"
)

type PlayService struct {
	repo *repository.PlayRepository
}

func NewPlayService(repo *repository.PlayRepository) *PlayService {
	return &PlayService{repo: repo}
}

func (s *PlayService) GetAllPlays(ctx context.Context) ([]models.Play, error) {
	return s.repo.GetAll(ctx)
}

func (s *PlayService) CreatePlay(ctx context.Context, play *models.Play) error {
	return s.repo.Create(ctx, play)
}
