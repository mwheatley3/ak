package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	// needed to drive connection to DB
	_ "github.com/jackc/pgx/stdlib"
	"github.com/mwheatley3/ak/server/db"
	"github.com/mwheatley3/ak/server/pg"
	"github.com/mwheatley3/ak/server/twitter"
	"github.com/mwheatley3/ak/server/web"
	// "html/template"
	"net/http"
	"os"
	"strconv"
)

func main() {
	s := web.Server{
		HTTPServer: http.Server{},
		Router:     httprouter.New(),
	}

	s.Router.GET("/", react)
	s.Router.GET("/login", react)
	s.Router.GET("/sentiments", react)
	s.Router.GET("/keri", react)
	s.Router.GET("/hello", hello)
	s.Router.GET("/public/*splat", fsHandler("/public/", "public"))
	s.Router.GET("/api/users/:userID", s.GetUser)
	s.Router.GET("/api/auth", auth)
	s.Router.GET("/tweet", s.WithTwitterClient(tweets))
	// s.Router.GET("/access", twitterAccess)

	twitterKey := os.Getenv("TWITTER_KEY")
	twitterSecret := os.Getenv("TWITTER_SECRET")

	l := logrus.New()
	twitterClient := twitter.New(twitter.BaseURL, twitterKey, twitterSecret, l)
	s.TwitterClient = twitterClient

	database := os.Getenv("WW_DB")
	if database == "" {
		database = "workingwheatleys"
	}
	dbPort := os.Getenv("WW_DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dp, _ := strconv.Atoi(dbPort)
	dbHost := os.Getenv("WW_DB_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	dbPass := os.Getenv("WW_DB_PASS")
	dbUser := os.Getenv("WW_DB_USER")
	if dbUser == "" {
		dbUser = "mwheatley"
	}

	db := db.NewFromConfig(l, pg.Config{
		Host:           dbHost,
		Port:           uint16(dp),
		Database:       database,
		Password:       dbPass,
		User:           dbUser,
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
	err = db.CreateUsersTable()
	if err != nil {
		fmt.Printf(err.Error())
	}

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

func auth(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// fmt.Fprintln(res, "AUTH API")
	resp := web.JSON(nil, res)
	u := web.User{}
	resp.Success(u)
	// resp.Error(errors.New("No user found"), http.StatusNotFound)
	println("***************************")
}

func hello(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintln(res, "hello, world")
}

func react(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	println("react")
	// if err := indexTmpl.Execute(res, ""); err != nil {
	// 	fmt.Print(err.Error())
	// }
	http.ServeFile(res, req, "./index.html")
}

// var indexTmpl = template.Must(template.New("").Parse(`
// <html>
// 	<head>
// 		<title>Working Wheatleys</title>
// 	</head>
// 	<body>
// 		<div id="root"></div>
// 		<script src='/public/main.js'></script>
// 	</body>
// </html>
// `))

func tweets(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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

func fsHandler(prefix, path string) httprouter.Handle {
	fs := http.StripPrefix(prefix, http.FileServer(http.Dir(path)))

	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fs.ServeHTTP(w, r)
	})
}
