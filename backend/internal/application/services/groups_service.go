package services

import (
	"context"
	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/repository"
	"social_network/internal/domain/ports/service"
)

type groupService struct {
	repo repository.GroupRepository
}

func NewGroupService(r repository.GroupRepository) service.GroupService {
	return &groupService{repo: r}
}

func (s *groupService) CreateGroup(ctx context.Context, g *models.Group) error {
	return s.repo.CreateGroup(ctx, g)
}
