package repository

import (
	"context"

	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/AleksK1NG/api-mc/internal/price_group_settings"
	"github.com/AleksK1NG/api-mc/pkg/utils"
	"github.com/opentracing/opentracing-go"
)

// News Repository
type priceGroupRepo struct {
	db *sqlx.DB
}

// News repository constructor
func NewPriceGroupRepository(db *sqlx.DB) price_group_settings.Repository {
	return &priceGroupRepo{db: db}
}
func (r *priceGroupRepo) GetAllByNewsID(ctx context.Context, newsID int, query *utils.PaginationQuery) (*models.PriceGroupSettings, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "newsRepo.GetNewsByID")
	defer span.Finish()

	n := &models.PriceGroupSettings{}
	if err := r.db.GetContext(ctx, n, getNewsByID, newsID); err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetNewsByID.GetContext")
	}

	return n, nil
}

func (r *priceGroupRepo) GetAllPriceGroup(ctx context.Context, query *utils.PaginationQuery) (*models.PriceGroupSettings, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "priceGroupRepo.GetAllPriceGroup")
	defer span.Finish()

	n := &models.PriceGroupSettings{}
	if err := r.db.GetContext(ctx, n, getAllPriceGroup); err != nil {
		return nil, errors.Wrap(err, "priceGroupRepo.GetAllPriceGroup.GetContext")
	}

	return n, nil
}
