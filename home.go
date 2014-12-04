package main

import (
	"net/http"
	//"github.com/nsan1129/unframed/log"
	//"github.com/nsan1129/unframed"
)

func codingChallengeReg() {
	net.Get("/", codingChallenge)
	net.Get("/help", cc_help)

	net.TemplateFiles(
		"tmpl/coding_challenge.html.tmpl",
	)
}

func codingChallenge(w http.ResponseWriter, r *http.Request) {
	net.ExeTmpl(w, "codingChallenge", nil)
}

func cc_help(w http.ResponseWriter, r *http.Request) {
	net.ExeTmpl(w, "help_cc", nil)
}
