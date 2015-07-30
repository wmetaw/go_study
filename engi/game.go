package main

import (
"github.com/ajhager/engi"
// "fmt"
// "reflect"
)

var windowHeight float32 = 800
var windowWidth  float32 = 1200

// 重力
const GRAVITY float32 = 5

// アニメーション定数
const (
    ANIME_TYPE_WAIT = iota
    ANIME_TYPE_RUN
    ANIME_TYPE_JUMP
    ANIME_TYPE_SHOOT
)

var JUMP_COUNTER = [...]float32 {-110,-90,-60,-30,-20,-10,-5,-3,-1,0,1,3,10,20,30,40,50,70}

// ゲームオブジェクト
type Game struct {
    *engi.Game

    // ゴーファーくん構造体
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
    posX float32
    posY float32

    // アニメーション
    animeNo int
    animeType int
    animeFrame float32
    waitImages  [2]engi.Drawable // 待機モーション
    runImages   [3]engi.Drawable // 移動モーション
    jumpImages  [2]engi.Drawable // ジャンプモーション
    shootImages [3]engi.Drawable // 発砲モーション

    isJump bool
    jumpCnt int
}

func (pl PL) Update(game *Game, dt float32) {

    // 慣性ジャンプ
    if game.PL.isJump && game.PL.jumpCnt < len(JUMP_COUNTER) {
        game.posY += JUMP_COUNTER[game.PL.jumpCnt]
        game.PL.jumpCnt++
    }

    // 落下
    game.PL.posY += GRAVITY
    if game.PL.posY > windowHeight - game.PL.img.Height()/4 {
        game.PL.posY = windowHeight - game.PL.img.Height()/4
        game.PL.jumpCnt = 0
        game.PL.isJump = false
    }

    // 5フレーム後にアニメーションを切り替え
    var isChange bool
    if game.PL.animeFrame >= (dt*5) {
        game.PL.animeFrame = dt
        isChange = true
    }  else {
        game.PL.animeFrame += dt
    }

    // アニメーション切り替え
    if isChange {
        switch game.PL.animeType {

        // 待機モーション
        case ANIME_TYPE_WAIT:

            // 画像の要素数を超えたらアニメ番号を0に戻す
            if game.PL.animeNo >= len(game.PL.waitImages) {
                game.PL.animeNo = 0
            }

            // キャラ画像を置換
            game.PL.img = game.PL.waitImages[game.PL.animeNo]

        // 走るモーション
        case ANIME_TYPE_RUN:
            if game.PL.animeNo >= len(game.PL.runImages) {
                game.PL.animeNo = 0
            }
            game.PL.img = game.PL.runImages[game.PL.animeNo]

        // 発砲モーション
        case ANIME_TYPE_SHOOT:
            if game.PL.animeNo >= len(game.PL.shootImages) {
                game.PL.animeNo = 0
            }
            game.PL.img = game.PL.shootImages[game.PL.animeNo]

        // ジャンプモーション
        case ANIME_TYPE_JUMP:
            if game.PL.animeNo > 0 {
                game.PL.animeNo = 1
            }
            game.PL.img = game.PL.jumpImages[game.PL.animeNo]
        }

        game.PL.animeNo++
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

    var dx float32
    game.PL.animeType = ANIME_TYPE_WAIT
    if game.keyMap[engi.ArrowLeft] {
        dx = -10
        game.PL.animeType = ANIME_TYPE_RUN
    }
    if game.keyMap[engi.ArrowRight] {
        dx = 10
        game.PL.animeType = ANIME_TYPE_RUN
    }
    if game.keyMap[engi.S] {
        game.PL.animeType = ANIME_TYPE_SHOOT
    }
    if game.keyMap[engi.Space] {
        game.PL.isJump = true
    }
    if game.PL.isJump {
        game.PL.animeType = ANIME_TYPE_JUMP
    }

    game.PL.posX += dx

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
    engi.Files.Add("plWait0", "img/player/wait/0.png")
    engi.Files.Add("plWait1", "img/player/wait/1.png")

    engi.Files.Add("plRun0", "img/player/run/0.png")
    engi.Files.Add("plRun1", "img/player/run/1.png")
    engi.Files.Add("plRun2", "img/player/run/2.png")

    engi.Files.Add("plShoot0", "img/player/shoot/0.png")
    engi.Files.Add("plShoot1", "img/player/shoot/1.png")
    engi.Files.Add("plShoot2", "img/player/shoot/2.png")

    engi.Files.Add("plJump0", "img/player/jump/0.png")
    engi.Files.Add("plJump1", "img/player/jump/1.png")

    engi.Files.Add("font", "img/font.png")

    game.batch = engi.NewBatch(engi.Width(), engi.Height())

    // キーマップの作成
    game.keyMap = make(map[engi.Key]bool)

    // プレイヤーキャラ初期化
    game.PL.posX, game.PL.posY = 512, 320
    game.PL.animeFrame = 0
    game.PL.animeNo = 0
    game.PL.animeType = ANIME_TYPE_WAIT
    game.PL.jumpCnt = 0;
}

/**
 * セットアップ
 *
 * Preloadの次に呼ばれる
 */
func (game *Game) Setup() {

    // 背景
    engi.SetBg(0xfffffa)
    game.PL.waitImages[0] = engi.Files.Image("plWait0")
    game.PL.waitImages[1] = engi.Files.Image("plWait1")

    game.PL.runImages[0]  = engi.Files.Image("plRun0")
    game.PL.runImages[1]  = engi.Files.Image("plRun1")
    game.PL.runImages[2]  = engi.Files.Image("plRun2")

    game.PL.shootImages[0]  = engi.Files.Image("plShoot0")
    game.PL.shootImages[1]  = engi.Files.Image("plShoot1")
    game.PL.shootImages[2]  = engi.Files.Image("plShoot2")

    game.PL.jumpImages[0] = engi.Files.Image("plJump0")
    game.PL.jumpImages[1] = engi.Files.Image("plJump1")

    game.PL.img = game.PL.waitImages[0]
    game.font = engi.NewGridFont(engi.Files.Image("font"), 20, 20)
}



