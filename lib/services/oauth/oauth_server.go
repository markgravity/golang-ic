package oauth

import (
	"context"
	"os"

	"github.com/markgravity/golang-ic/helpers/log"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
)

var oauthServer *server.Server
var clientStore *pg.ClientStore

func SetUpOAuthServer() error {
	manager := manage.NewDefaultManager()
	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()

	err := clientStore.Set(os.Getenv("CLIENT_ID"), &models.Client{
		ID:     os.Getenv("CLIENT_ID"),
		Secret: os.Getenv("CLIENT_SECRET"),
		Domain: os.Getenv("DOMAIN"),
	})
	if err != nil {
		return err
	}

	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)

	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetPasswordAuthorizationHandler(passwordAuthorizationHandler)
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	oauthServer = srv

	return nil
}

func GetOAuthServer() *server.Server {
	return oauthServer
}

func GetClientStore() *pg.ClientStore {
	return clientStore
}

func passwordAuthorizationHandler(ctx context.Context, clientID, email string, password string) (string, error) {
	// TODO: Implement the logic in Sign In task
	return "1", nil
}
