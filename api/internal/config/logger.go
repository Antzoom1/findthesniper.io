package config

import (
	"time"

	"github.com/RagOfJoes/findthesniper.io/domains"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"
)

var _ domains.Domain = (*Logger)(nil)

// Logger defines the configuration for the logger
type Logger struct {
	// Level is the log level
	Level int
	// ReportCaller determines whether the calling method will be reported
	ReportCaller bool
}

func (l Logger) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Level, validation.Min(0), validation.Max(6)),
		validation.Field(&l.ReportCaller),
	)
}

// SetupLogger configures the logger with the provided configuration
func SetupLogger(config Configuration, logger *logrus.Logger) error {
	logrus.SetLevel(logrus.Level(config.Logger.Level))
	logrus.SetReportCaller(config.Logger.ReportCaller)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: false,
		DisableColors:    false,
		DisableQuote:     false,
		FullTimestamp:    true,
		TimestampFormat:  time.RFC3339,
	})

	return nil
}
