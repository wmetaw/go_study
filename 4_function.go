package main

import (
    "fmt"
)

// 関数、関数リテラル
func main() {

    // hello world
    a := hello("wolrd")
    fmt.Println(a)


    // 複数戻り値
    b , c := swap(111, 222)
    fmt.Println(b, c)


    // 無名関数 (swap関数を変数に代入)
    tmp := func(a, b int) (int, int) {
        return b , a
    }
    fmt.Println( tmp(44, 33) )


    // 即時関数
    func (msg string) {
        fmt.Println(msg)
    }("関数を定義して実行")

}

/**
 * 引数で受け取った文字列に連結させて返す
 *
 * @return string
 */
//func hello(変数名 型) (戻り値 型)
func hello(msg string) (ret string) {

    ret = "hello " + msg
    return
}

/**
 * 引数を逆にしてまとめて返す
 */
func swap(a, b int) (int, int) {
    return b , a
}




/*
    リテラルとは

    ソースコード内に値となる、文字列、数字、式を直接表記したもの。
    一例として、変数が箱とたとえられるのであれば、その変数の中に入る値をリテラルと呼ぶ。

        例）
        var string = "Hello World";
        var num = 10;

    ここでいう代入された値の"Hello World"や"10"のことをリテラルという。


    関数リテラルとは
    変数に関数を代入して記述することを関数リテラルと言う。

        例1）
        var func01 = function(){処理};

    関数リテラル、無名関数、匿名関数この三つは同義語。

 */



