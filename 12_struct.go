package main

import(
    "fmt"
)

// 構造体の宣言
type user struct {
    first_name string
    last_name string
    score int
}

/**
 * 構造体
 */
func main() {

    u := new(user) // アドレスが返る

    // 代入方法1
    u.first_name = "Tachi"

    // 代入方法2
    (*u).last_name = "Hiroshi"

    // 代入方法3 フィールド順
    uu  := user {"Tachi", "Hiroshi", 100}

    // 代入方法4 キー指定っぽく
    uuu  := user {first_name:"Tachi", last_name:"Hiroshi", score:100}

    fmt.Println(u)
    fmt.Println(uu)
    fmt.Println(uuu)
}