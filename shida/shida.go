package main

// imageパッケージは画像データを簡単に扱えるようにしたパッケージ。
// Go1.2の時点ではGIF、JPEG、PNGなどを簡単に読み込み・書き出しが可能
import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "log"
    "math/rand"
    "os"
    "time"
)

// 描画ロジック用変数
var (
    N  int =20
    xm float64 = 0.0
    ym float64 = 0.5
    h  float64 = 0.6
)

// 画像用変数
var (
    width  int = 500
    height int = 500
    filename string = "shida.png"
)

// 色設定
var (

    // 白背景
    bgcolor color.Color = color.RGBA{255, 255, 255, 255}

    // 葉の色
    linecolor color.Color = color.RGBA{0, 128, 0, 255}
)

//======================================
// 謎の計算式
//======================================
func W1x(x, y float64) float64 {
    return 0.836*x + 0.044*y
}

func W1y(x, y float64) float64 {
    return -0.044*x + 0.836*y + 0.169
}

func W2x(x, y float64) float64 {
    return -0.141*x + 0.302*y
}

func W2y(x, y float64) float64 {
    return 0.302*x + 0.141*y + 0.127
}

func W3x(x, y float64) float64 {
    return 0.141*x - 0.302*y
}

func W3y(x, y float64) float64 {
    return 0.302*x + 0.141*y + 0.169
}

func W4x(x, y float64) float64 {
    return 0
}

func W4y(x, y float64) float64 {
    return 0.175337 * y
}

/**
 * 描画
 * @param   m *image.RGBA   [image.RGBAのポインタ]
 * @param   k int           [最大描画数]
 * @param   x float64       [x座標]
 * @param   y float64       [y座標]
 */
func dotDraw(m *image.RGBA, k int, x, y float64) {

    if 0 < k {
        dotDraw(m, k-1, W1x(x, y), W1y(x, y))

        // rand.Float64() 0.0 ~ 1.0 までの乱数
        if rand.Float64() < 0.3 {
            dotDraw(m, k-1, W2x(x, y), W2y(x, y))
        }
        if rand.Float64() < 0.3 {
            dotDraw(m, k-1, W2x(x, y), W2y(x, y))
        }
        if rand.Float64() < 0.3 {
            dotDraw(m, k-1, W3x(x, y), W3y(x, y))
        }
        if rand.Float64() < 0.3 {
            dotDraw(m, k-1, W4x(x, y), W4y(x, y))
        }
    } else {
        var s float64 = 490.0

        // 描画位置を指定
        m.Set(int(x*s+float64(width)*0.5), int(float64(height)-y*s), linecolor)
    }
}


func main() {

    // 時刻を乱数の種に設定して擬似乱数を発生させる
    rand.Seed(time.Now().Unix())

    // 画像領域の確保
    m := image.NewRGBA(image.Rect(0, 0, width, height))

    // Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
    draw.Draw(m, m.Bounds(), &image.Uniform{bgcolor}, image.ZP, draw.Src)

    dotDraw(m, N, 0, 0)

    // 画像データの書き込みはio.Writerを用意しておいて、Encode
    file, err := os.Create(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    err = png.Encode(file, m)
    if err != nil {
        log.Fatal(err)
    }
}

