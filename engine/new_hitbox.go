package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type NewHitbox struct {
	Origin Vec2
	Closed bool
	Points []Vec2
	Width  int
	Height int
}

func CreateNewHitbox(origin Vec2, points []Vec2, close bool) *NewHitbox {
	box := NewHitbox{origin, close, points, 0, 0}
	box.Width = box.SetWidth()
	box.Height = box.SetHeight()
	return &box
}

func (h *NewHitbox) SetWidth() int {
	max_x := 0
	for _, p := range h.Points {
		if int(p.X) > max_x {
			max_x = int(p.X)
		}
	}
	min_x := MaxInt
	for _, p := range h.Points {
		if int(p.X) < min_x {
			min_x = int(p.X)
		}
	}
	return max_x - min_x
}

func (h *NewHitbox) SetHeight() int {
	max_y := 0
	for _, p := range h.Points {
		if int(p.Y) > max_y {
			max_y = int(p.Y)
		}
	}
	min_y := MaxInt
	for _, p := range h.Points {
		if int(p.Y) < min_y {
			min_y = int(p.Y)
		}
	}
	return max_y - min_y
}

func (h *NewHitbox) Center() Vec2 {
	return Vec2{
		h.Origin.X + float32(h.Width/2),
		h.Origin.Y + float32(h.Height/2),
	}
}

func (o *NewHitbox) Collidiing(p *NewHitbox) bool {

	for i := 0; i < len(o.Points)-1; i++ {
		a := o.Points[i]
		b := o.Points[i+1]

		for i := 0; i < len(p.Points)-1; i++ {
			c := o.Points[i]
			d := o.Points[i+1]

			interesect := ccw(a, b, d) != ccw(b, c, d) && ccw(a, b, c) != ccw(a, b, d)
			if interesect {
				return true
			}

		}

	}
	return false
}

func ccw(a Vec2, b Vec2, c Vec2) bool {
	return (c.Y-a.Y)*(b.X-a.X) > (b.Y-a.Y)*(c.X-a.X)
}

func (h *NewHitbox) Draw() {

	for i, p := range h.Points {

		if i+1 > (len(h.Points) - 1) {
			if !h.Closed {
				return
			}
			target := h.Points[0]
			rl.DrawLine(
				int32(h.Origin.X+p.X),
				int32(h.Origin.Y+p.Y),
				int32(h.Origin.X+target.X),
				int32(h.Origin.Y+target.Y),
				rl.White,
			)
			return
		}

		target := h.Points[i+1]
		rl.DrawLine(
			int32(h.Origin.X+p.X),
			int32(h.Origin.Y+p.Y),
			int32(h.Origin.X+target.X),
			int32(h.Origin.Y+target.Y),
			rl.White,
		)

	}

}
