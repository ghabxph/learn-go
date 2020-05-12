package blog

import (
	"net/http"
)

func Router() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", staticFileServer)
	return r
}
