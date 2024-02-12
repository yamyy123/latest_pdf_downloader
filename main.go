package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/yameen0603/echolabstack/routes"
)

func main() {

  e := echo.New()
  routes.NewRoutes(e)
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
  e.Logger.Fatal(e.Start(":8080"))
}

