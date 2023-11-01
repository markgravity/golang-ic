package modelsresponse

type TokenResponse struct {
	ID           string `jsonapi:"primary,token"`
	AccessToken  string `jsonapi:"attr,access_token"`
	TokenType    string `jsonapi:"attr,token_type"`
	ExpiresIn    int64  `jsonapi:"attr,expires_in"`
	RefreshToken string `jsonapi:"attr,refresh_token"`
}
