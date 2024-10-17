package repositories

import (
	"context"

	"github.com/RagOfJoes/findthesniper.io/domains"
)

// Challenge defines methods for a challenge repository
type Challenge interface {
	// Create creates a new session
	Create(ctx context.Context, newChallenge domains.Challenge) (*domains.Challenge, error)

	// Get retrieves a challenge with its id
	Get(ctx context.Context, id string) (*domains.Challenge, error)

	// Update updates a challenge
	Update(ctx context.Context, updateChallenge domains.Challenge) (*domains.Challenge, error)

	// Delete deletes a challenge
	Delete(ctx context.Context, id string) error
}
