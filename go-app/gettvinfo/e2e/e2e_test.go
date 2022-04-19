//go:build e2e

package e2e

import (
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Makoto87/tv-comment-app/go-app/gettvinfo"
)

func TestMain(m *testing.M) {
	template := template.Must(template.ParseFiles("../testdata/e2e.html"))
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := template.ExecuteTemplate(w, "e2e.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))
	defer ts.Close()

	url := ts.URL
	const selector = ".epg-item_seriesTitleText__RnbO0"

	programs := gettvinfo.Scraping(url, selector)

	if err := gettvinfo.ProgramInsert(programs); err != nil {
		log.Println("Fatal: ProgramInsert ", err)
	}

	if err := gettvinfo.EpisodeInsert(programs); err != nil {
		log.Println("Fatal: EpisodeInsert ", err)
	}
}
