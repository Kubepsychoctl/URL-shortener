package user

import (
	"app/url-shorter/pkg/db"
)

type UserRepository struct {
	database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{database: database}
}

func (r *UserRepository) Create(user *User) (*User, error) {
	result := r.database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := r.database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
