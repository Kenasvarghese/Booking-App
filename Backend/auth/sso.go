package auth

type ssoConfig struct {
	clientID     string
	clientSecret string
	redirectURL  string
}

type option func(c *ssoConfig)


// NewSSOConfig creates a new ssoConfig and applies the provided options.
func NewSSOConfig(options ...option) *ssoConfig {
	c := &ssoConfig{}
	for _, option := range options {
		option(c)
	}
	return c
}

// WithClientID returns an option that sets the clientID field in ssoConfig.
func WithClientID(clientID string) option {
	return func(c *ssoConfig) {
		c.clientID = clientID
	}
}

// WithClientSecret returns an option that sets the clientSecret field in ssoConfig.
func WithClientSecret(clientSecret string) option {
	return func(c *ssoConfig) {
		c.clientSecret = clientSecret
	}
}

// WithRedirectURL returns an option that sets the redirectURL field in ssoConfig.
func WithRedirectURL(redirectURL string) option {
	return func(c *ssoConfig) {
		c.redirectURL = redirectURL
	}
}

