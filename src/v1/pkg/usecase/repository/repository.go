package repository

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}

type chatRepository struct {
	db *gorm.DB
}

type relationshipRepository struct {
	db *gorm.DB
}

type messagesRepository struct {
	db *gorm.DB
}
