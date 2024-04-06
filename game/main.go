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
		"player.png",
	)
	test_platform := engine.CreateObject(
		100, 500,
		700, 600,
	)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		player.Entity.Draw()
		player.Update()
		test_platform.Draw()

		player.Entity.Hitbox.Draw()
		rl.DrawFPS(16, 16)

		rl.EndDrawing()
	}
}
