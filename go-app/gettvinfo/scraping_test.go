package gettvinfo_test

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/Makoto87/tv-comment-app/go-app/gettvinfo"
	"golang.org/x/net/html"
)

func TestScraping(t *testing.T) {
	cases := []struct {
		name, html, inputSelector, want string
	}{
		{"test1", "test1.html", "div1", "test1_result.html"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// test用のhttpサーバーを立ち上げる
			f := "testdata/" + c.html
			template := template.Must(template.ParseFiles(f))
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if err := template.ExecuteTemplate(w, c.html, nil); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}))
			defer ts.Close()
			url := ts.URL
			fmt.Println(url)

			// 希望の出力結果を読み込む
			rf, err := os.Open("testdata/" + c.want)
			if err != nil {
				log.Println(err)
			}
			r, err := html.Parse(rf)

			// 立ち上げたサーバーのURLを利用
			got := gettvinfo.Scraping(url, c.inputSelector)
			gh, err := html.Parse(strings.NewReader(got))
			if err != nil {
				log.Println(err)
			}
			
			if reflect.DeepEqual(gh, r) {
				t.Errorf("Want = %v, \nGot = %v", r, gh)
			}
		})
	}
}
