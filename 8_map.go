package main

import (
    "fmt"
)

// MAP (連想配列のようなもの)
func main() {

    // map[キーの型]値の型
    m := make(map[string]string)
    m["first_name"] = "yamada"
    m["last_name"]  = "taro"

    fmt.Println(m);

    // 宣言と代入
    i := map[string]string{"first_name":"yamada", "last_name":"taro"}
    fmt.Println(i);

    // キーを指定し、要素を削除
    delete(i, "first_name")
    fmt.Println(i);

    // 値が存在するか
    value, flg := i["yamada"]
    fmt.Println(value, flg);
}