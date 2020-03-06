package repsitory

import (
	"context"

	"github.com/zlp-ecommerce/customer-service/models"
)

// PostRepo explain...
type CustomerRepo interface {
	Fetch(ctx context.Context, num int64) ([]*models.Customer, error)
	GetByID(ctx context.Context, id int64) (*models.Customer, error)
	Create(ctx context.Context, p *models.Customer) (int64, error)
	Update(ctx context.Context, p *models.Customer) (*models.Customer, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
