package gettvinfo_test

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Makoto87/tv-comment-app/go-app/gettvinfo"
)

func TestScraping(t *testing.T) {
	cases := []struct {
		name          string
		targetHtml    string
		inputSelector string
		want          string
	}{
		{"test1", "test1.html", ".div1 p", "test1,test1,test1"},
		{"test2", "test1.html", ".t1", "test1,test1,test1"},
		{"test3", "test2.html", "li.js1", "test"},
		{"test4", "test1.html", "div2", ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// test用のhttpサーバーを立ち上げる
			f := "testdata/" + c.targetHtml
			template := template.Must(template.ParseFiles(f))
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if err := template.ExecuteTemplate(w, c.targetHtml, nil); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}))
			defer ts.Close()
			url := ts.URL

			// 立ち上げたサーバーのURLを利用
			got := gettvinfo.Scraping(url, c.inputSelector)
			gotStrs := strings.Join(got, ",")

			if gotStrs != c.want {
				t.Errorf("Want = %#v, \nGot = %#v", c.want, gotStrs)
			}
		})
	}
}
