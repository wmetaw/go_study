package main

import (
"github.com/ajhager/engi"
)

// ゲームオブジェクト
type Game struct {
    *engi.Game

    // ロゴ画像
    bot   engi.Drawable

    // 描画用バッチ？
    batch *engi.Batch

    // フォント
    font  *engi.Font

    // ロゴの座標
    botX ,botY float32

    // キーマッピング
    keyMap map[engi.Key]bool
}

/**
 * ロード
 *
 * 最初に呼び出される
 */
 func (game *Game) Preload() {

    // spriteの読み込み
    engi.Files.Add("bot", "data/icon.png")
    engi.Files.Add("font", "data/font.png")
    game.batch = engi.NewBatch(engi.Width(), engi.Height())

    // キーマップの作成
    game.keyMap = make(map[engi.Key]bool)
    game.botX, game.botY = 512, 320
}

/**
 * セットアップ
 *
 * Preloadの次に呼ばれる
 */
 func (game *Game) Setup() {

    // 背景
    engi.SetBg(0x2d3739)
    game.bot = engi.Files.Image("bot")
    game.font = engi.NewGridFont(engi.Files.Image("font"), 20, 20)
}

/**
 * 描画
 *
 * engiが描画ループしてくれるっぽい
 */
 func (game *Game) Render() {

    game.batch.Begin()

    // フォントの表示
    game.font.Print(game.batch, "ENGI", 475, 200, 0xffffff)

    // ロゴを表示
    game.batch.Draw(game.bot, game.botX, game.botY, 0.5, 0.5, 10, 10, 0, 0xffffff, 1)

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
        dy = -10
    }
    if game.keyMap[engi.ArrowDown] {
        dy = 10
    }
    if game.keyMap[engi.ArrowLeft] {
        dx = -10
    }
    if game.keyMap[engi.ArrowRight] {
        dx = 10
    }

    game.botX += dx
    game.botY += dy
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

    //      タイトル, height, width, fullscreen, Gameオブジェクト
    engi.Open("Hello", 1024, 640, false, &Game{})
}



