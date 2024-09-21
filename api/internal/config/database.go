package config

import (
	"github.com/RagOfJoes/findthesniper.io/domains"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var _ domains.Domain = (*Database)(nil)

// Database defines the configuration for the database
type Database struct {
	// Driver is the database driver
	Driver string
	// Host is the database host
	Host string
	// Name is the database name
	Name string
	// Password is the database password
	Password string
	// Port is the database port
	Port string
	// User is the database user
	User string
}

func (d Database) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Driver, validation.Required, validation.In("mysql", "postgres")),
		validation.Field(&d.Host, validation.Required),
		validation.Field(&d.Name, validation.Required),
		validation.Field(&d.Password, validation.Required),
		validation.Field(&d.Port, validation.Required, is.Port),
		validation.Field(&d.User, validation.Required),
	)
}
