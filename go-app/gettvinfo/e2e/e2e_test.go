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
	// HTTPサーバー立ち上げ
	template := template.Must(template.ParseFiles("../testdata/e2e.html"))
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := template.ExecuteTemplate(w, "e2e.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))
	defer ts.Close()

	// URLとセレクタ
	url := ts.URL
	const selector = ".epg-item_seriesTitleText__RnbO0"

	// テスト(main関数と同じ動きをする)
	// スクレイピング
	// paragraph、test1、test3、test5、test8 が取得される
	programs := gettvinfo.Scraping(url, selector)

	// 番組名と放送回をinsert
	if err := gettvinfo.ProgramInsert(programs); err != nil {
		log.Println("Fatal: ProgramInsert ", err)
	}

	if err := gettvinfo.EpisodeInsert(programs); err != nil {
		log.Println("Fatal: EpisodeInsert ", err)
	}
}
