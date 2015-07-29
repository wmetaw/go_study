package main

import (
"github.com/ajhager/engi"
// "fmt"
// "reflect"
)

var gravity float32 = 25
var windowHeight float32 = 800
var windowWidth  float32 = 1200


// ゲームオブジェクト
type Game struct {
    *engi.Game

    // ゴーファーくん
    PL

    // 描画用バッチ？
    batch *engi.Batch

    // フォント
    font  *engi.Font

    // キーマッピング
    keyMap map[engi.Key]bool
}

// ゴーファーくん
type PL struct {

    img engi.Drawable
    animeNo int
    posX float32
    posY float32
    waitFrame float32
    waitImg [2]engi.Drawable
    // isJump bool = false
}

func (pl PL) Update(game *Game, dt float32) {

    // 落下
    game.PL.posY += gravity
    if game.PL.posY > windowHeight - game.PL.img.Height()/4 {
        game.PL.posY = windowHeight - game.PL.img.Height()/4
    }

    // 5フレーム後にアニメーションを切り替え
    var isChange bool
    if game.PL.waitFrame >= (dt*5) {
        game.PL.waitFrame = dt
        isChange = true
    }  else {
        game.PL.waitFrame += dt
    }

    // アニメーション切り替え
    if isChange {

        // キャラ画像を置換
        game.PL.img = game.PL.waitImg[game.PL.animeNo]

        // アニメーション番号をカウントアップ
        game.PL.animeNo++
        if len(game.PL.waitImg) == game.PL.animeNo {
            game.PL.animeNo = 0
        }
    }

}

/**
 * 描画
 *
 * engiが描画ループしてくれるっぽい
 */
 func (game *Game) Render() {

    game.batch.Begin()

    // フォントの表示
    game.font.Print(game.batch, "ENGI", 475, 200, 0x000000)

    // ロゴを表示
    game.batch.Draw(game.PL.img, game.PL.posX, game.PL.posY, 1, 1, 1, 1, 0, 0xffffff, 1)

    game.batch.End()
}


/**
 * 更新
 *
 * 【引数】時間差 float32(0.01635106とか）
 */
    func (game *Game) Update(dt float32) {

        var dx, dy float32
        if game.keyMap[engi.ArrowUp] {
            dy = -30
        }
        if game.keyMap[engi.ArrowDown] {
            dy = 30
        }
        if game.keyMap[engi.ArrowLeft] {
            dx = -10
        }
        if game.keyMap[engi.ArrowRight] {
            dx = 10
        }
        if game.keyMap[engi.Space] {
            game.PL.posY -= 30
        }

        game.PL.posX += dx
        game.PL.posY += dy

        game.PL.Update(game, dt)
    }

/**
 * キー入力判定
 */
 func (game *Game) Key(key engi.Key, modifier engi.Modifier, action engi.Action) {

    // 対応するキーのフラグをON
    switch action {
    case engi.PRESS:
        game.keyMap[key] = true
    case engi.RELEASE:
        game.keyMap[key] = false
    }
}

/**
 * Main
 */
 func main() {

    //      タイトル, width, height, fullscreen, Gameオブジェクト
    engi.Open("Hello", 1200, 800, false, &Game{})
}



/**
 * ロード
 *
 * 最初に呼び出される
 */
 func (game *Game) Preload() {

    // spriteの読み込み
    engi.Files.Add("plWait1", "img/player/wait/1.png")
    engi.Files.Add("plWait2", "img/player/wait/2.png")
    engi.Files.Add("font", "img/font.png")

    game.batch = engi.NewBatch(engi.Width(), engi.Height())

    // キーマップの作成
    game.keyMap = make(map[engi.Key]bool)

    // プレイヤーキャラ初期化
    game.PL.posX, game.PL.posY = 512, 320
    game.PL.waitFrame = 0
    game.PL.animeNo = 0
}

/**
 * セットアップ
 *
 * Preloadの次に呼ばれる
 */
 func (game *Game) Setup() {

    // 背景
    engi.SetBg(0xfffffa)
    game.PL.waitImg[0] = engi.Files.Image("plWait1")
    game.PL.waitImg[1] = engi.Files.Image("plWait2")
    game.PL.img = game.PL.waitImg[0]
    game.font = engi.NewGridFont(engi.Files.Image("font"), 20, 20)
}



