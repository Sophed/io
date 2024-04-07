package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func LoadMap(path string) {

	image := rl.LoadImage(ASSETS_DIR + path)

	scaler := float32(PIXEL_SCALE * 16)

	for w := range image.Width {
		for h := range image.Height {

			pixel := image.ToImage().At(int(w), int(h))

			if pixel == rl.White {
				GAME_OBJECTS = append(GAME_OBJECTS,
					CreateObject(
						float32(w)*scaler,
						float32(h)*scaler,
						(float32(w)*scaler)+scaler,
						(float32(h)*scaler)+scaler,
					),
				)
			}

		}
	}

}
