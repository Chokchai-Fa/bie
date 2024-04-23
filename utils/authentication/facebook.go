package authentication

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

func NewFacebookAuth(clientID string, clientSecret, redirectUrl string) (*oauth2.Config, error) {
	facebookConf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Scopes:       []string{"email"},
		Endpoint:     facebook.Endpoint,
	}

	return facebookConf, nil
}

func FacebookLogin(facebookConfig oauth2.Config) string {
	url := facebookConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)

	return url
}

func FacebookCallBack(code string, facebookConfig oauth2.Config, ctx *gin.Context) (*oauth2.Token, error) {
	token, err := facebookConfig.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return token, err
}
