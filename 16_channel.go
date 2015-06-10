package main

import (
"fmt"
"time"
)


func task1(result chan string) {

    // ２秒間停止
    time.Sleep(time.Second * 2)

    // チャネルに値を送信
    result<- "task1 finish"
}

func task2() {

    fmt.Println("task2 finish")
}

/**
 * Channel を用いたメッセージング
 * Channel は参照型なので make() でインスタンスを生成
 * 送信が "channel<-value"  受信が "<-channel"
 */
 func main() {

    // チャネルのインスタンスを生成
    result := make(chan string)

    go task1(result)
    go task2()

    /**
     * <-result チャネルの値を取り出す
     * resultに値が入るまで処理がブロックされる
     */
     fmt.Println(<-result)


    // 即時関数を使い、擬似ローディングしてみる
     complete := make(chan bool)
     go func() {

        fmt.Printf("Now Loading")
        time.Sleep(time.Second * 1)
        fmt.Printf(".")
        time.Sleep(time.Second * 1)
        fmt.Printf(".")
        time.Sleep(time.Second * 1)
        fmt.Printf(".")

        complete<- true
    }()

    <-complete // 受信した値自体は必要ないため、捨てる

    fmt.Println("\r\nend")
}


