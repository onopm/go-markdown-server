package mdserver

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func getHandler(c echo.Context) error {

	b, ok := mdCache[c.Request().URI()]
	if ok {
		return c.HTML(http.StatusOK, string(b))
	}
	return c.String(http.StatusOK, "Hello, World!!")
}

func postHandler(c echo.Context) error {

	m, _ := c.Request().MultipartForm()

	for _, fh := range m.File {
		r, err := fh[0].Open()
		if err != nil {
			continue
		}

		html, err := readerRender(r)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusOK, err.Error())

		}

		//fmt.Printf("cached [%s] [%s]\n", c.Request().URI(), name)
		mdCache[c.Request().URI()] = html
		return c.String(http.StatusOK, "Hello, World!!")
	}

	return c.String(http.StatusOK, "Hello, World!!")
}

func putHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!!")
}

func deleteHandler(c echo.Context) error {

	// delete cache
	delete(mdCache, c.Request().URI())

	return c.String(http.StatusOK, "Hello, World!!")
}
