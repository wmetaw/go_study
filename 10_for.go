package main

import(
    "fmt"
)

/**
 * for文
 *
 * ※goにはwhileがないので、ループは基本的にforで
 */
func main() {

    // 基本構文
    for i := 0; i < 3; i++ {
        if i == 0 {
            continue
        }
        fmt.Println(i)
    }

    // whileっぽくする
    i := 0
    for i < 3 {
        fmt.Println(i)
        i++
    }

    // 無限ループ
    n := 0
    for {
        fmt.Println(n)
        if n == 5 {
            break
        }
        n++
    }
}