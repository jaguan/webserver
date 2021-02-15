package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func init() {
	fmt.Println("statichandler package initialized")
}

// GetStaticHandler is ...
func GetStaticHandler() StaticHandler {
	handler := new(StaticHandler)
	return *handler
}

// StaticHandler is ...
type StaticHandler struct {
	http.Handler
}

func (h *StaticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	localPath := "wwwroot" + req.URL.Path
	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	contentType := getContentType(localPath)
	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)

	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".png":
		contentType = "image/png"
	case ".jpg":
		contentType = "image/jpeg"
	default:
		contentType = "text/plain"
	}

	return contentType
}
