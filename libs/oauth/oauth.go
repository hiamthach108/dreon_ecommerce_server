package oauth

import (
	"context"
	"dreon_ecommerce_server/configs"
	"dreon_ecommerce_server/shared/constants"
	"dreon_ecommerce_server/shared/helpers"
	"dreon_ecommerce_server/shared/interfaces"
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type appOAuth struct {
	logger    interfaces.ILogger
	config    *configs.AppConfig
	googleCfg *oauth2.Config
}

type UserData struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
}

type IAppOAuth interface {
	GetGoogleOAuthUrl() (url string, state string)
	GetGoogleUserInfo(ctx context.Context, code string) (userInfo *UserData, err error)
}

func NewAppOAuth(config *configs.AppConfig, logger interfaces.ILogger) *appOAuth {
	googleCfg := &oauth2.Config{
		ClientID:     config.Auth.Google.ClientID,
		ClientSecret: config.Auth.Google.ClientSecret,
		RedirectURL:  fmt.Sprintf("http://%s:%s/auth/google/callback", config.Server.Host, config.Server.Port),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	return &appOAuth{
		logger,
		config,
		googleCfg,
	}
}

func (a *appOAuth) GetGoogleOAuthUrl() (url string, state string) {
	state, _ = helpers.GenerateRandomString(constants.REFRESH_TOKEN_OAUTH_LEN)
	url = a.googleCfg.AuthCodeURL(state)
	return
}

func (a *appOAuth) GetGoogleUserInfo(ctx context.Context, code string) (userInfo *UserData, err error) {
	token, err := a.googleCfg.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	client := a.googleCfg.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	userInfo = &UserData{}
	err = json.Unmarshal(userData, userInfo)
	if err != nil {
		return nil, err
	}

	return
}
