package mysql

import (
	"context"

	"github.com/RagOfJoes/findthesniper.io/domains"
	"github.com/RagOfJoes/findthesniper.io/repositories"
	"github.com/uptrace/bun"
)

var _ repositories.Challenge = (*challenge)(nil)

type challenge struct {
	db *bun.DB
}

func NewChallenge(db *bun.DB) repositories.Challenge {
	return &challenge{
		db: db,
	}
}

func (c *challenge) Create(ctx context.Context, newChallenge domains.Challenge) (*domains.Challenge, error) {
	if _, err := c.db.NewInsert().Model(&newChallenge).Exec(ctx); err != nil {
		return nil, err
	}

	return &newChallenge, nil
}

func (c *challenge) Get(ctx context.Context, id string) (*domains.Challenge, error) {
	var foundChallenge domains.Challenge

	if err := c.db.NewSelect().
		Model(&foundChallenge).
		Where("`challenge`.id = ?", id).
		Relation("User").
		Scan(ctx); err != nil {
		return nil, err
	}

	return &foundChallenge, nil
}

func (c *challenge) Update(ctx context.Context, updateChallenge domains.Challenge) (*domains.Challenge, error) {
	if _, err := c.db.NewUpdate().Model(&updateChallenge).WherePK().Exec(ctx); err != nil {
		return nil, err
	}

	return &updateChallenge, nil
}

func (c *challenge) Delete(ctx context.Context, id string) error {
	if _, err := c.db.NewDelete().Model(&domains.Challenge{}).Where("id = ?", id).Exec(ctx); err != nil {
		return err
	}

	return nil
}
