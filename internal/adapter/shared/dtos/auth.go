package dtos

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type Refresh struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type AuthJWT struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
