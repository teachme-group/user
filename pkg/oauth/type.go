package oauth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

type Provider string

const (
	Google Provider = "google"
	GitHub Provider = "github"
)

type (
	ProvidersConfig map[Provider]config
	providersCreds  map[Provider]providerConfig
)

var (
	endpoints = map[Provider]oauth2.Endpoint{
		Google: google.Endpoint,
		GitHub: github.Endpoint,
	}
	authURLs = map[Provider]string{
		Google: "https://accounts.google.com/o/oauth2/auth",
		GitHub: "https://api.github.com/user",
	}
)

func (p providersCreds) Fill(cfg ProvidersConfig) {
	for provider, config := range cfg {
		p[provider] = providerConfig{
			authURL: authURLs[provider],
			oauth: &oauth2.Config{
				ClientID:     config.ClientID,
				ClientSecret: config.ClientSecret,
				RedirectURL:  config.RedirectURL,
				Scopes:       config.Scopes,
				Endpoint:     endpoints[provider],
			},
		}
	}
}
