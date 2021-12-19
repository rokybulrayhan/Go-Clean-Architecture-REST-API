//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package price_group_settings

import (
	"context"

	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/pkg/utils"
)

// News Repository
type Repository interface {
	Create(ctx context.Context, priceGroup *models.PriceGroupSettings) (*models.PriceGroupSettings, error)
	GetAllByNewsID(ctx context.Context, priceGroupID int) (*models.PriceGroupSettings, error)
	GetAllPriceGroup(ctx context.Context, query *utils.PaginationQuery) (*models.PriceGroupSettings, error)
	Update(ctx context.Context, priceGroup *models.PriceGroupSettings) (*models.PriceGroupSettings, error)
	GetAllPriceGroupNew(ctx context.Context, filterQuery string, pq *utils.PaginationQuery) (*models.PriceGroupSettingsList, error)
	Delete(ctx context.Context, id int) error
}
