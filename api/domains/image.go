package domains

import (
	"github.com/RagOfJoes/findthesniper.io/internal"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/uptrace/bun"
)

var _ Domain = (*Image)(nil)

type Image struct {
	bun.BaseModel

	// ID defines the unique id for the image
	ID  string `bun:"type:varchar(26),pk,notnull" json:"id"`
	URL string `bun:"type:text,notnull" json:"url"`

	Height float32 `bun:"type:float,notnull" json:"height"`
	Width  float32 `bun:"type:float,notnull" json:"width"`

	SniperHeight float32 `bun:"type:float,notnull" json:"sniper_height"`
	SniperWidth  float32 `bun:"type:float,notnull" json:"sniper_width"`
	SniperX      float32 `bun:"type:float,notnull" json:"sniper_x"`
	SniperY      float32 `bun:"type:float,notnull" json:"sniper_y"`
}

func (i Image) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.ID, validation.Required, validation.By(internal.IsULID)),
		validation.Field(&i.URL, validation.Required, is.URL),

		validation.Field(&i.Height, validation.Required, validation.Min(0)),
		validation.Field(&i.Width, validation.Required, validation.Min(0)),

		validation.Field(&i.SniperHeight, validation.Required, validation.Min(0)),
		validation.Field(&i.SniperWidth, validation.Required, validation.Min(0)),
		validation.Field(&i.SniperX, validation.Required),
		validation.Field(&i.SniperY, validation.Required),
	)
}
