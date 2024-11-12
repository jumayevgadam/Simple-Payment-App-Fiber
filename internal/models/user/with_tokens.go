package user

// UserWithTokens model is
type UserWithTokens struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	User         AllUserDTO
}
