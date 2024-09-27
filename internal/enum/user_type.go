package enum

type UserTypeEnum string

const (
	ADMIN    UserTypeEnum = "ADMIN"
	MERCHANT UserTypeEnum = "MERCHANT"
	USER     UserTypeEnum = "USER"
)
