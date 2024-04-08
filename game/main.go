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

	test_hitbox := engine.CreateNewHitbox(
		engine.Vec2{
			X: 100, Y: 100,
		},
		[]engine.Vec2{
			{
				X: 300, Y: 100,
			},
			{
				X: 200, Y: 200,
			},
			{
				X: 100, Y: 200,
			},
			{
				X: 0, Y: 0,
			},
		},
	)

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
			player.Entity.Vel.X = 0
			player.Entity.Vel.Y = 0
		}

		if rl.IsMouseButtonDown(rl.MouseButtonRight) {
			test_hitbox.Origin = engine.Vec2{
				X: float32(rl.GetMouseX()),
				Y: float32(rl.GetMouseY()),
			}
		}

		for _, o := range engine.GAME_OBJECTS {
			o.Draw()
			o.Hitbox.Draw()
		}

		test_hitbox.Draw()
		rl.DrawPixel(
			int32(test_hitbox.Center().X),
			int32(test_hitbox.Center().Y),
			rl.Red,
		)

		player.Entity.Draw()

		player.Entity.Hitbox.Draw()
		rl.DrawFPS(16, 16)

		rl.EndDrawing()
	}
}
