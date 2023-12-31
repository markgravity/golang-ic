package oauth

import (
	"context"
	"os"

	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/helpers"
	"github.com/markgravity/golang-ic/helpers/log"
	models "github.com/markgravity/golang-ic/lib/models"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	oauthmodels "github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
)

var oauthServer *server.Server
var clientStore *pg.ClientStore

func SetUpOAuthServer() error {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()

	client := oauthmodels.Client{
		ID:     os.Getenv("CLIENT_ID"),
		Secret: os.Getenv("CLIENT_SECRET"),
		Domain: os.Getenv("DOMAIN"),
	}
	err := clientStore.Set(client.ID, &client)

	if err != nil {
		return err
	}

	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)

	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetPasswordAuthorizationHandler(passwordAuthorizationHandler)
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("OAuth Internal Error:", err.Error())
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
	db := database.GetDB()

	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Error(err)
		return "", err
	}

	err = helpers.ComparePassword(user.EncryptedPassword, password)
	if err != nil {
		log.Error("Incorrect username or password", err)
		return "", err
	}

	return user.Base.ID.String(), nil
}
