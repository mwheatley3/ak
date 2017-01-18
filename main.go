package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/mwheatley3/ak/server"
	"github.com/mwheatley3/ak/server/db"
	"github.com/mwheatley3/ak/server/pg"
	"strconv"
	// "github.com/jmoiron/sqlx"
	"github.com/mwheatley3/ak/server/twitter"
	"net/http"
	"os"
)

func main() {
	s := server.Server{
		HTTPServer: http.Server{},
		Router:     http.NewServeMux(),
	}
	fs := http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
	s.Router.Handle("/public/", fs)
	s.Router.HandleFunc("/", react)
	s.Router.HandleFunc("/hello", hello)
	s.Router.HandleFunc("/api/auth", auth)
	s.Router.HandleFunc("/tweet", s.WithTwitterClient(tweets))
	// s.Router.HandleFunc("/access", twitterAccess)

	twitterKey := os.Getenv("TWITTER_KEY")
	twitterSecret := os.Getenv("TWITTER_SECRET")

	l := logrus.New()
	twitterClient := twitter.New(twitter.BaseURL, twitterKey, twitterSecret, l)
	s.TwitterClient = twitterClient

	database := os.Getenv("WW_DB")
	dbPort := os.Getenv("WW_DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dp, _ := strconv.Atoi(dbPort)
	dbHost := os.Getenv("WW_DB_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}

	db := db.NewFromConfig(l, pg.Config{
		Host:     dbHost,
		Port:     uint16(dp),
		Database: database,
		// Password:       "abc",
		User:           "mwheatley",
		SslMode:        "prefer",
		SlowThreshold:  40,
		MaxConnections: 30,
	})
	err := db.Init()
	if err != nil {
		fmt.Printf(err.Error())
	}

	a, err := db.Test("")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("return from postgres%#+v\n", a)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	fmt.Println("listening... port" + port)
	s.HTTPServer.Addr = ":" + port
	s.HTTPServer.Handler = s.Router
	err = s.HTTPServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world")
}

func auth(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "AUTH API")
	println("***************************")
}

func react(res http.ResponseWriter, req *http.Request) {
	println("react")
	http.ServeFile(res, req, "./index.html")
}

func tweets(res http.ResponseWriter, req *http.Request) {
	// is storing the twitter client on context correct?
	twitterClient := req.Context().Value("twitterClient").(*twitter.Client)
	if twitterClient.AccessToken == "" {
		twitterClient.Access()
	}
	url := fmt.Sprintf("%s/statuses/user_timeline.json?screen_name=ItsFlo&count=2", twitterClient.BaseURL)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf(err.Error())
	}

	// var dest map[string]interface{}
	var dest []twitter.Tweet
	resp, err := twitterClient.Call(request, &dest)

	fmt.Printf("dest%#+v\n", dest)
	if err != nil {
		fmt.Printf(err.Error())
	}

	println("inside twitter handler\n")
	fmt.Printf("%#+v request context\n", req.Context().Value("twitterClient"))
	fmt.Printf("resp %#+v\n", resp)

	// twitterClient
	twitter.Feed()
}
