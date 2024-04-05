package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sophed/io/engine"
)

func main() {

	//rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(
		340*engine.PIXEL_SCALE,
		180*engine.PIXEL_SCALE,
		"io",
	)
	defer rl.CloseWindow()

	player := engine.CreatePlayer(
		100, 300,
		rl.LoadTexture(engine.ASSETS_DIR+"player.png"),
	)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		player.Update()

		player.Entity.Draw()

		player.Entity.Hitbox.Draw()
		rl.DrawFPS(16, 16)

		rl.EndDrawing()
	}
}
