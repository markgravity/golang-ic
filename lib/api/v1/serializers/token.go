package serializers

import (
	"time"

	responsemodels "github.com/markgravity/golang-ic/lib/models/response"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/google/uuid"
)

type TokenSerializer struct {
	Token     oauth2.TokenInfo
	TokenType string
}

func (s *TokenSerializer) Data() (response *responsemodels.TokenResponse) {
	response = &responsemodels.TokenResponse{
		ID:           uuid.NewString(),
		AccessToken:  s.Token.GetAccess(),
		TokenType:    s.TokenType,
		ExpiresIn:    int64(s.Token.GetAccessExpiresIn() / time.Second),
		RefreshToken: s.Token.GetRefresh(),
	}

	return response
}
