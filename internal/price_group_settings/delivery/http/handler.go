package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AleksK1NG/api-mc/config"
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
// @Summary Get comments by news
// @Description Get all comment by news id
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param id path int true "news_id"
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Success 200 {object} models.CommentsList
// @Failure 500 {object} httpErrors.RestErr
// @Router /comments/byNewsId/{id} [get]
func (h *priceGroupHandlers) GetAllByNewsID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "commentsHandlers.GetAllByNewsID")
		defer span.Finish()
		//	fmt.Println("KKKKKKKK")

		newsID, err := strconv.Atoi(c.Param("news_id"))
		fmt.Println(newsID)

		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		commentsList, err := h.priceGroupUC.GetAllByNewsID(ctx, newsID, pq)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, commentsList)
	}
}

func (h *priceGroupHandlers) GetAllPriceGroup() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "commentsHandlers.GetAllByNewsID")
		defer span.Finish()
		//	fmt.Println("KKKKKKKK")

		//	newsID, err := strconv.Atoi(c.Param("news_id"))
		//fmt.Println(newsID)

		/*if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		*/

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
