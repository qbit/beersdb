package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	_ "github.com/lib/pq"
	"suah.dev/beersdb/db"
)

var (
	pd, err     = sql.Open("postgres", "host=localhost dbname=qbit sslmode=disable password=''")
	ctx, cancel = context.WithCancel(context.Background())
	base        = db.New(pd)
	ere         = regexp.MustCompile(`^\/api\/(.+)\/.+$`)
)

type apiHandler struct {
	handlerMap map[string]interface{}
}

func (apiHandler) Error(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (a apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//if r.Header.Get("Content-Type") != "application/json" {
	//	a.Error(w, http.StatusBadRequest)
	//	return
	//}

	token := r.Header.Get("X-Access-Token")
	part := ere.ReplaceAllString(r.URL.Path, "$1")

	if token == "" {
		a.Error(w, http.StatusUnauthorized)
		return
	}

	user, err := base.GetUserByToken(ctx, token)
	if err != nil {
		a.Error(w, http.StatusInternalServerError)
		return
	}

	if time.Now().After(user.TokenExpires) {
		a.Error(w, http.StatusUnauthorized)
		return
	}

	switch part {
	case "beer":
	case "brewery":
	case "type":
	case "user":
	default:
		a.Error(w, http.StatusBadRequest)
		return
	}
}

func main() {
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}

	defer pd.Close()

	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "BeersDB, a db for all your beer needs!")
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
