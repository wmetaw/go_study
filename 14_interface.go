package main

import (
    "fmt"
)

// インターフェース
type greeter interface {
    greet()
}

// JPN構造体とメソッド
type jpn struct {}
func (j jpn) greet() {
    fmt.Println("こんにちは！")
}

// USA構造体とメソッド
type usa struct {}
func (u usa) greet() {
    fmt.Println("Hello!")
}

/**
 * インターフェース
 * http://jxck.hatenablog.com/entry/20130325/1364251563
 */
func main() {

    // greeter型でスライスを作成
    greeters := []greeter{ jpn{}, usa{} }

    for _, v := range greeters {
        v.greet()

        // 型チェックし、文字列を出力
        checkInterface(v)
    }


}

/*
 * 型チェック
 * t interface{} は空のインターフェースであり、全ての型を受け取ることができる
 */
func checkInterface(t interface{} ) {

    /**
     * 型アサーション
     * 値, フラグ := t.(type) // t が type を満たすかを調べる
     */
    _, flg := t.(jpn)
    if flg {
        fmt.Printf("I am ")
    } else {
        fmt.Printf("I'm ")
    }

    // 型switch
    switch t.(type) {
    case jpn:
        fmt.Println("japanese")
    case usa:
        fmt.Println("american")
    default:
        fmt.Println("human")
    }
}






