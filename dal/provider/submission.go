package provider

import (
	"context"

	"github.com/evscott/z3-e2c-api/models"
)

func (c *Config) CreateSubmission(ctx context.Context, submission *models.Submission) error {
	return c.db.Insert(submission)
}
