package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/models"
	price_group "github.com/AleksK1NG/api-mc/internal/price_group_settings"
	"github.com/AleksK1NG/api-mc/pkg/httpErrors"
	"github.com/AleksK1NG/api-mc/pkg/logger"
	"github.com/AleksK1NG/api-mc/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

type priceGroupHandlers struct {
	cfg          *config.Config
	priceGroupUC price_group.UseCase
	logger       logger.Logger
}

func NewPriceGroupHandlers(cfg *config.Config, priceGroupUC price_group.UseCase, logger logger.Logger) price_group.Handlers {
	return &priceGroupHandlers{cfg: cfg, priceGroupUC: priceGroupUC, logger: logger}
}

// GetAllByNewsID
// @Summary Get comments by priceGroup
// @Description Get all comment by priceGroup id
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param id path int true "priceGroup_id"
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Success 200 {object} models.CommentsList
// @Failure 500 {object} httpErrors.RestErr
// @Router /comments/byNewsId/{id} [get]
func (h *priceGroupHandlers) GetAllByNewsID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "priceGroupHandlers.GetAllByNewsID")
		defer span.Finish()
		//	fmt.Println("KKKKKKKK")

		priceGroupID, err := strconv.Atoi(c.Param("priceGroup_id"))
		fmt.Println(priceGroupID)

		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		commentsList, err := h.priceGroupUC.GetAllByNewsID(ctx, priceGroupID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, commentsList)
	}
}

/*
func (h *priceGroupHandlers) GetAllPriceGroup() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "priceGroupHandlers.GetAllByNewsID")
		defer span.Finish()
		//	fmt.Println("KKKKKKKK")

		//	priceGroupID, err := strconv.Atoi(c.Param("priceGroup_id"))
		//fmt.Println(priceGroupID)

		/*if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
*/
/*
		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		commentsList, err := h.priceGroupUC.GetAllPriceGroup(ctx, pq)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, commentsList)
	}
}
*/

func (h priceGroupHandlers) Create() echo.HandlerFunc {
	//fmt.Println("jjjjjjjjjjjjjjjjjjjjjjjj")
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "priceGroupHandlers.Create")
		defer span.Finish()

		n := &models.PriceGroupSettings{}
		if err := c.Bind(n); err != nil {

			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		priceGroupCreate, err := h.priceGroupUC.Create(ctx, n)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, priceGroupCreate)
	}
}

/*
func (h priceGroupHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "priceGroupHandlers.Update")
		defer span.Finish()

		priceGroupId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		n := &models.PriceGroupSettings{}
		if err = c.Bind(n); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		n.ID = priceGroupId
		fmt.Println(*n)

		updatedPriceGroup, err := h.priceGroupUC.Update(ctx, n)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, updatedPriceGroup)
	}
}
*/
/*
func (h priceGroupHandlers) GetAllPriceGroupNew() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "priceGroupHandlers.GetPriceGroup")
		defer span.Finish()

		fmt.Println(len(c.Request().URL.Query()))
		var filterQuery string
		var where string = "where deleted_at IS NULL "
		descriptionParam := c.QueryParam("description")
		nameParam := c.QueryParam("name")
		activeParam := c.QueryParam("active")
		if len(c.Request().URL.Query()) != 0 {

			if descriptionParam != "" {
				filterQuery = filterQuery + " AND " + "description like '%" + c.QueryParam("description") + "%'"

			}
			if nameParam != "" {
				filterQuery = filterQuery + " AND " + "name = '" + c.QueryParam("name") + "'"

			}
			if activeParam != "" {
				filterQuery = filterQuery + " AND " + "active = " + c.QueryParam("active")

			}

		}
		filterQuery = where + filterQuery
		fmt.Println(filterQuery)
		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		priceGroupList, err := h.priceGroupUC.GetAllPriceGroupNew(ctx, filterQuery, pq)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, priceGroupList)
	}
}
*/

/*
func (h priceGroupHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "newsHandlers.Delete")
		defer span.Finish()
		fmt.Println("123567")
		id, err := strconv.Atoi(c.Param("id"))
		fmt.Println(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err = h.priceGroupUC.Delete(ctx, id); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.NoContent(http.StatusOK)
	}
}
*/
