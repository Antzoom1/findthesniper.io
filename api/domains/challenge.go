package domains

import (
	"time"

	"github.com/RagOfJoes/findthesniper.io/internal"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/uptrace/bun"
)

var _ Domain = (*Challenge)(nil)

// Challenge
type Challenge struct {
	bun.BaseModel

	// ID defines the unique id for the game
	ID string `bun:"type:varchar(26),pk,notnull" json:"id"`
	// Prompt defines the prompt for the game
	Prompt string `bun:"type:varchar(256),notnull" json:"prompt"`

	// CreatedAt defines when the user was created
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	// UpdatedAt defines when the user was last updated
	UpdatedAt bun.NullTime `bun:",nullzero,default:NULL" json:"updated_at"`
	// DeletedAt defines when and if the user was deleted
	DeletedAt bun.NullTime `bun:",soft_delete,nullzero,default:NULL" json:"-"`

	// ImageID defines the unique id of the image for the challenge
	ImageID string `bun:"type:varchar(26)" json:"-"`
	// Image is the image for the challenge
	Image Image `bun:"rel:belongs-to,join:image_id=id" json:"image"`

	UserID string `bun:"type:varchar(26)" json:"-"`
	// User is the user that created the challenge
	User User `bun:"rel:belongs-to,join:user_id=id" json:"user"`
}

func (c Challenge) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ID, validation.Required, validation.By(internal.IsULID)),
		validation.Field(&c.Prompt, validation.Required, validation.Length(1, 256)),

		validation.Field(&c.CreatedAt, validation.Required),
		validation.Field(&c.UpdatedAt, validation.When(!c.UpdatedAt.IsZero(), validation.By(internal.IsAfter(c.CreatedAt)))),
		validation.Field(&c.DeletedAt, validation.When(!c.DeletedAt.IsZero(), validation.By(internal.IsAfter(c.CreatedAt)))),

		validation.Field(&c.ImageID, validation.Required, validation.By(internal.IsULID)),
		validation.Field(&c.Image, validation.Required),

		validation.Field(&c.UserID, validation.Required, validation.By(internal.IsULID)),
		validation.Field(&c.User, validation.Required),
	)
}
