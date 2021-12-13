package usecase

import (
	"context"

	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/models"
	price_group "github.com/AleksK1NG/api-mc/internal/price_group_settings"

	"github.com/AleksK1NG/api-mc/pkg/logger"
	"github.com/AleksK1NG/api-mc/pkg/utils"
	"github.com/opentracing/opentracing-go"
)

const (
	basePrefix    = "api-news:"
	cacheDuration = 3600
)

// News UseCase
type priceGroupUC struct {
	cfg            *config.Config
	priceGroupRepo price_group.Repository
	logger         logger.Logger
}

// News UseCase constructor
func NewPriceGroupUseCase(cfg *config.Config, priceGroupRepo price_group.Repository, logger logger.Logger) price_group.UseCase {
	return &priceGroupUC{cfg: cfg, priceGroupRepo: priceGroupRepo, logger: logger}
}

//GetAllByNewsID(ctx context.Context, newsID uuid.UUID, query *utils.PaginationQuery) (*models.PriceGroupSettings, error)

func (u *priceGroupUC) GetAllByNewsID(ctx context.Context, newsID int, query *utils.PaginationQuery) (*models.PriceGroupSettings, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "commentsUC.GetAllByNewsID")
	defer span.Finish()

	return u.priceGroupRepo.GetAllByNewsID(ctx, newsID, query)
	//ctx, newsID, query)
}

func (u *priceGroupUC) GetAllPriceGroup(ctx context.Context, query *utils.PaginationQuery) (*models.PriceGroupSettings, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "commentsUC.GetAllByNewsID")
	defer span.Finish()

	return u.priceGroupRepo.GetAllPriceGroup(ctx, query)
	//ctx, newsID, query)
}

// Create news
