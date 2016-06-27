package mdserver

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"

	"github.com/russross/blackfriday"
)

func fileRender(filename string) (string, error) {

	md, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	b := blackfriday.MarkdownCommon([]byte(md))
	return string(b), nil
}

func readerRender(r multipart.File) (string, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	b := blackfriday.MarkdownCommon([]byte(body))
	return string(b), nil
}
