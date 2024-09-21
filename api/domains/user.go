package domains

import (
	"time"

	"github.com/RagOfJoes/findthesniper.io/internal"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/oklog/ulid/v2"
	"github.com/uptrace/bun"
)

var _ Domain = (*User)(nil)

// User defines a user
type User struct {
	bun.BaseModel

	// ID defines the unique id for the user
	ID string `bun:"type:varchar(26),pk,notnull" json:"id"`
	// State defines the current state of the User
	State string `bun:"type:varchar(8),default:'PENDING',notnull" json:"state"`
	// Username defines the username of the User
	Username string `bun:"type:varchar(64),unique,notnull" json:"username"`

	// CreatedAt defines when the user was created
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	// UpdatedAt defines when the user was last updated
	UpdatedAt bun.NullTime `bun:",nullzero,default:NULL" json:"updated_at"`
	// DeletedAt defines when and if the user was deleted
	DeletedAt bun.NullTime `bun:",soft_delete,nullzero,default:NULL" json:"-"`
}

// NewUser creates a new user
func NewUser() User {
	return User{
		ID:       ulid.Make().String(),
		State:    "PENDING",
		Username: ulid.Make().String(),

		CreatedAt: time.Now(),
	}
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.ID, validation.Required, validation.By(internal.IsULID)),
		validation.Field(&u.State, validation.Required, validation.In("PENDING", "COMPLETE")),
		validation.Field(&u.Username, validation.Required, validation.Length(4, 64), is.Alphanumeric),

		validation.Field(&u.CreatedAt, validation.Required),
		validation.Field(&u.UpdatedAt, validation.When(!u.UpdatedAt.IsZero(), validation.By(internal.IsAfter(u.CreatedAt)))),
		validation.Field(&u.DeletedAt, validation.When(!u.DeletedAt.IsZero(), validation.By(internal.IsAfter(u.CreatedAt)))),
	)
}
