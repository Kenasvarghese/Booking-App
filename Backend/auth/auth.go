package auth

import (
	"context"
	"encoding/json"
	"errors"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type authProvider struct {
	authConfig       *oauth2.Config
	userInfoEndpoint string
}

type AuthProvider interface {
	GetURL(state string) string
	GetUserEmail(code string) (string, error)
}

// NewAuthProvider creates a new AuthProvider instance using the provided SSO configuration.
func NewAuthProvider(c *ssoConfig) AuthProvider {
	return &authProvider{
		authConfig: &oauth2.Config{
			ClientID:     c.clientID,
			ClientSecret: c.clientSecret,
			RedirectURL:  c.redirectURL,
			Scopes:       []string{"email"},
			Endpoint:     google.Endpoint,
		},
		userInfoEndpoint: "https://www.googleapis.com/oauth2/v2/userinfo",
	}
}

// GetURL generates the OAuth2 authorization URL for the given state.
func (a *authProvider) GetURL(state string) string {
	return a.authConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

// GetUserEmail exchanges the provided authorization code for the user Email,
func (a *authProvider) GetUserEmail(code string) (string, error) {
	token, err := a.authConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	client := a.authConfig.Client(context.Background(), token)
	resp, err := client.Get(a.userInfoEndpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return "", err
	}

	// Extract email from userInfo response
	if email, ok := userInfo["email"].(string); ok {
		return email, nil
	}
	return "", errors.New("email not found in user info")
}
