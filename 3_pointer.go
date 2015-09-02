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

// ポインタ操作
func main (){

    a := 5

    // int 型のアドレスを宣言
    var pa * int

    pa = &a

    // 出力
    fmt.Println(*pa)
    fmt.Println(pa)

}