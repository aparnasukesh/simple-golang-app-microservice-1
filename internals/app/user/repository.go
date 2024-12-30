package user

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	CreateUser(ctx context.Context, user User) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserDetailsById(ctx context.Context, id int) (*User, error)
	UpdateUserProfile(ctx context.Context, updateUser UserProfileDetails, id int) error
	DeleteUserProfile(ctx context.Context, id int64) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(ctx context.Context, user User) (int64, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return int64(user.ID), nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	dbData := &User{}
	result := r.db.Where("email = ?", email).First(&dbData)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return dbData, nil
}

func (r *repository) GetUserDetailsById(ctx context.Context, id int) (*User, error) {
	user := &User{}
	result := r.db.Where("id =?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (r *repository) UpdateUserProfile(ctx context.Context, updateUser UserProfileDetails, id int) error {
	r.db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)

	result := r.db.Model(&User{}).Where("id = ?", id).Updates(updateUser)
	if result.Error != nil {
		return result.Error
	}
	updatedUser := User{}
	if err := r.db.Where("id=?", id).First(&updatedUser).Error; err != nil {
		return err
	}
	return nil
}
func (r *repository) DeleteUserProfile(ctx context.Context, id int64) error {
	err := r.db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete user profile with id %d: %w", id, err)
	}
	return nil
}