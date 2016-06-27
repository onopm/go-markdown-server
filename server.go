package mdserver

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func ListenAndServe() {

	e := echo.New()
	rl := initLog()

	// Middleware
	//e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: rl,
	}))
	e.Use(middleware.Recover())

	// handler
	e.GET("/*", getHandler)
	e.POST("/*", postHandler)

	html, err := fileRender("README.md")
	if err != nil {
		return
	}
	mdCache["/"] = html

	fmt.Println("start server")
	e.Run(standard.New(":5000"))
}
