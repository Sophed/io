package engine

import (
	"fmt"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Entity   Entity
	Jumping  bool
	LastJump int64
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
		0,
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

	if rl.IsKeyDown(rl.KeySpace) {
		p.Jump()
	}

	p.Entity.Update()
	fmt.Println("{}", p.LastJump)

}

func (p *Player) Jump() {
	if p.Jumping {
		if time.Now().UnixMilli()-p.LastJump > 25 {
			p.Jumping = false
		}
		return
	}
	if !p.Entity.OnGround() {
		return
	}
	p.Jumping = true
	p.LastJump = time.Now().UnixMilli()
	p.Entity.Vel.Y -= PLAYER_JUMP_STRENGTH * 10
}
