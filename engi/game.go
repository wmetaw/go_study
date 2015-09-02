package main

import (
	// "fmt"
	"github.com/ajhager/engi"
)

// 型宣言
type AnimeType int

// 重力
var GRAVITY float32 = 9.6

var windowHeight float32 = 800
var windowWidth float32 = 1200

// アニメーション定数
const (
	ANIME_TYPE_WAIT AnimeType = iota
	ANIME_TYPE_RUN
	ANIME_TYPE_JUMP
	ANIME_TYPE_SHOOT
)

var JUMP_COUNTER = [...]float32{-110, -90, -60, -30, -20, -10, -5, -3, -1, 0, 1, 3, 10, 20, 30, 40, 50, 70}

//=================================
// ゲームオブジェクト
//=================================
type Game struct {
	*engi.Game

	// ゴーファーくん構造体
	Player

	// 描画用バッチ？
	batch *engi.Batch

	// フォント
	font *engi.Font

	// BG
	bg engi.Drawable

	// キーマッピング
	keyMap map[engi.Key]bool
}

//=================================
// 弾構造体
//=================================
type Bullet struct {
	img  engi.Drawable
	posX float32
	posY float32
	flg  bool
}


// 弾更新
func (bt Bullet) Update(bullet *Bullet, dt float32) {

	if bullet.flg == true {
		bullet.posX += 10

		if bullet.posX > windowWidth {
			bullet.flg = false
		}
	}
}


//=================================
// プレイヤー構造体
//=================================
type Player struct {
	img  engi.Drawable
	posX float32
	posY float32

	// アニメーション
	animeNo     int
	animeType   AnimeType
	animeFrame  float32
	waitImages  [2]engi.Drawable // 待機モーション
	runImages   [3]engi.Drawable // 移動モーション
	jumpImages  [2]engi.Drawable // ジャンプモーション
	shootImages [3]engi.Drawable // 発砲モーション

	isJump    bool
	jumpCnt   int
	bulletMax int
	bullets   []Bullet
}

// プレイヤー更新
func (pl Player) Update(game *Game, dt float32) {

	// 慣性ジャンプ
	if game.Player.isJump && game.Player.jumpCnt < len(JUMP_COUNTER) {
		game.posY += JUMP_COUNTER[game.Player.jumpCnt]
		game.Player.jumpCnt++
	}

	// 落下
	game.Player.posY += GRAVITY
	if game.Player.posY > 670 {
		game.Player.posY = 670
		game.Player.jumpCnt = 0
		game.Player.isJump = false
	}

	if game.Player.posX < 230 {
		game.Player.posX = 230
	}

	// fmt.Printf("X：%f", game.Player.posX)
	// fmt.Printf("Y：%f\n", game.Player.posY)

	if game.Player.posX > windowWidth + game.Player.img.Width()/4 {
		game.Player.posX = windowWidth + game.Player.img.Width()/4
	}

	// 5フレーム後にアニメーションを切り替え
	var isChange bool
	if game.Player.animeFrame >= (dt * 5) {
		game.Player.animeFrame = dt
		isChange = true
	} else {
		game.Player.animeFrame += dt
	}

	// アニメーション切り替え
	if isChange {
		switch game.Player.animeType {

		// 待機モーション
		case ANIME_TYPE_WAIT:

			// 画像の要素数を超えたらアニメ番号を0に戻す
			if game.Player.animeNo >= len(game.Player.waitImages) {
				game.Player.animeNo = 0
			}

			// キャラ画像を置換
			game.Player.img = game.Player.waitImages[game.Player.animeNo]

		// 走るモーション
		case ANIME_TYPE_RUN:
			if game.Player.animeNo >= len(game.Player.runImages) {
				game.Player.animeNo = 0
			}
			game.Player.img = game.Player.runImages[game.Player.animeNo]

		// 発砲モーション
		case ANIME_TYPE_SHOOT:
			if game.Player.animeNo >= len(game.Player.shootImages) {
				game.Player.animeNo = 0
			}
			game.Player.img = game.Player.shootImages[game.Player.animeNo]

		// ジャンプモーション
		case ANIME_TYPE_JUMP:
			if game.Player.animeNo > 0 {
				game.Player.animeNo = 1
			}
			game.Player.img = game.Player.jumpImages[game.Player.animeNo]
		}

		game.Player.animeNo++
	}

	if game.Player.animeType == ANIME_TYPE_SHOOT && isChange{

		// 弾初期化
		for i := 0; i < len(game.Player.bullets); i++ {
			if game.Player.bullets[i].flg == false {
				game.Player.bullets[i].posX = game.Player.posX
				game.Player.bullets[i].posY = game.Player.posY - game.Player.img.Height()/2
				game.Player.bullets[i].flg = true
				break
			}
		}
	}

	// 弾更新
	for i := 0; i < len(game.Player.bullets); i++ {
		game.Player.bullets[i].Update(&game.Player.bullets[i], dt)
	}

}



/**
 * 更新
 *
 * 【引数】時間差 float32(0.01635106とか）
 */
func (game *Game) Update(dt float32) {

	var dx float32
	game.Player.animeType = ANIME_TYPE_WAIT
	if game.keyMap[engi.ArrowLeft] {
		dx = -10
		game.Player.animeType = ANIME_TYPE_RUN
	}
	if game.keyMap[engi.ArrowRight] {
		dx = 10
		game.Player.animeType = ANIME_TYPE_RUN
	}
	if game.keyMap[engi.D] {
		game.Player.animeType = ANIME_TYPE_SHOOT
	}
	if game.keyMap[engi.Space] {
		game.Player.isJump = true
	}
	if game.Player.isJump {
		game.Player.animeType = ANIME_TYPE_JUMP
	}

	game.Player.posX += dx

	game.Player.Update(game, dt)
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
 * 描画
 *
 * engiが描画ループしてくれるっぽい
 */
func (game *Game) Render() {

	game.batch.Begin()

	// フォントの表示
	game.font.Print(game.batch, "ENGI", 475, 200, 0x000000)

	// 表示
	game.batch.Draw(game.bg, 1200, 760, 1, 1, 1, 0.95, 0, 0xffffff, 1)
	game.batch.Draw(game.Player.img, game.Player.posX, game.Player.posY, 1, 1, 1, 1, 0, 0xffffff, 1)

	// プレイヤーの弾描画
	for i := 0; i < len(game.Player.bullets); i++ {
		if game.Player.bullets[i].flg == true {
			game.batch.Draw(game.Player.bullets[i].img, game.Player.bullets[i].posX, game.Player.bullets[i].posY, 1, 1, 1, 1, 0, 0xffffff, 1)
		}
	}

	game.batch.End()
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

	engi.Files.Add("bg", "img/bg.png")

	engi.Files.Add("font", "img/font.png")

	engi.Files.Add("bom", "img/bullet/bom.png")

	game.batch = engi.NewBatch(engi.Width(), engi.Height())

	// キーマップの作成
	game.keyMap = make(map[engi.Key]bool)

	// プレイヤーキャラ初期化(TODO初期化関数を作る)
	game.Player.posX, game.Player.posY = 300, 500
	game.Player.animeFrame = 0
	game.Player.animeNo = 0
	game.Player.animeType = ANIME_TYPE_WAIT
	game.Player.jumpCnt = 0
	game.Player.bulletMax = 20
}

/**
 * セットアップ
 *
 * Preloadの次に呼ばれる
 */
func (game *Game) Setup() {

	// 背景
	engi.SetBg(0xfffffa)
	game.Player.waitImages[0] = engi.Files.Image("plWait0")
	game.Player.waitImages[1] = engi.Files.Image("plWait1")

	game.Player.runImages[0] = engi.Files.Image("plRun0")
	game.Player.runImages[1] = engi.Files.Image("plRun1")
	game.Player.runImages[2] = engi.Files.Image("plRun2")

	game.Player.shootImages[0] = engi.Files.Image("plShoot0")
	game.Player.shootImages[1] = engi.Files.Image("plShoot1")
	game.Player.shootImages[2] = engi.Files.Image("plShoot2")

	game.Player.jumpImages[0] = engi.Files.Image("plJump0")
	game.Player.jumpImages[1] = engi.Files.Image("plJump1")

	game.bg = engi.Files.Image("bg")

	for i := 0; i < game.Player.bulletMax; i++ {
		game.Player.bullets = append(game.Player.bullets, Bullet{engi.Files.Image("bom"), 0, 0, false})
	}

	game.Player.img = game.Player.waitImages[0]
	game.font = engi.NewGridFont(engi.Files.Image("font"), 20, 20)
}
