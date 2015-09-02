/**
 * [基本的なデータ型]
 * string    "hello"
 * int       1234
 * float64   3.141592
 * bool      false true
 * nil       不定
 *
 * [宣言なしの場合]場合
 * var s string // ""
 * var a int    // 0
 * var f bool   // false
 */

package main

import (
    "fmt"
)

// 基本的な変数
func main() {

    a := 5           // int
    b := 1.35        // float
    c := "hoge"      // 文字列
    var d bool       // bool
    const e = "定数"  // 定数

    fmt.Printf("\n a:%d  b:%f  c:%s  d:%t  e:%s \n\n", a,b,c,d,e)

    // 定数を列挙
    // 識別子 iota を使用することで0から始まる連番にできる
    // 識別子 iota+1 を使用すること1から始まる連番にできる
    const (
        sun = iota+1
        mon
        tue
    )
    fmt.Println(sun, mon, tue)

}
