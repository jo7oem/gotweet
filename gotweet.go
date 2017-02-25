package gotweet

import (
	"encoding/json"
	"github.com/mrjones/oauth"
	"io/ioutil"
)

type TwitterApp struct {
	consumer    oauth.Consumer
	accessToken oauth.AccessToken
}

func NewTwitterapp(APIkye, APIsec, accessToken, accessTokenSec string) *TwitterApp {
	app := new(TwitterApp)
	app.consumer = oauth.NewConsumer(
		APIkye,
		APIsec,
		oauth.ServiceProvider{
			RequestTokenUrl: "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "	https://api.twitter.com/oauth/authorize",
			AccessTokenUrl: "https://api.twitter.com/oauth/access_token"})
	app.accessToken = &oauth.AccessToken{accessToken, accessTokenSec}
	return app
}
func (t *TwitterApp) Get(url string, params map[string]string) (interface{}, error) {
	response, err := t.consumer.Get(url, params, t.accessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// decode
	var result interface{}
	err = json.Unmarshal(b, &result)
	return result, err
}

func (t *TwitterApp) Post(url string, params map[string]string) (interface{}, error) {
	response, err := t.consumer.Post(url, params, t.accessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// decode
	var result interface{}
	err = json.Unmarshal(b, &result)
	return result, err
}
