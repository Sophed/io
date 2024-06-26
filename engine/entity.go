package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Entity struct {
	Pos     Vec2
	Vel     Vec2
	Dir     int
	Hitbox  Hitbox
	Texture rl.Texture2D
}

func CreateEntity(pos Vec2, texture rl.Texture2D) *Entity {

	top_left := Vec2{pos.X, pos.Y}

	texture.Width *= PIXEL_SCALE
	texture.Height *= PIXEL_SCALE

	bottom_right := Vec2{
		pos.X + float32(texture.Height),
		pos.Y + float32(texture.Width),
	}

	return &Entity{
		pos,
		Vec2{0, 0},
		1,
		Hitbox{top_left, bottom_right},
		texture,
	}

}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

const ENTITY_GRAVITY = 2000

func (e *Entity) Move(pos Vec2) {
	e.Pos = pos
	bottom_right := Vec2{
		pos.X + float32(e.Texture.Height),
		pos.Y + float32(e.Texture.Width),
	}
	e.Hitbox = Hitbox{pos, bottom_right}
}

func (e *Entity) Update() {

	decel := 2000 * rl.GetFrameTime()

	if e.Vel.X > 0 {
		e.Vel.X = max(e.Vel.X-decel, 0)
	} else if e.Vel.X < 0 {
		e.Vel.X = min(e.Vel.X+decel, 0)
	}
	e.Gravity()

	new_pos := Vec2{
		e.Pos.X + e.Vel.X*rl.GetFrameTime(),
		e.Pos.Y + e.Vel.Y*rl.GetFrameTime(),
	}

	for _, o := range GAME_OBJECTS {
		if e.Hitbox.Colliding(o.Hitbox) {
			if e.Hitbox.TopLeft.X < o.Hitbox.TopLeft.X {
				new_pos.X = e.Pos.X - rl.GetFrameTime()
			} else {
				new_pos.X = e.Pos.X + rl.GetFrameTime() + 1
			}
		}
	}

	e.Move(new_pos)
}

func (e *Entity) Gravity() {
	if e.OnGround() {
		if e.Vel.Y > 0 {
			e.Vel.Y = 0
		}
		return
	}
	e.Vel.Y += ENTITY_GRAVITY * rl.GetFrameTime()
}

func (e *Entity) OnGround() bool {
	simulated_vec1 := Vec2{e.Hitbox.BottomRight.X, e.Hitbox.BottomRight.Y}
	simulated_vec2 := Vec2{e.Hitbox.TopLeft.X, e.Hitbox.BottomRight.Y}
	for _, o := range GAME_OBJECTS {
		if (o.Hitbox.Contains(simulated_vec1) || o.Hitbox.Contains(simulated_vec2)) &&
			simulated_vec1.Y-o.Hitbox.TopLeft.Y < 4*PIXEL_SCALE {
			e.Move(Vec2{
				e.Pos.X,
				o.Hitbox.TopLeft.Y - float32(e.Texture.Height),
			})
			return true
		}
	}
	return false
}

func (e *Entity) Draw() {
	x := int32(e.Pos.X)
	y := int32(e.Pos.Y)
	rl.DrawTexture(e.Texture, x, y, rl.White)
}
