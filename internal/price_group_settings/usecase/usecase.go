package usecase

import (
	"context"
	"fmt"

	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/models"
	price_group "github.com/AleksK1NG/api-mc/internal/price_group_settings"

	"github.com/AleksK1NG/api-mc/pkg/logger"
	"github.com/AleksK1NG/api-mc/pkg/utils"
	"github.com/opentracing/opentracing-go"
)

const (
	basePrefix    = "api-priceGroup:"
	cacheDuration = 3600
)

// News UseCase
type priceGroupUC struct {
	cfg            *config.Config
	priceGroupRepo price_group.Repository
	redisRepo      price_group.RedisRepository
	logger         logger.Logger
}

// News UseCase constructor
func NewPriceGroupUseCase(cfg *config.Config, priceGroupRepo price_group.Repository, redisRepo price_group.RedisRepository, logger logger.Logger) price_group.UseCase {
	return &priceGroupUC{cfg: cfg, priceGroupRepo: priceGroupRepo, redisRepo: redisRepo, logger: logger}
}

//GetAllByNewsID(ctx context.Context, priceGroupID uuid.UUID, query *utils.PaginationQuery) (*models.PriceGroupSettings, error)

func (u *priceGroupUC) GetAllByNewsID(ctx context.Context, priceGroupID int) (*models.PriceGroupSettings, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "priceGroupUC.GetAllByNewsID")
	defer span.Finish()

	return u.priceGroupRepo.GetAllByNewsID(ctx, priceGroupID)
	//ctx, priceGroupID, query)
}

func (u *priceGroupUC) GetAllPriceGroup(ctx context.Context, query *utils.PaginationQuery) (*models.PriceGroupSettings, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "commentsUC.GetAllByNewsID")
	defer span.Finish()

	return u.priceGroupRepo.GetAllPriceGroup(ctx, query)
	//ctx, priceGroupID, query)
}

func (u *priceGroupUC) Create(ctx context.Context, priceGroup *models.PriceGroupSettings) (*models.PriceGroupSettings, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "priceGroupUC.Create")
	defer span.Finish()

	/*user, err := utils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, httpErrors.NewUnauthorizedError(errors.WithMessage(err, "priceGroupUC.Create.GetUserFromCtx"))
	}

	priceGroup.AuthorID = user.UserID

	if err = utils.ValidateStruct(ctx, priceGroup); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.WithMessage(err, "priceGroupUC.Create.ValidateStruct"))
	}
	*/

	n, err := u.priceGroupRepo.Create(ctx, priceGroup)
	if err != nil {
		return nil, err
	}

	if err = u.redisRepo.DeletePriceGroupCtx(ctx, "price_group_redis"); err != nil {
		fmt.Printf("LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLl")
		u.logger.Errorf("priceGroup.create.DeleteNewsCtx: %v", err)
	}

	return n, err
}

// Update priceGroup item

func (u *priceGroupUC) Update(ctx context.Context, priceGroup *models.PriceGroupSettings) (*models.PriceGroupSettings, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "priceGroupUC.Update")
	defer span.Finish()
	fmt.Println(priceGroup)
	updatedUser, err := u.priceGroupRepo.Update(ctx, priceGroup)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *priceGroupUC) GetAllPriceGroupNew(ctx context.Context, filterQuery string, pq *utils.PaginationQuery) (*models.PriceGroupSettingsList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "price_group.GetAllPriceGroupNew")
	priceGroupBase, err := u.redisRepo.GetPriceGroupCtx(ctx, "price_group_redis")
	if err != nil {
		u.logger.Errorf("priceGroup.GetNewsByID.GetNewsByIDCtx: %v", err)
	}
	if priceGroupBase != nil {
		return priceGroupBase, nil
	}
	defer span.Finish()
	fmt.Println("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	n, err := u.priceGroupRepo.GetAllPriceGroupNew(ctx, filterQuery, pq)
	if err != nil {
		return nil, err
	}

	if err = u.redisRepo.SetPriceGroupCtx(ctx, "price_group_redis", cacheDuration, n); err != nil {
		u.logger.Errorf("newsUC.GetNewsByID.SetNewsCtx: %s", err)
	}

	return u.priceGroupRepo.GetAllPriceGroupNew(ctx, filterQuery, pq)
}

func (u *priceGroupUC) Delete(ctx context.Context, newsID int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "newsUC.Delete")
	defer span.Finish()

	_, err := u.priceGroupRepo.GetAllByNewsID(ctx, newsID)
	if err != nil {
		return err
	}

	if err = u.priceGroupRepo.Delete(ctx, newsID); err != nil {
		return err
	}

	return nil
}
func (u *priceGroupUC) getKeyWithPrefix(newsID string) string {
	return fmt.Sprintf("%s: %s", basePrefix, newsID)
}
