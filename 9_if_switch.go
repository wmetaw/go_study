package main

import(
    "fmt"
)

// ifとswitch
func main() {

    // ifの中でしか生存しない変数の場合、この書き方ができる
    if _score := 60; _score > 80 {
        fmt.Println("Great")
    } else if _score >  60 {
        fmt.Println("Nice")
    } else {
        fmt.Println("Oh...")
    }

    // switch
    signal := "red"
    switch signal {
    case "red":
        fmt.Println("Stop")
    case "yellow":
        fmt.Println("Caution")
    case "green", "blue":
        fmt.Println("Go!!")
    default:
        fmt.Println("Accident")
    }

    // switchにifを使う場合
    score := 80
    switch {
    case score > 80:
        fmt.Println("Great")
    case score > 60:
        fmt.Println("Nice")
    default:
        fmt.Println("Oh...")
    }

}