package web

import (
	"github.com/annakozyreva1/banner_show/log"
	"github.com/annakozyreva1/banner_show/selector"
	"net/http"
	"strings"
)

var (
	logger = log.Logger
)

func getShowHandler(sel *selector.Selector) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		categories := r.FormValue("categories")
		url, ok := sel.GetBanner(strings.Split(categories, ","))
		if ok {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("<a href=" + url + "/>"))
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func Run(address string, sel *selector.Selector) {
	http.HandleFunc("/", getShowHandler(sel))
	if err := http.ListenAndServe(address, nil); err != nil {
		logger.Fatalf("failed web server: %v", err.Error())
	}
}
