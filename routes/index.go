package routes

import(
	"github.com/labstack/echo/v4"
   "github.com/yameen0603/echolabstack/service"
)

func NewRoutes(e *echo.Echo){
	e.GET("/files", service.Filehandler)
}