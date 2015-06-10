package main

import(
    "fmt"
    "net/http"
)

/**
 * webサーバーを作成
 *
 * リクエストを受け取り、表示
 */
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

    // 1文字目から最後まで(0文字目は/がはいる)
    fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

/**
 * 実行後、ローカルにアクセス
 * http://localhost:8080/太刀ひろし
 *
 *
 * yosemiteの場合は下記リンクを参照
 * http://gijutsubu.backstage.jp/?eid=71
 */