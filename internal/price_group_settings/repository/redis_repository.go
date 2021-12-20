package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/internal/price_group_settings"
)

// News redis repository
type priceGroupRedisRepo struct {
	redisClient *redis.Client
}

// News redis repository constructor
func NewpriceGroupRepo(redisClient *redis.Client) price_group_settings.RedisRepository {
	return &priceGroupRedisRepo{redisClient: redisClient}
}

// Get new by id
func (n *priceGroupRedisRepo) GetPriceGroupCtx(ctx context.Context, key string) (*models.PriceGroupSettingsList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "priceGroupRedisRepo.GetNewsByIDCtx")
	defer span.Finish()

	priceGroupBytes, err := n.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "priceGroupRedisRepo.GetNewsByIDCtx.redisClient.Get")
	}
	priceGroupBase := &models.PriceGroupSettingsList{}
	if err = json.Unmarshal(priceGroupBytes, priceGroupBase); err != nil {
		return nil, errors.Wrap(err, "priceGroupRedisRepo.GetNewsByIDCtx.json.Unmarshal")
	}

	return priceGroupBase, nil
}

// Cache priceGroup item
func (n *priceGroupRedisRepo) SetPriceGroupCtx(ctx context.Context, key string, seconds int, priceGroup *models.PriceGroupSettingsList) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "priceGroupRedisRepo.SetNewsCtx")
	defer span.Finish()

	priceGroupBytes, err := json.Marshal(priceGroup)
	if err != nil {
		return errors.Wrap(err, "priceGroupRedisRepo.SetNewsCtx.json.Marshal")
	}
	if err = n.redisClient.Set(ctx, key, priceGroupBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return errors.Wrap(err, "priceGroupRedisRepo.SetNewsCtx.redisClient.Set")
	}
	return nil
}

// Delete new item from cache
func (n *priceGroupRedisRepo) DeletePriceGroupCtx(ctx context.Context, key string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "priceGroupRedisRepo.DeleteNewsCtx")
	defer span.Finish()
	//fmt.Printf(key)
	//fmt.Printf("LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLl")
	if err := n.redisClient.Del(ctx, key).Err(); err != nil {
		return errors.Wrap(err, "priceGroupRedisRepo.DeleteNewsCtx.redisClient.Del")
	}
	return nil
}
