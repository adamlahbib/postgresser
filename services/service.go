package services

import (
	"context"

	"github.com/adamlahbib/postgresser/models"
)

type Service interface {
	Create(ctx context.Context, request models.CreateRequest) (models.CreateResponse, error)
	Delete(ctx context.Context, request models.DeleteRequest) error
	Update(ctx context.Context, request models.UpdateRequest) error
}
