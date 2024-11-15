package user

// UserWithTokens model is user details with token.
type UserWithTokens struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	User         AllUserDTO
}
