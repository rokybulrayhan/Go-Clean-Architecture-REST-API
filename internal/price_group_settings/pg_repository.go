//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package price_group_settings

import (
	"context"

	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/pkg/utils"
)

// News Repository
type Repository interface {
	GetAllByNewsID(ctx context.Context, newsID int, query *utils.PaginationQuery) (*models.PriceGroupSettings, error)
	GetAllPriceGroup(ctx context.Context, query *utils.PaginationQuery) (*models.PriceGroupSettings, error)
}
