package main

import (
	"context"
	"database/sql"
	"encoding/json"
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
	nounRe      = regexp.MustCompile(`^\/api\/(.+)\/.+$`)
	verbRe      = regexp.MustCompile(`&\/api\/.+/(.+)$`)
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
	noun := nounRe.ReplaceAllString(r.URL.Path, "$1")
	verb := verbRe.ReplaceAllString(r.URL.Path, "$1")

	if token == "" {
		a.Error(w, http.StatusUnauthorized)
		return
	}

	user, err := base.GetUserByToken(ctx, token)
	if err != nil {
		a.Error(w, http.StatusUnauthorized)
		return
	}

	if time.Now().After(user.TokenExpires) {
		a.Error(w, http.StatusUnauthorized)
		return
	}

	switch noun {
	case "beer":
		switch verb {
		case "all":
			beers, err := base.GetAllBeers(ctx)
			if err != nil {
				a.Error(w, http.StatusInternalServerError)
				log.Println(err)
				return
			}

			json.NewEncoder(w).Encode(beers)
		default:
			beers, err := base.SearchBeers(ctx, noun)
			if err != nil {
				a.Error(w, http.StatusInternalServerError)
				log.Println(err)
				return
			}
			json.NewEncoder(w).Encode(beers)
		}
	case "brewery":
	case "type":
	case "user":
		switch verb {
		case "regenToken":
			newToken, err := base.GenerateNewToken(ctx, token)
			if err != nil {
				a.Error(w, http.StatusInternalServerError)
				log.Println(err)
				return
			}

			json.NewEncoder(w).Encode(newToken)
		}

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

	log.Fatal(http.ListenAndServe(":8336", mux))
}
