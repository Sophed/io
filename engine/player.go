package engine

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Entity  Entity
	Jumping bool
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
		false,
	}
}

const PLAYER_ACCELERATION = 4
const PLAYER_MAX_SPEED = 500
const PLAYER_JUMP_STRENGTH = 60

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

	if p.Entity.OnGround() {
		p.Jumping = false
	}

	if rl.IsKeyDown(rl.KeySpace) {
		p.Jump()
	}

	p.Entity.Update()
	fmt.Println("{}", p.Entity.Vel.Y)

}

func (p *Player) Jump() {
	if !p.Entity.OnGround() || p.Jumping {
		return
	}
	p.Jumping = true
	p.Entity.Vel.Y -= PLAYER_JUMP_STRENGTH * 2
}
