package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/AleksK1NG/api-mc/internal/price_group_settings"
	"github.com/AleksK1NG/api-mc/pkg/utils"
	"github.com/opentracing/opentracing-go"
)

// Price Group Repository
type priceGroupRepo struct {
	db *sqlx.DB
}

// Price Group repository constructor
func NewPriceGroupRepository(db *sqlx.DB) price_group_settings.Repository {
	return &priceGroupRepo{db: db}
}

func (r *priceGroupRepo) Create(ctx context.Context, priceGroup *models.PriceGroupSettings) (*models.PriceGroupSettings, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "priceGroupRepo.Create")
	defer span.Finish()

	var n models.PriceGroupSettings
	if err := r.db.QueryRowxContext(
		ctx,
		createPriceGroup,
		&priceGroup.Name,
		&priceGroup.Description,
		&priceGroup.CurrencyType,
		&priceGroup.Active,
		&priceGroup.CreatedBy,
	).StructScan(&n); err != nil {
		return nil, errors.Wrap(err, "priceGroupRepo.Create.QueryRowxContext")
	}

	return &n, nil
}

func (r *priceGroupRepo) GetAllByNewsID(ctx context.Context, priceGroupID int) (*models.PriceGroupSettings, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "priceGroupRepo.GetNewsByID")
	defer span.Finish()

	n := &models.PriceGroupSettings{}
	if err := r.db.GetContext(ctx, n, getNewsByID, priceGroupID); err != nil {
		return nil, errors.Wrap(err, "priceGroupRepo.GetNewsByID.GetContext")
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

/*
span, ctx := opentracing.StartSpanFromContext(ctx, "priceGroupRepo.GetNews")
	defer span.Finish()

	/*var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getTotalCount); err != nil {
		return nil, errors.Wrap(err, "priceGroupRepo.GetNews.GetContext.totalCount")
	}
*/
/*
	var priceGroupList = make([]*models.PriceGroupSettings, 0, 2)
	rows, err := r.db.QueryxContext(ctx, getAllPriceGroup)
	if err != nil {
		return nil, errors.Wrap(err, "priceGroupRepo.GetNews.QueryxContext")
	}
	defer rows.Close()

	for rows.Next() {
		n := &models.PriceGroupSettings{}
		if err = rows.StructScan(n); err != nil {
			return nil, errors.Wrap(err, "priceGroupRepo.GetNews.StructScan")
		}
		priceGroupList = append(priceGroupList, n)
	}

	priceGroupLis := priceGroupList
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "priceGroupRepo.GetNews.rows.Err")
	}

	return priceGroupLis, nil
*/

func (r *priceGroupRepo) Update(ctx context.Context, priceGroup *models.PriceGroupSettings) (*models.PriceGroupSettings, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "priceGroupRepo.Update")
	defer span.Finish()

	var n models.PriceGroupSettings
	if err := r.db.QueryRowxContext(
		ctx,
		updatePriceGroup,
		&priceGroup.Name,
		&priceGroup.Description,
		&priceGroup.CurrencyType,
		&priceGroup.Active,
		&priceGroup.UpdatedBy,
		&priceGroup.ID,
	).StructScan(&n); err != nil {
		return nil, errors.Wrap(err, "priceGroupRepo.Update.QueryRowxContext")
	}

	return &n, nil
}

func (r *priceGroupRepo) GetAllPriceGroupNew(ctx context.Context, filterQuery string, pq *utils.PaginationQuery) (*models.PriceGroupSettingsList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "priceGroupRepo.GetNews")
	defer span.Finish()

	//var getTotalCountDummy string

	//getTotalCountDummy = ""
	fmt.Println(ctx)
	fmt.Println(filterQuery)
	getTotalCountDummy := "SELECT COUNT(id) FROM price_group_settings " + filterQuery

	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getTotalCountDummy); err != nil {
		return nil, errors.Wrap(err, "priceGroupRepo.GetNews.GetAllPriceGroupNew.totalCount")
	}

	if totalCount == 0 {
		return &models.PriceGroupSettingsList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
			Page:       pq.GetPage(),
			Size:       pq.GetSize(),
			HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
			PriceGroup: make([]*models.PriceGroupSettings, 0),
		}, nil
	}

	//getPriceGroup
	getPriceGroupDummy := "SELECT * FROM price_group_settings " + filterQuery + " ORDER BY name, updated_at OFFSET $1 LIMIT $2"
	var priceGroupList = make([]*models.PriceGroupSettings, 0, pq.GetSize())
	rows, err := r.db.QueryxContext(ctx, getPriceGroupDummy, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetNews.QueryxContext")
	}
	defer rows.Close()

	for rows.Next() {
		n := &models.PriceGroupSettings{}
		if err = rows.StructScan(n); err != nil {
			return nil, errors.Wrap(err, "newsRepo.GetNews.StructScan")
		}
		priceGroupList = append(priceGroupList, n)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetNews.rows.Err")
	}
	//fmt.Printf("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	return &models.PriceGroupSettingsList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
		PriceGroup: priceGroupList,
	}, nil
}

func (r *priceGroupRepo) Delete(ctx context.Context, id int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "newsRepo.Delete")
	defer span.Finish()

	result, err := r.db.ExecContext(ctx, deletePriceGroup, id)
	if err != nil {
		return errors.Wrap(err, "newsRepo.Delete.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "newsRepo.Delete.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "newsRepo.Delete.rowsAffected")
	}

	return nil
}
