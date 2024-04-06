package engine

import (
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Entity     Entity
	Jumping    bool
	JumpBuffer bool
	LastJump   int64
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
		false,
		0,
	}
}

const PLAYER_ACCELERATION = 5000
const PLAYER_MAX_SPEED = 450
const PLAYER_JUMP_STRENGTH = 60

func (p *Player) Update() {

	accel := PLAYER_ACCELERATION * rl.GetFrameTime()

	if rl.IsKeyDown(rl.KeyA) {
		if !(math.Abs(float64(p.Entity.Vel.X-accel)) > PLAYER_MAX_SPEED) {
			p.Entity.Vel.X -= accel
		}
		if !rl.IsKeyDown(rl.KeyD) {
			p.Entity.Dir = -1
		}
	}

	if rl.IsKeyDown(rl.KeyD) {
		if !(math.Abs(float64(p.Entity.Vel.X+accel)) > PLAYER_MAX_SPEED) {
			p.Entity.Vel.X += accel
		}
		if !rl.IsKeyDown(rl.KeyA) {
			p.Entity.Dir = 1
		}
	}

	if rl.IsKeyDown(rl.KeySpace) {
		p.Jump()
	}

	p.Entity.Update()

}

func (p *Player) Jump() {
	if !p.Entity.OnGround() || time.Now().UnixMilli()-p.LastJump <= 15 {
		return
	}
	p.Jumping = true
	p.LastJump = time.Now().UnixMilli()
	p.Entity.Vel.Y = 0
	p.Entity.Vel.Y -= PLAYER_JUMP_STRENGTH * 10
}
