package user

// UserWithTokens model is user details with token.
type UserWithTokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	User         AllUserDTO
}
