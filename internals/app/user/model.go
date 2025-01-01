package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"not null" json:"username" validate:"required,min=8,max=24"`
	Password    string `gorm:"not null" json:"password" validate:"required,min=6,max=12"`
	Email       string `gorm:"not null;unique" json:"email" validate:"email,required"`
	PhoneNumber string `gorm:"not null" json:"phone" validate:"required,len=10"`
	FirstName   string `gorm:"not null" json:"firstname" validate:"required"`
	LastName    string `gorm:"not null" json:"lastname" validate:"required"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
}

type UserProfileDetails struct {
	Username    string `json:"username" validate:"required,min=8,max=24"`
	FirstName   string `gorm:"not null" json:"firstname" validate:"required,min=4,max=10"`
	LastName    string `gorm:"not null" json:"lastname" validate:"required,min=4,max=10"`
	PhoneNumber string `json:"phone" validate:"required,len=10"`
	Email       string `json:"email" validate:"email,required"`
	DateOfBirth string `json:"dateofbirth"`
	Gender      string `json:"gender"`
}

type ListUserRequest struct {
	Method   int32 `json:"method"`
	WaitTime int32 `json:"wait_time"`
}
