package DOIpUpdate

import (
	"golang.org/x/oauth2"
	"github.com/digitalocean/godo"
	"golang.org/x/net/context"
)


type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func GetClientWithToken(token string) (*godo.Client) {
	ts := &TokenSource{
		AccessToken: token,
	}

	oauthClient := oauth2.NewClient(context.TODO(), ts);
	client := godo.NewClient(oauthClient)

	return client
}