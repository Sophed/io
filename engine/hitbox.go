package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Hitbox struct {
	TopLeft, BottomRight Vec2
}

func (h *Hitbox) Width() float32 {
	return h.BottomRight.X - h.TopLeft.X
}

func (h *Hitbox) Height() float32 {
	return h.BottomRight.Y - h.TopLeft.Y
}

func (h *Hitbox) Center() Vec2 {
	return Vec2{
		h.TopLeft.X + (h.Width() / 2),
		h.TopLeft.Y + (h.Height() / 2),
	}
}

func (h *Hitbox) Colliding(other Hitbox) bool {
	return h.TopLeft.X < other.BottomRight.X &&
		h.BottomRight.X > other.TopLeft.X &&
		h.TopLeft.Y < other.BottomRight.Y &&
		h.BottomRight.Y > other.TopLeft.Y
}

func (h *Hitbox) Contains(other Vec2) bool {
	return h.TopLeft.X <= other.X &&
		h.BottomRight.X >= other.X &&
		h.TopLeft.Y <= other.Y &&
		h.BottomRight.Y >= other.Y
}

func (h *Hitbox) Draw() {
	rl.DrawRectangleLines(
		int32(h.TopLeft.X),
		int32(h.TopLeft.Y),
		int32(h.Width()),
		int32(h.Height()),
		rl.Red,
	)
}
