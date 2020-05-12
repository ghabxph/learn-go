package main

import (
	"github.com/ghabxph/gabriel/blog-server/internal/app/blog"
	"log"
	"net/http"
	"os"
)

func main() {
	if os.Getenv("SFS_WEB_DIR") == "" {
		log.Printf("SFS_WEB_DIR env not set.")
		return
	}
	http.ListenAndServe(":8888", blog.Router())
}
