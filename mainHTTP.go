package main

import (
	"net/http"
)

func main() {

	// 8080ポートで起動
	http.ListenAndServe(":8080", nil)
}
