package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// "/"の場合の処理
	router.Path("/").HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		// テンプレート
		tmpl := template.Must(template.New("index").ParseFiles("templates/index.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			panic(err)
		}

		// logの代わり
		fmt.Print(time.Now())
		fmt.Printf(" path: " + r.URL.Path + "\n")
	}).Methods("GET")

	// "/save"の場合の処理
	r.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {

		// formの値を取得
		r.ParseForm()
		data1 := r.FormValue("aaa")
		data2 := r.FormValue("bbb")

		// テンプレート
		tmpl := template.Must(template.New("save").ParseFiles("templates/save.html"))
		err := tmpl.Execute(w, struct {
			Data1 string
			Data2 string
		}{
			Data1: data1,
			Data2: data2,
		})
		if err != nil {
			panic(err)
		}

		// logの代わり
		fmt.Print(time.Now())
		fmt.Printf(" path: " + r.URL.Path + "\n")
	}).Methods("POST")

	// 可変URL1
	r.HandleFunc("/test1/{sample1}/", func(w http.ResponseWriter, r *http.Request) {
		// 画面サンプル
		fmt.Fprintf(w, "sample1\n")

		// logの代わり
		fmt.Print(time.Now())
		fmt.Printf(" path: " + r.URL.Path + "\n")
	})

	// 可変URL2
	r.HandleFunc("/test2/{sample1}/{id:[0-9]{1,6}}/", func(w http.ResponseWriter, r *http.Request) {
		// 画面サンプル
		fmt.Fprintf(w, "sample2\n")

		// logの代わり
		fmt.Print(time.Now())
		fmt.Printf(" path: " + r.URL.Path + "\n")
	})

	// handing
	http.Handle("/", r)

	// 設定
	http.ListenAndServe(":9001", nil)
}
