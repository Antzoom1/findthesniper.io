package config

import (
	"github.com/RagOfJoes/findthesniper.io/domains"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var _ domains.Domain = (*Provider)(nil)

// Provider defines the required configuration for OAuth providers
type Provider struct {
	// URL is where the access token will be used to check whether the user is authenticated or not
	URL string
	// ClientID defines the client id for Provider
	ClientID string
	// ClientSecret defines the client secret for Provider
	ClientSecret string
}

func (p Provider) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.URL, validation.Required, is.URL),
		validation.Field(&p.ClientID, validation.Required),
		validation.Field(&p.ClientSecret, validation.Required),
	)
}
