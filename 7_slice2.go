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

    // 0初期化された配列[3]を作成し、それを参照するeスライスを作成
    e := []int{1,3,5}

    // スライスした配列に要素を追加(可変長配列のスライスのみ可能)
    e = append(e, 7, 9)
    fmt.Println(e)

    // eの要素数の配列を作成
    f := make([]int, len(e))

    // eスライスをコピーし、コピーした要素数をgに代入
    g := copy(f,e)
    fmt.Println(f)
    fmt.Println(g)
}

// 参考リンク
// http://jxck.hatenablog.com/entry/golang-slice-internals
// http://www.slideshare.net/yasi_life/go-14075425