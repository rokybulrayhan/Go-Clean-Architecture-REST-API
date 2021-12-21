package http

import (
	"github.com/AleksK1NG/api-mc/internal/middleware"
	price_group "github.com/AleksK1NG/api-mc/internal/price_group_settings"
	"github.com/labstack/echo/v4"
)

// Map priceGroup routes

func MapPriceGroupRoutes(priceGroup *echo.Group, h price_group.Handlers, mw *middleware.MiddlewareManager) {

	priceGroup.GET("/group/:priceGroup_id", h.GetAllByNewsID())
	//priceGroup.GET("/", h.GetAllPriceGroup())
	//priceGroup.GET("/group", h.GetAllPriceGroupNew())
	priceGroup.POST("/group/create", h.Create())
	//priceGroup.PUT("/group/update/:id", h.Update())
	//	priceGroup.DELETE("/group/delete/:id", h.Delete())
}
