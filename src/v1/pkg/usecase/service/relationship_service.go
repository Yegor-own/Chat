package service

import (
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/repository"
)

type RelationshipService interface {
	InsertRelationship(uId, cId uint) (*entities.Relationship, error)
	GetRelationshipByUserId(id uint) (*entities.Relationship, error)
	GetRelationshipByChatId(id uint) (*entities.Relationship, error)
	RemoveRelationship(relationshipId uint) error
}

type relationshipService struct {
	repository repository.RelationshipRepository
}

func NewRelationshipService(r repository.RelationshipRepository) RelationshipService {
	return &relationshipService{
		repository: r,
	}
}

func (s *relationshipService) InsertRelationship(uId, cId uint) (*entities.Relationship, error) {
	return s.repository.CreateRelationship(uId, cId)
}

func (s *relationshipService) GetRelationshipByUserId(id uint) (*entities.Relationship, error) {
	return s.repository.ReadByUserId(id)
}

func (s *relationshipService) GetRelationshipByChatId(id uint) (*entities.Relationship, error) {
	return s.repository.ReadByChatId(id)
}

func (s *relationshipService) RemoveRelationship(relationshipId uint) error {
	return s.repository.DeleteRelationship(relationshipId)
}
