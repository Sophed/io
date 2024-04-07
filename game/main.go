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
	//rl.SetTargetFPS(60)

	player := engine.CreatePlayer(
		100, 300,
		"player.png",
	)
	test_platform := engine.CreateObject(
		100, 400,
		700, 500,
	)
	other_platform := engine.CreateObject(
		800, 450,
		1200, 600,
	)
	wall := engine.CreateObject(
		1000, 250,
		1050, 450,
	)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		player.Update()

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			player.Entity.Move(
				engine.Vec2{
					X: float32(rl.GetMouseX()),
					Y: float32(rl.GetMouseY()),
				},
			)
		}

		player.Entity.Draw()
		test_platform.Draw()
		other_platform.Draw()
		wall.Draw()

		player.Entity.Hitbox.Draw()
		test_platform.Hitbox.Draw()
		other_platform.Hitbox.Draw()
		wall.Hitbox.Draw()
		rl.DrawFPS(16, 16)

		rl.EndDrawing()
	}
}
