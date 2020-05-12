package blog

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func staticFileServer(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v", r.URL.Path)
	webdir := os.Getenv("SFS_WEB_DIR")
	if !isFile(webdir + p(r.URL.Path)) {
		err404, _ := ioutil.ReadFile(webdir + "/assets/error/404.html")
		w.WriteHeader(404)
		w.Write(err404)
		return
	}
	http.ServeFile(w, r, webdir+p(r.URL.Path))
}

func p(urlpath string) string {
	if urlpath == "/" {
		return "/index.html"
	}
	return urlpath
}

func isFile(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}
