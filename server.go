package mdserver

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/russross/blackfriday"
)

func rootHandler(c echo.Context) error {

	md, _ := ioutil.ReadFile("README.md")
	b := blackfriday.MarkdownCommon([]byte(md))

	return c.HTML(http.StatusOK, string(b))
}

func getHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!!")
}
func postHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!!")
}
func putHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!!")
}
func deleteHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!!")
}

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
	e.GET("/", rootHandler)
	e.GET("/*", getHandler)
	e.POST("/*", postHandler)

	fmt.Println("start server")
	e.Run(standard.New(":5000"))
}
