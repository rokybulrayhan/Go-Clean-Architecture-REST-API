package http

import (
	"github.com/AleksK1NG/api-mc/internal/middleware"
	price_group "github.com/AleksK1NG/api-mc/internal/price_group_settings"
	"github.com/labstack/echo/v4"
)

// Map news routes

func MapPriceGroupRoutes(priceGroup *echo.Group, h price_group.Handlers, mw *middleware.MiddlewareManager) {

	priceGroup.GET("/group/:news_id", h.GetAllByNewsID())
	priceGroup.GET("/group", h.GetAllPriceGroup())
}
