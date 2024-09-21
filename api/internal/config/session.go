package config

import (
	"time"

	"github.com/RagOfJoes/findthesniper.io/domains"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ domains.Domain = (*Session)(nil)

// Session defines the configuration for session management
type Session struct {
	// Lifetime controls how long a session can be valid for
	//
	// Default: 336h (2 weeks)
	Lifetime time.Duration
	// Cookie is the configuration for the session cookie
	Cookie Cookie
}

func (s Session) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Lifetime, validation.Required),
		validation.Field(&s.Cookie),
	)
}
