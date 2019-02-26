package main

import (
	"fmt"
	"net/http"
)

func main() {

	// /healthハンドラ作成
	http.HandleFunc("/health", healthHandler)

	// 8080ポートで起動
	http.ListenAndServe(":8080", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Go.")
}
