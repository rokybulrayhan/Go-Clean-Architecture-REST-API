//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package price_group_settings

import (
	"context"

	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/pkg/utils"
)

// News use case
type UseCase interface {
	GetAllByNewsID(ctx context.Context, newsID int, query *utils.PaginationQuery) (*models.PriceGroupSettings, error)
	GetAllPriceGroup(ctx context.Context, query *utils.PaginationQuery) (*models.PriceGroupSettings, error)
}
