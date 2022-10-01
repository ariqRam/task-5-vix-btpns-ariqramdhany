package models

type User struct {
	CommonModel
	Username     *string `json:"username" validate:"required"`
	Email        *string `json:"email" validate:"required"`
	Password     *string `json:"password" validate:"email, required"`
	Token        *string `json:"token"`
	RefreshToken *string `json:"refresh_token"`
	UserType     *string `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
}
