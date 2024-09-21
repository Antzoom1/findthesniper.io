package domains

import (
	"time"

	"github.com/RagOfJoes/findthesniper.io/internal"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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

	// ImageURL defines the ImageURL for the image for the challenge
	ImageURL string `bun:"type:text,notnull" json:"url"`
	// ImageHeight defines the original height of the image for the challenge
	ImageHeight float32 `bun:"type:float,notnull" json:"height"`
	// ImageWidth defines the original width of the image for the challenge
	ImageWidth float32 `bun:"type:float,notnull" json:"width"`

	// SniperHeight defines the height of the sniper object in the image relative to the original image's dimensions
	SniperHeight float32 `bun:"type:float,notnull" json:"sniper_height"`
	// SniperWidth defines the width of the sniper object in the image relative to the original image's dimensions
	SniperWidth float32 `bun:"type:float,notnull" json:"sniper_width"`
	// SniperX defines the x coordinate of the center of the circle in the image relative to the original image's dimensions
	SniperX float32 `bun:"type:float,notnull" json:"sniper_x"`
	// SniperY defines the y coordinate of the center of the circle in the image relative to the original image's dimensions
	SniperY float32 `bun:"type:float,notnull" json:"sniper_y"`

	// CreatedAt defines when the user was created
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	// UpdatedAt defines when the user was last updated
	UpdatedAt bun.NullTime `bun:",nullzero,default:NULL" json:"updated_at"`
	// DeletedAt defines when and if the user was deleted
	DeletedAt bun.NullTime `bun:",soft_delete,nullzero,default:NULL" json:"-"`

	// UserID defines unique id for the user that created the challenge
	UserID string `bun:"type:varchar(26)" json:"-"`
	// User is the user that created the challenge
	User User `bun:"rel:belongs-to,join:user_id=id" json:"user"`
}

func (c Challenge) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ID, validation.Required, validation.By(internal.IsULID)),
		validation.Field(&c.Prompt, validation.Required, validation.Length(1, 256)),
		validation.Field(&c.ImageURL, validation.Required, is.URL),

		validation.Field(&c.ImageHeight, validation.Required, validation.Min(0)),
		validation.Field(&c.ImageWidth, validation.Required, validation.Min(0)),

		validation.Field(&c.SniperHeight, validation.Required, validation.Max(c.ImageHeight), validation.Min(0)),
		validation.Field(&c.SniperWidth, validation.Required, validation.Max(c.ImageWidth), validation.Min(0)),
		validation.Field(&c.SniperX, validation.Required, validation.Min(c.ImageWidth-c.SniperWidth/2), validation.Min(c.SniperWidth/2)),
		validation.Field(&c.SniperY, validation.Required, validation.Max(c.ImageHeight-c.SniperHeight/2), validation.Min(c.SniperHeight/2)),

		validation.Field(&c.CreatedAt, validation.Required),
		validation.Field(&c.UpdatedAt, validation.When(!c.UpdatedAt.IsZero(), validation.By(internal.IsAfter(c.CreatedAt)))),
		validation.Field(&c.DeletedAt, validation.When(!c.DeletedAt.IsZero(), validation.By(internal.IsAfter(c.CreatedAt)))),

		validation.Field(&c.UserID, validation.Required, validation.By(internal.IsULID)),
		validation.Field(&c.User, validation.Required),
	)
}
