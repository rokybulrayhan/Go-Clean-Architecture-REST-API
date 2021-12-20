package repository

/*import (
	"context"
	"log"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/internal/price_group_settings"
)

func SetupRedis() price_group_settings.RedisRepository {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	priceGroupRedisRepo := priceGroupRedisRepo(client)
	return priceGroupRedisRepo
}

func TestNewsRedisRepo_SetNewsCtx(t *testing.T) {
	t.Parallel()

	priceGroupRedisRepo := SetupRedis()

	t.Run("SetNewsCtx", func(t *testing.T) {
		priceGroupUID := uuid.New()
		key := "key"
		n := &models.NewsBase{
			NewsID:  priceGroupUID,
			Title:   "Title",
			Content: "Content",
		}

		err := priceGroupRedisRepo.SetNewsCtx(context.Background(), key, 10, n)
		require.NoError(t, err)
		require.Nil(t, err)
	})
}

func TestNewsRedisRepo_GetNewsByIDCtx(t *testing.T) {
	t.Parallel()

	priceGroupRedisRepo := SetupRedis()

	t.Run("GetNewsByIDCtx", func(t *testing.T) {
		priceGroupUID := uuid.New()
		key := "key"
		n := &models.NewsBase{
			NewsID:  priceGroupUID,
			Title:   "Title",
			Content: "Content",
		}

		priceGroupBase, err := priceGroupRedisRepo.GetNewsByIDCtx(context.Background(), key)
		require.Nil(t, priceGroupBase)
		require.NotNil(t, err)

		err = priceGroupRedisRepo.SetNewsCtx(context.Background(), key, 10, n)
		require.NoError(t, err)
		require.Nil(t, err)

		priceGroupBase, err = priceGroupRedisRepo.GetNewsByIDCtx(context.Background(), key)
		require.NoError(t, err)
		require.Nil(t, err)
		require.NotNil(t, priceGroupBase)
	})
}

func TestNewsRedisRepo_DeleteNewsCtx(t *testing.T) {
	t.Parallel()

	priceGroupRedisRepo := SetupRedis()

	t.Run("SetNewsCtx", func(t *testing.T) {
		key := "key"

		err := priceGroupRedisRepo.DeleteNewsCtx(context.Background(), key)
		require.NoError(t, err)
		require.Nil(t, err)
	})
}
*/
