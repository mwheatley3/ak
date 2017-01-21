package main

import (
	"github.com/mwheatley3/ak/server/web"
)

var confPath string

// Config is portal service configuration
type Config struct {
	// Web struct {
	// 	web.Config
	// 	Cookie struct {
	// 		BlockKey goconfig.HexBytes
	// 		HashKey  goconfig.HexBytes
	// 	}
	// }
	// Postgres pg.Config
	// Log      log.Config
}

// func loadConfig() Config {
// 	var c Config
// 	config.MustLoad(&c, config.FromFileWithOverride(confPath))
//
// 	c.Web.Config.Cookie.BlockKey = c.Web.Cookie.BlockKey
// 	c.Web.Config.Cookie.HashKey = c.Web.Cookie.HashKey
//
// 	return c
// }
//
// func pgLoadConfig() (pg.Config, log.Config) {
// 	c := loadConfig()
// 	return c.Postgres, c.Log
// }

// srv := web.Sesrver{
//   HTTPServer: http.Server{},
//   Router:     http.NewServeMux(),
// }
// fs := http.StripPrefix("/public/", http.FileServer(http.Dir("../../public")))
// s.Router.Handle("/public/", fs)
// s.Router.HandleFunc("/", react)
// s.Router.HandleFunc("/hello", hello)
// s.Router.HandleFunc("/api/auth", auth)
// s.Router.HandleFunc("/tweet", s.WithTwitterClient(tweets))
