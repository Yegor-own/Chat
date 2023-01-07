package repository

import (
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"gorm.io/gorm"
)

type RelationshipRepository interface {
	CreateRelationship(uId, cId uint) (*entities.Relationship, error)
	ReadByUserId(id uint) (*entities.Relationship, error)
	ReadByChatId(id uint) (*entities.Relationship, error)
	DeleteRelationship(relationshipId uint) error
}

func NewRelationshipRepositorySQL(db *gorm.DB) RelationshipRepository {
	return &relationshipRepository{
		db: db,
	}
}

func (r *relationshipRepository) CreateRelationship(uId, cId uint) (*entities.Relationship, error) {
	rel := entities.Relationship{
		UserID: uId,
		ChatID: cId,
	}

	res := r.db.Create(&rel)
	if res.Error != nil {
		return nil, res.Error
	}
	return &rel, nil
}

func (r *relationshipRepository) ReadByUserId(id uint) (*entities.Relationship, error) {
	rel := entities.Relationship{}
	res := r.db.First(&rel, "UserID = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &rel, nil
}

func (r *relationshipRepository) ReadByChatId(id uint) (*entities.Relationship, error) {
	rel := entities.Relationship{}
	res := r.db.First(&rel, "ChatID = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &rel, nil
}

func (r *relationshipRepository) DeleteRelationship(relationshipId uint) error {
	rel := entities.Relationship{ID: relationshipId}
	res := r.db.Delete(&rel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
