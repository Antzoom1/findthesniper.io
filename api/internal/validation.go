package internal

import (
	"database/sql"
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/oklog/ulid/v2"
	"github.com/uptrace/bun"
)

// IsULID is a custom validation rule that checks if the value is a valid ULID
func IsULID(value interface{}) error {
	switch value.(type) {
	case string:
		str, ok := value.(string)
		if !ok {
			return errors.New("must be a valid ULID")
		}

		if _, err := ulid.Parse(str); err != nil {
			return err
		}

		break
	case sql.NullString:
		str, ok := value.(sql.NullString)
		if !ok {
			return errors.New("must be a valid ULID")
		}

		if _, err := ulid.Parse(str.String); err != nil {
			return err
		}

		break
	}

	return nil
}

// IsAfter is a custom validation rule that checks if the value is after a given time
func IsAfter(a time.Time) validation.RuleFunc {
	return func(value interface{}) error {
		switch value.(type) {
		case time.Time:
			t, _ := value.(time.Time)
			if t.Before(a) {
				return errors.New("must be after " + a.String())
			}

			return nil
		case bun.NullTime:
			t, _ := value.(bun.NullTime)
			if t.IsZero() {
				return errors.New("must not be zero")
			}
			if t.Time.Before(a) {
				return errors.New("must be after " + a.String())
			}

			return nil
		}

		return errors.New("must be a valid time.Time or bun.NullTime")
	}
}

// IsBefore is a custom validation rule that checks if the value is before a given time
func IsBefore(b time.Time) validation.RuleFunc {
	return func(value interface{}) error {
		switch value.(type) {
		case time.Time:
			t, _ := value.(time.Time)
			if t.After(b) {
				return errors.New("must be before " + b.String())
			}

			return nil
		case bun.NullTime:
			t, _ := value.(bun.NullTime)
			if t.IsZero() {
				return errors.New("must not be zero")
			}
			if t.Time.After(b) {
				return errors.New("must be before " + b.String())
			}

			return nil
		}

		return errors.New("must be a valid time.Time or bun.NullTime")
	}
}
