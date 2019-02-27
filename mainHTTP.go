package main

import (
	"fmt"
	"net/http"
)

func main() {

	// /healthハンドラ作成
	http.HandleFunc("/health", healthHandler)

	// /sample.htmlハンドラ作成
	http.Handle("/sample.html", http.FileServer(http.Dir("contents")))

	// /profile.htmlハンドラ作成
	http.Handle("/profile.html", http.FileServer(http.Dir("contents")))

	// 8080ポートで起動
	http.ListenAndServe(":8080", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Go.")
}

func add(x int, y int) (int, error) {
	z := x + y
	return z, nil
}
