package config

// AccessTokenConfig for access token.
type AccessTokenConfig struct {
	AccessTokenSecret string `envconfig:"ACCESS_JWT_SECRET_KEY" validate:"required"`
	AccessTokenName   string `envconfig:"ACCESS_TOKEN_NAME" validate:"required"`
}

// RefreshTokenConfig for refresh token.
type RefreshTokenConfig struct {
	RefreshTokenSecret string `envconfig:"REFRESH_JWT_SECRET_KEY" validate:"required"`
	RefreshTokenName   string `envconfig:"REFRESH_TOKEN_NAME" validate:"required"`
}

// embedding two structs, we got JWTOps struct.
type JWTOps struct {
	AccessTokenConfig
	RefreshTokenConfig
}
