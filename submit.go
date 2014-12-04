package main

import (
	//"github.com/nsan1129/unframed/log"
	"github.com/nsan1129/unframed"
	"net/http"
)

func submitReg() {
	sr := net.Subrouter("/submit")
	sr.Post("/verify", submitVerify)
	sr.Get("/", submit)

	net.TemplateFiles(
		"tmpl/submit.html.tmpl",
	)

}

func submit(w http.ResponseWriter, r *http.Request) {
	net.ExeTmpl(w, "submit", nil)
}

type DivisionData struct {
	Paris  int
	London int
	Berlin int
}

type Results map[string]string

type Styler struct {
	awesomeness int
}

func (s Styler) Style(rs string) (style string) {
	if rs == "Correct" {
		style = "color:green"
	} else if rs == "Incorrect" {
		style = "color:red"
	} else {
		style = "color:grey"
	}
	return style
}

func submitVerify(w http.ResponseWriter, r *http.Request) {

	subs := &DivisionData{}
	ans := map[string]int{}
	unframed.File2Json("answers.json", &ans)
	net.DecodeForm(subs, r)

	results := Results{}

	checkAns(subs.Paris, ans, results, "Paris")
	checkAns(subs.London, ans, results, "London")
	checkAns(subs.Berlin, ans, results, "Berlin")

	embVideo := false

	if results["Paris"] == "Correct" && results["London"] == "Correct" && results["Berlin"] == "Correct" {
		embVideo = true
	}

	styler := Styler{}

	net.ExeTmpl(w, "submitVerify", results, styler, embVideo)

}

func checkAns(subm int, ans map[string]int, results Results, city string) {
	if subm == ans[city] {
		results[city] = "Correct"
	} else {
		results[city] = "Incorrect"
	}
}
