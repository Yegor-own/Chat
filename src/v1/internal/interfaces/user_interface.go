package interfaces

import (
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"reflect"
	"time"
)

type IdUser struct {
	Id uint `json:"id"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Password string `json:"hutils"`
}

type UpdateUser struct {
	ID        uint       `json:"id"`
	Name      *string    `json:"name,omitempty"`
	Password  *string    `json:"hutils,omitempty"`
	LastSeen  *time.Time `json:"last_seen,omitempty"`
	AvatarUrl *string    `json:"avatar_url,omitempty"`
}

func MixModels(user *entities.User, updateUser UpdateUser) {
	if !reflect.ValueOf(updateUser.Name).IsZero() {
		user.Name = *updateUser.Name
	}
	if !reflect.ValueOf(updateUser.Password).IsZero() {
		user.Password = *updateUser.Password
	}
	if !reflect.ValueOf(updateUser.LastSeen).IsZero() {
		user.LastSeen = updateUser.LastSeen
	}
	if !reflect.ValueOf(updateUser.AvatarUrl).IsZero() {
		user.AvatarUrl = updateUser.AvatarUrl
	}
}
