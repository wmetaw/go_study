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

func main() {

    // 通常の代入
    var msg = "hello world"

    // 短縮代入
    m := "hello world!!"

    fmt.Println(msg)
    fmt.Println(m)

}
