package engine

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Entity Entity
}

func CreatePlayer(x, y int, texture string) *Player {
	data := rl.LoadTexture(ASSETS_DIR + texture)
	return &Player{
		*CreateEntity(
			Vec2{
				float32(x),
				float32(y),
			},
			data,
		),
	}
}

const PLAYER_ACCELERATION = 4
const PLAYER_MAX_SPEED = 500

func (p *Player) Update() {

	if rl.IsKeyDown(rl.KeyA) {
		if !(math.Abs(float64(p.Entity.Vel.X-PLAYER_ACCELERATION)) > PLAYER_MAX_SPEED) {
			p.Entity.Vel.X -= PLAYER_ACCELERATION
		}
		if !rl.IsKeyDown(rl.KeyD) {
			p.Entity.Dir = -1
		}
	}

	if rl.IsKeyDown(rl.KeyD) {
		if !(math.Abs(float64(p.Entity.Vel.X+PLAYER_ACCELERATION)) > PLAYER_MAX_SPEED) {
			p.Entity.Vel.X += PLAYER_ACCELERATION
		}
		if !rl.IsKeyDown(rl.KeyA) {
			p.Entity.Dir = 1
		}
	}

	p.Entity.Update()

}
