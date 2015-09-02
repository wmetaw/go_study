package main

import (
    "fmt"
)


// 構造体の宣言
type user struct {
    first_name string
    last_name string
    score int
}


/**
 * メソッド
 *
 * 【関数とメソッドの違い】
 * オブジェクト自身を操作する場合に使う手続きをメソッド。
 * それ以外を関数（非オブジェクト指向言語）
 * http://deepfolte.hatenablog.com/entry/20080513/1210662275
 */
func main() {

    u := user{first_name:"Tachi", last_name:"Hiroshi", score:79}
    u.clear()
    u.show()
    fmt.Println(u)
}


// メソッド
func (u *user)clear(){
    u.score++
}

func (u *user)show(){
    fmt.Printf("Name:%s %s Score:%d \r\n", u.first_name, u.last_name, u.score)
}