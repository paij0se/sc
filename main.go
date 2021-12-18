package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	tl "github.com/JoelOtter/termloop"
	"github.com/paij0se/sc/music"
	ymp3cli "github.com/paij0se/ymp3cli/src/cli"
)

type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

type Enemy struct {
	*tl.Entity
}

func (Enemy *Enemy) Tick(event tl.Event) {
	// put the enemies in randoms positions
	Enemy.SetPosition(rand.Intn(100), rand.Intn(100))

}

func (player *Player) Draw(screen *tl.Screen) {
	// This is like the camera
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.Entity.Draw(screen)
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}
	}
}

func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
}

func (Enemy *Enemy) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		Enemy.SetPosition(rand.Intn(100), rand.Intn(100))
	}
	// Check if colliding with a player
	if _, ok := collision.(*Player); ok {
		go ymp3cli.PlaySongCli("music/v.mp3")
	}
}

func GenerateRandomEnemy() *Enemy {
	enemy := &Enemy{tl.NewEntity(1, 1, 1, 1)}
	enemy.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '💀'})
	return enemy
}

func main() {
	go music.PlayMusic()
	game := tl.NewGame()
	game.Screen().SetFps(30)
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorBlack,
		Ch: ' ',
	})
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorWhite))

	player := Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
		level:  level,
	}

	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorGreen, Ch: '옷'})
	level.AddEntity(&player)
	var numEnemies int
	fmt.Print(`
		Select A level

		Levels:
		1. Easy
		2. Medium
		3. Hard
		4. Impossible

		💀---<

		     
	$:`)

	fmt.Scanf("%d", &numEnemies)
	// set the difficulty
	switch numEnemies {
	case 1:
		for i := 0; i < 10; i++ {
			level.AddEntity(GenerateRandomEnemy())
		}
		text := tl.NewText(0, 1, "Number of enemies: "+strconv.Itoa(10), tl.ColorBlue, tl.ColorBlack)
		game.Screen().AddEntity(text)
	case 2:
		for i := 0; i < 33; i++ {
			level.AddEntity(GenerateRandomEnemy())
		}
		text := tl.NewText(0, 1, "Number of enemies: "+strconv.Itoa(33), tl.ColorBlue, tl.ColorBlack)
		game.Screen().AddEntity(text)
	case 3:
		for i := 0; i < 50; i++ {
			level.AddEntity(GenerateRandomEnemy())
		}
		text := tl.NewText(0, 1, "Number of enemies: "+strconv.Itoa(50), tl.ColorBlue, tl.ColorBlack)
		game.Screen().AddEntity(text)
	case 4:
		for i := 0; i < 100; i++ {
			level.AddEntity(GenerateRandomEnemy())
		}
		text := tl.NewText(0, 1, "Number of enemies: "+strconv.Itoa(100), tl.ColorBlue, tl.ColorBlack)
		game.Screen().AddEntity(text)
	default:
		fmt.Println("Invalid level")
		os.Exit(1)

	}
	game.Screen().AddEntity(tl.NewText(0, 0, "Don't let the enemies touch you!\npress <ctrl + c> to exit", tl.ColorBlue, tl.ColorBlack))
	game.Screen().SetLevel(level)
	game.Start()

}
