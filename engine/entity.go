package engine

import rl "github.com/gen2brain/raylib-go/raylib"

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

func (e *Entity) Update() {

	if e.Vel.X > 0 {
		e.Vel.X--
	} else if e.Vel.X < 0 {
		e.Vel.X++
	}

	new_pos := Vec2{
		e.Pos.X + e.Vel.X*rl.GetFrameTime(),
		e.Pos.Y + e.Vel.Y*rl.GetFrameTime(),
	}
	e.Pos = new_pos
	bottom_right := Vec2{
		e.Pos.X + float32(e.Texture.Height),
		e.Pos.Y + float32(e.Texture.Width),
	}
	e.Hitbox = Hitbox{new_pos, bottom_right}
}

func (e *Entity) Draw() {
	x := int32(e.Pos.X)
	y := int32(e.Pos.Y)
	rl.DrawTexture(e.Texture, x, y, rl.White)
}
