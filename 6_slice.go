package main

import (
    "fmt"
)

/**
 * スライスは配列の部分列を簡単に取り出すことができるデータ構造
 *
 * Go の配列は固定長、スライスは可変長配列のようなもの
 *
 * a := [3]int{1,2,3}   // 配列
 * a := [...]int{1,2,3} // 配列
 * a := []int{1,2,3}    // スライス
 */
func main() {

    // 要素数をしていせず、配列を作成(コンパイラが数える)
    c := [...]int{1,3,5,7,9}

    // 2番目から(4-1)番目をスライス
    d := c[2:4]

    // [5,7]
    fmt.Println(d)
    fmt.Println(c)

    // スライスし生成した配列に値を代入 ※元の配列の参照なのでc[4]も12となる
    d[1] = 12
    fmt.Println(c)

    // len() 配列の長さ
    // cap() 配列の先頭からきりだせる最大数
    fmt.Println(d, len(d), cap(d))
}

// 参考リンク
// http://jxck.hatenablog.com/entry/golang-slice-internals
// http://www.slideshare.net/yasi_life/go-14075425