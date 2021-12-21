//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package price_group_settings

import (
	"context"

	"github.com/AleksK1NG/api-mc/internal/models"
)

// News use case
type UseCase interface {
	Create(ctx context.Context, priceGroup *models.PriceGroupSettings) (*models.PriceGroupSettings, error)
	GetAllByNewsID(ctx context.Context, priceGroupID int) (*models.PriceGroupSettings, error)
	//	GetAllPriceGroup(ctx context.Context, query *utils.PaginationQuery) (*models.PriceGroupSettings, error)
	//	Update(ctx context.Context, priceGroup *models.PriceGroupSettings) (*models.PriceGroupSettings, error)
	//	GetAllPriceGroupNew(ctx context.Context, filterQuery string, pq *utils.PaginationQuery) (*models.PriceGroupSettingsList, error)
	//	Delete(ctx context.Context, id int) error
}
