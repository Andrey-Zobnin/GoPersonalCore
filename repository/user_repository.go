
package repository

import (
    "persona-core/model"
    "gorm.io/gorm"
)

type UserRepository interface {
    Create(user *model.User) (*model.User, error)
    GetByID(id string) (*model.User, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) (*model.User, error) {
    if err := r.db.Create(user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

func (r *userRepository) GetByID(id string) (*model.User, error) {
    var user model.User
    if err := r.db.First(&user, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
