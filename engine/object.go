package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type GameObject struct {
	Hitbox Hitbox
	//Texture rl.Texture2D
}

func CreateObject(x1, y1, x2, y2 float32) GameObject {
	hitbox := Hitbox{
		Vec2{x1, y1},
		Vec2{x2, y2},
	}
	object := GameObject{hitbox}
	GAME_OBJECTS = append(GAME_OBJECTS, object)
	return object
}

func (o *GameObject) Draw() {
	rl.DrawRectangle(
		int32(o.Hitbox.TopLeft.X),
		int32(o.Hitbox.TopLeft.Y),
		int32(o.Hitbox.Width()),
		int32(o.Hitbox.Height()),
		rl.Blue,
	)
}
