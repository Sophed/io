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

const ENTITY_GRAVITY = 2000

func (e *Entity) Update() {

	if e.Vel.X > 0 {
		e.Vel.X--
	} else if e.Vel.X < 0 {
		e.Vel.X++
	}
	e.Gravity()

	new_pos := Vec2{
		e.Pos.X + e.Vel.X*rl.GetFrameTime(),
		e.Pos.Y + e.Vel.Y*rl.GetFrameTime(),
	}

	// We'll figure this out later
	/*for _, o := range GAME_OBJECTS {
		if e.Hitbox.Colliding(o.Hitbox) {
			if e.Hitbox.TopLeft.X < o.Hitbox.TopLeft.X {
				new_pos.X = e.Pos.X - rl.GetFrameTime() - 1
			} else {
				new_pos.X = e.Pos.X + rl.GetFrameTime() + 1
			}
		}
	}*/

	bottom_right := Vec2{
		new_pos.X + float32(e.Texture.Height),
		new_pos.Y + float32(e.Texture.Width),
	}
	e.Pos = new_pos
	e.Hitbox = Hitbox{new_pos, bottom_right}
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
	simulated_vec1 := Vec2{e.Hitbox.BottomRight.X, e.Hitbox.BottomRight.Y + 1}
	simulated_vec2 := Vec2{e.Hitbox.TopLeft.X, e.Hitbox.BottomRight.Y + 1}
	for _, o := range GAME_OBJECTS {
		if o.Hitbox.Contains(simulated_vec1) || o.Hitbox.Contains(simulated_vec2) {
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
