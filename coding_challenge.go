package main

import (
	_ "github.com/lib/pq"
	"github.com/nsan1129/unframed"
	"github.com/nsan1129/unframed/log"
)

var net *unframed.NetHandle
var db *unframed.DbHandle

var programTitle = "Coding Challenge"
var cfgFile string = "/etc/coding-challenge/config.json"

func main() {
	cfg := unframed.ReadConfig(cfgFile)

	db = unframed.NewDB(cfg.DbType, cfg.ConnStr)
	defer db.Close()
	log.Message("Working Directory:", cfg.Wd)
	net = unframed.NewNet()

	unframed.DefaultPageTitle = programTitle
	/* Removed because all the crap is in coding_challenge.html.tmpl template
	net.TemplateFiles(
		"tmpl/_base.html.tmpl",
	)
	*/
	net.Dir("assets/")
	net.Dir("public/")

	registerControllers()

	db.PrepareStatements()
	net.LoadTemplates()

	log.Message("Serving", programTitle, "on port:", cfg.ListenPort)
	net.Serve(cfg.ListenPort)

}

func registerControllers() {
	codingChallengeReg()
	downloadReg()
	submitReg()
}
