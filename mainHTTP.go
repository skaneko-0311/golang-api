package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {

	// /healthハンドラ作成
	http.HandleFunc("/health", healthHandler)

	// ルートディレクトリ設定
	http.Handle("/", http.FileServer(http.Dir("contents")))

	// 8080ポートで起動
	http.ListenAndServe(":8080", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("free")
	out, err := cmd.Output()
	fmt.Fprintf(w, "<結果>\n %s", out)
	fmt.Fprintf(w, "<エラー>\n %s", err)
}

func add(x int, y int) (int, error) {
	z := x + y
	return z, nil
}
