package repository

import (
	"context"
	"log"
	"todo-golang/domain/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	log.Printf("[FindByEmail] called with email: %s", email)

	var user entity.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	log.Printf("found user: %+v", user)

	return &user, nil
}

// func (r *UserRepository) Update(user entity.User) (entity.User, error) {
// 	if err := r.db.Save(&user).Error; err != nil {
// 		return entity.User{}, err
// 	}
// 	return user, nil
// }

// func (r *UserRepository) Delete(user entity.User) error {
// 	if err := r.db.Delete(&entity.User{}, user).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
