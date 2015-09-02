package main

import (
    "fmt"
    "time"
)

func task1() {

    // ２秒間停止
    time.Sleep(time.Second * 2)
    fmt.Println("task1 finish")
}

func task2() {

    fmt.Println("task2 finish")
}

// ゴルーチン
func main(){

    // 並行処理開始
    go task1()
    go task2()

    // task1が終了する前にmain関数が終了するため３秒まつ
    time.Sleep(time.Second * 3)
}


/**
 * 【Tips】並行と並列の違い
 *
 * 並行: 時分割でスレッドを処理
 * 並列: マルチコアで処理
 *
 * 並行: 複数の動作が、順不同もしくは同時に起こりうる
 * 並列: 複数の動作が、同時に起こること(非同期処理)
 *
 * 並行: 実行状態を複数保てる
 * 並列: 複数の動作を同時に出来る
 *
 *
 * 単純に「行」と「列」で考えるとわかりやすいかも
 *
 * 参考リンク
 * http://www.m-tea.info/2011/03/concurrent-parallel-01.html
 */