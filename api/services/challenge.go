package services

import (
	"context"
	"errors"

	"github.com/RagOfJoes/findthesniper.io/domains"
	"github.com/RagOfJoes/findthesniper.io/internal"
	"github.com/RagOfJoes/findthesniper.io/repositories"
	"github.com/sirupsen/logrus"
)

// Errors
var (
	ErrChallengeCreate       = errors.New("Failed to create challenge.")
	ErrChallengeDoesNotExist = errors.New("Challenge does not exist.")
)

// Challenge defines the challenge service
type Challenge struct {
	repository repositories.Challenge
}

type ChallengeDependencies struct {
	Repository repositories.Challenge
}

// NewChallenge instantiates a challenge service
func NewChallenge(dependencies ChallengeDependencies) Challenge {
	logrus.Print("Created Challenge Service")

	return Challenge{
		repository: dependencies.Repository,
	}
}

func (c *Challenge) New(ctx context.Context, newChallenge domains.Challenge) (*domains.Challenge, error) {
	if err := newChallenge.Validate(); err != nil {
		return nil, internal.NewErrorf(internal.ErrorCodeBadRequest, "%v", err)
	}

	createdChallenge, err := c.repository.Create(ctx, newChallenge)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeInternal, "%v", ErrChallengeCreate)
	}
	if err := createdChallenge.Validate(); err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeInternal, "%v", ErrChallengeCreate)
	}

	return createdChallenge, nil
}

// Find retrieves a user with their id. If strict is set to true then only completed users will be returned
func (c *Challenge) Find(ctx context.Context, id string) (*domains.Challenge, error) {
	foundChallenge, err := c.repository.Get(ctx, id)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "%v", ErrChallengeDoesNotExist)
	}
	if err := foundChallenge.Validate(); err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "%v", ErrChallengeDoesNotExist)
	}

	return foundChallenge, nil
}
