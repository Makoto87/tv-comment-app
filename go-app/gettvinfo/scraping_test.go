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
		{ name: "test1", targetHtml: "test1.html", inputSelector: ".div1 p", want: "test1,test1,test1" },
		{ name: "test2", targetHtml: "test1.html", inputSelector: ".t1", want: "test1,test1,test1" },
		{ name: "test3", targetHtml: "test2.html", inputSelector: "li.js1", want: "test" },
		{ name: "test4", targetHtml: "test1.html", inputSelector: "div2", want: ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := "testdata/" + c.targetHtml
			template := template.Must(template.ParseFiles(f))
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if err := template.ExecuteTemplate(w, c.targetHtml, nil); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}))
			defer ts.Close()
			url := ts.URL

			got := gettvinfo.Scraping(url, c.inputSelector)
			gotStrs := strings.Join(got, ",")

			if gotStrs != c.want {
				t.Errorf("Want = %#v, \nGot = %#v", c.want, gotStrs)
			}
		})
	}
}
