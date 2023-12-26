package main

import (
	"go-lang-templ/handler"

	"github.com/labstack/echo/v4"
)

type DB struct {}
func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}
	app.Use()
	app.GET("/user", userHandler.HandleUserShow)
	app.Start(":3000")
	// fmt.Println("it is wrok")
}
// https://www.youtube.com/watch?v=wttTTFVrQiw
// 24:49
// func withUser(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Set("user","a@gg.com")
// 		ctx := context.WithValue("user","a@gg.com")
// 		c.SetRequest(c.Request().WithContext(ctx))
// 		return next(c)
// 	}
// }