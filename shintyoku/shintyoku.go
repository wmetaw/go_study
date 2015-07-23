package main

import(
    "fmt"
    "math/rand"
    "time"
)

var max int = 0

func main() {

    // 時刻を乱数の種に設定して擬似乱数を発生させる
    rand.Seed(time.Now().Unix())

    // 煽り関数
    dis(0)
}

// 煽り関数
func dis(cnt int) {

    if cnt == 4 {

        fmt.Println(`
 ＿人人人人人人人＿
 ＞進捗どうですか＜
 ￣Y^Y^Y^Y^Y^Y^Y￣`)
    fmt.Printf("%d回煽りました\n", max)

        return
    }

    // 煽った回数
    max++

    rnd := rand.Intn(4)
    switch rnd {
        case 0:
            fmt.Print("進捗")
            cnt = 1
            dis(cnt)

        case 1:
            fmt.Print("どう")
            if rnd == cnt {
                cnt++
            } else {
                cnt = 0
            }
            dis(cnt)

        case 2:
            fmt.Print("です")
            if rnd == cnt {
                cnt++
            } else {
                cnt = 0
            }
            dis(cnt)

        case 3:
            fmt.Print("か")
            if rnd == cnt {
                cnt++
            } else {
                cnt = 0
            }
            dis(cnt)
    }

}