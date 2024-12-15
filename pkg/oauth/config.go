package oauth

import "golang.org/x/oauth2"

type (
	config struct {
		RedirectURL  string   `yaml:"redirect_url"`
		ClientID     string   `yaml:"client_id"`
		ClientSecret string   `yaml:"client_secret"`
		Scopes       []string `yaml:"scopes"`
		Endpoint     string   `yaml:"endpoint"`
	}
	providerConfig struct {
		authURL string
		oauth   *oauth2.Config
	}
)
