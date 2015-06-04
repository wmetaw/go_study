package main

import(
    "fmt"
)

/**
 * range
 * 「Array」「Slice」「map」などをイテレートするrange
 */
func main() {

    // スライスを作成
    slice := []int{2, 3, 8}

    // sliceの要素文ループ  i:インデックス  v:値
    for i , v := range slice {

        // rangeは二つの値を返す
        fmt.Println(i ,v)
    }


    // 値を破棄したければ、ブランク修飾子「 _ 」を使用
    for _ , v := range slice {
        fmt.Println(v)
    }


    // mapの場合は key value が返る
    m := map[string]string{"first_name":"Tachi", "last_name":"Hiroshi"}
    for k ,v := range m {
        fmt.Println(k ,v)
    }
}