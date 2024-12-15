package oauth

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"github.com/teachme-group/user/internal/domain"
	"github.com/teachme-group/user/pkg/errlist"
	"golang.org/x/oauth2"
)

type oauth struct {
	oauthCfg providersCreds
}

func New(cfg ProvidersConfig) *oauth {
	creds := providersCreds{}
	creds.Fill(cfg)

	return &oauth{
		oauthCfg: creds,
	}
}

func (o *oauth) AuthCodeURLs(state string, provider *string, opts ...oauth2.AuthCodeOption) (map[string]string, error) {
	result := make(map[string]string)
	if provider == nil {
		for prov, creds := range o.oauthCfg {
			authCodeOpts := append(opts, oauth2.SetAuthURLParam("response_type", "code"))
			result[string(prov)] = creds.oauth.AuthCodeURL(state, authCodeOpts...)
		}
		return result, nil
	}

	crds, ok := o.oauthCfg[Provider(*provider)]
	if !ok {
		return nil, errlist.ErrProviderNotFound
	}

	authCodeOpts := append(opts, oauth2.SetAuthURLParam("response_type", "code"))
	result[*provider] = crds.oauth.AuthCodeURL(state, authCodeOpts...)

	return result, nil
}

func (o *oauth) exchangeToken(ctx context.Context, provider Provider, code string) (*oauth2.Token, error) {
	crds, ok := o.oauthCfg[provider]
	if !ok {
		return nil, errlist.ErrProviderNotFound
	}

	return crds.oauth.Exchange(ctx, code)
}

func (o *oauth) ProcessCallback(ctx context.Context, provider Provider, _ []byte, callBackURL string) (user domain.User, err error) {
	crds, ok := o.oauthCfg[provider]
	if !ok {
		return user, errlist.ErrProviderNotFound
	}

	parsedParams, err := url.ParseQuery(callBackURL)
	if err != nil {
		return user, fmt.Errorf("failed to parse callback URL: %v", err)
	}

	code := parsedParams.Get("code")
	token, err := o.exchangeToken(ctx, provider, code)
	if err != nil {
		return user, fmt.Errorf("failed to exchange token: %v", err)
	}

	client := crds.oauth.Client(ctx, token)

	resp, err := client.Get(crds.authURL)
	if err != nil {
		return user, fmt.Errorf("failed to fetch user profile: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return user, fmt.Errorf("failed to fetch user profile: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return user, fmt.Errorf("failed to read response body: %v", err)
	}

	unmarshal, ok := unmarshaler[provider]
	if !ok {
		return user, fmt.Errorf("unmarshaler not found for provider %s", provider)
	}

	unmarshalResponse, err := unmarshal(body)
	if err != nil {
		return user, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return unmarshalResponse.ToUser(), nil
}
