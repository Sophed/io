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

	engine.LoadMap("map.png")

	player := engine.CreatePlayer(
		100, 300,
		"player.png",
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

		for _, o := range engine.GAME_OBJECTS {
			o.Draw()
			o.Hitbox.Draw()
		}

		player.Entity.Draw()

		player.Entity.Hitbox.Draw()
		rl.DrawFPS(16, 16)

		rl.EndDrawing()
	}
}
