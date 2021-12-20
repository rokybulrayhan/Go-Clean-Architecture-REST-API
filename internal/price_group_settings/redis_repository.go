//go:generate mockgen -source redis_repository.go -destination mock/redis_repository_mock.go -package mock
package price_group_settings

import (
	"context"

	"github.com/AleksK1NG/api-mc/internal/models"
)

// News redis repository
type RedisRepository interface {
	GetPriceGroupCtx(ctx context.Context, key string) (*models.PriceGroupSettingsList, error)
	SetPriceGroupCtx(ctx context.Context, key string, seconds int, priceGroup *models.PriceGroupSettingsList) error
	DeletePriceGroupCtx(ctx context.Context, key string) error
}
