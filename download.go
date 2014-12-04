package main

import (
	"net/http"
	//"github.com/nsan1129/unframed/log"
	//"github.com/nsan1129/unframed"
)

func downloadReg() {
	sr := net.Subrouter("/download")
	sr.Get("/", download)

	net.TemplateFiles(
		"tmpl/download.html.tmpl",
	)

}

func download(w http.ResponseWriter, r *http.Request) {
	net.ExeTmpl(w, "download", nil)
}
