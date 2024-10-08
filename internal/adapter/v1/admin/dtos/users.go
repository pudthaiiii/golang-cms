package dtos

import (
	"time"
)

type CreateUser struct {
	FullName   string `json:"fullName" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6"`
	Mobile     string `json:"mobile" validate:"required,min=10"`
	Company    string `json:"company"`
	IsActive   int    `json:"isActive" validate:"required,one_or_zero"`
	Type       string `json:"type" validate:"required,enum_user_type"`
	RoleID     uint   `json:"roleId" validate:"required"`
	IsAllBU    int    `json:"isAllBu" validate:"required,one_or_zero"`
	MerchantID uint   `json:"merchantId" validate:"required_if_type_not_admin"`
}

type ResponseUserID struct {
	ID uint `json:"id"`
}

type ResponseUserData struct {
	ID              uint       `json:"id"`
	FullName        string     `json:"fullName"`
	Email           string     `json:"email"`
	Mobile          string     `json:"mobile"`
	Company         string     `json:"company"`
	IsActive        int        `json:"isActive"`
	Type            string     `json:"type"`
	RoleID          uint       `json:"roleId"`
	IsAllBU         int        `json:"isAllBu"`
	ProfileImage    string     `json:"profileImage"`
	EmailVerifiedAt *time.Time `json:"emailVerifiedAt"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
}
