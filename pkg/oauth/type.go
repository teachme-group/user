package oauth

import (
	"encoding/json"

	"github.com/teachme-group/user/internal/domain"
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
		Google: "https://www.googleapis.com/oauth2/v3/userinfo",
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

var unmarshaler = map[Provider]func([]byte) (Response, error){
	Google: unmarshalGoogleResponse,
}

func unmarshalGoogleResponse(data []byte) (Response, error) {
	var resp GoogleAuthResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type (
	Response interface {
		ToUser() domain.User
	}

	GoogleAuthResponse struct {
		Sub           string `json:"sub"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
	}
)

func (r GoogleAuthResponse) ToUser() domain.User {
	return domain.User{
		Login:         r.Name,
		Email:         r.Email,
		EmailVerified: r.EmailVerified,
		ProfileImage:  r.Picture,
	}
}
