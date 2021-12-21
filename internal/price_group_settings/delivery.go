package price_group_settings

import "github.com/labstack/echo/v4"

// News HTTP Handlers interface
type Handlers interface {
	GetAllByNewsID() echo.HandlerFunc
	//GetAllPriceGroup() echo.HandlerFunc
	Create() echo.HandlerFunc
	//Update() echo.HandlerFunc
	//GetAllPriceGroupNew() echo.HandlerFunc
	//Delete() echo.HandlerFunc
}
