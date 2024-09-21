package config

import (
	"github.com/RagOfJoes/findthesniper.io/domains"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ domains.Domain = (*Providers)(nil)

// Providers defines the required configuration for supported OAuth providers
type Providers struct {
	Discord Provider
	GitHub  Provider
	Google  Provider
}

func (p Providers) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Discord, validation.Required),
		validation.Field(&p.GitHub, validation.Required),
		validation.Field(&p.Google, validation.Required),
	)
}
