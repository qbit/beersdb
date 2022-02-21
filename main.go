package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"suah.dev/beersdb/db"
)

var (
	pd, err     = sql.Open("postgres", "host=localhost dbname=qbit sslmode=disable")
	ctx, cancel = context.WithCancel(context.Background())
	base        = db.New(pd)
	nounRe      = regexp.MustCompile(`^\/api\/(.+)\/.+$`)
	verbRe      = regexp.MustCompile(`^\/api\/.+\/(.+)$`)
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
	err = r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	noun := nounRe.ReplaceAllString(r.URL.Path, "$1")
	verb := verbRe.ReplaceAllString(r.URL.Path, "$1")

	log.Println(noun, verb)

	if noun == "user" && verb == "login" {
		tokenData, err := base.Login(ctx, db.LoginParams{
			Username: r.Form.Get("username"),
			Crypt:    r.Form.Get("password"),
		})
		if err != nil {
			a.Error(w, http.StatusUnauthorized)
			log.Println(err)
			return
		}

		json.NewEncoder(w).Encode(tokenData)
		return
	}

	if token == "" {
		log.Println("Here")
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
		case "add":
			var b db.CreateBeerParams
			json.NewDecoder(r.Body).Decode(&b)

			beer, err := base.CreateBeer(ctx, b)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println(err)
				return
			}

			json.NewEncoder(w).Encode(beer)
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
		switch verb {
		case "add":
			var b db.CreateBreweryParams
			json.NewDecoder(r.Body).Decode(&b)

			brewery, err := base.CreateBrewery(ctx, b)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println(err)
				return
			}

			json.NewEncoder(w).Encode(brewery)
		}
	case "type":
		switch verb {
		case "add":
			var t db.CreateTypeParams
			json.NewDecoder(r.Body).Decode(&t)

			bType, err := base.CreateType(ctx, t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println(err)
				return
			}

			json.NewEncoder(w).Encode(bType)
		}
	case "user":
		switch verb {
		case "login":
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
