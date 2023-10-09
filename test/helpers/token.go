package helpers

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/markgravity/golang-ic/lib/services/oauth"
	"os"
)

func GenerateToken(userID string) string {
	request := oauth2.TokenGenerateRequest{
		ClientID:            os.Getenv("CLIENT_ID"),
		ClientSecret:        os.Getenv("CLIENT_SECRET"),
		UserID:              userID,
		RedirectURI:         "",
		Scope:               "",
		Code:                "",
		CodeChallenge:       "",
		CodeChallengeMethod: "",
		Refresh:             "",
		CodeVerifier:        "",
		AccessTokenExp:      0,
		Request:             nil,
	}
	server := oauth.GetOAuthServer()
	tokenInfo, _ := server.Manager.GenerateAccessToken(context.Background(), "password", &request)

	return tokenInfo.GetAccess()
}
