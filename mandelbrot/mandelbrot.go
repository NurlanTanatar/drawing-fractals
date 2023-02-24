package mandelbrot

import (
	"math"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	FPS       int32 = 60
	winHeight int32 = 900
	winWidth  int32 = 600

	// zoom sett-s
	renderHeight int32 = winHeight * 4
	renderWidth  int32 = winWidth * 4

	// mandelbrot set range
	mMin float64 = -2
	mMax float64 = 2

	// depth: z_i+1 = f(z_i)
	iterations int = 100

	// sliders
	mMinSlider float64 = mMin
	mMaxSlider float64 = mMax
	iterSlider float64 = float64(iterations)
	offsetX    int32   = int32(float64((-1)*renderHeight) / 3.0)
	offsetY    int32   = int32(float64((-1)*renderWidth) / 2.4)
)

// linear mapping: x in [xMin, xMax] -> y in [outMin, outMax]
func mapTo(x, xMin, xMax, outMin, outMax float64) float64 {
	return (x-xMin)*(outMax-outMin)/(xMax-xMin) + outMin
}

func calcColor(seed uint8) rl.Color {
	return rl.NewColor(
		uint8(mapTo(math.Sqrt(float64(seed)), 0, math.Sqrt(255), 0, 255)),
		uint8(mapTo(float64(seed*seed), 0, 255*255, 0, 255)),
		// uint8(mapTo(math.Log(float64(seed)), 0, math.Log(255), 0, 255)),
		uint8(mapTo(0.5*float64(seed), 0, math.Log(255), 0, 255)),
		255,
	)
}

func DrawMandelbrot() {
	rl.InitWindow(winHeight, winWidth, "raylib [core] example - basic window")
	rl.SetTargetFPS(FPS)

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyQ) {
			mMinSlider /= 1.05
		} else if rl.IsKeyDown(rl.KeyW) {
			mMinSlider *= 1.05
		}

		if rl.IsKeyDown(rl.KeyA) {
			mMaxSlider /= 1.05
		} else if rl.IsKeyDown(rl.KeyS) {
			mMaxSlider *= 1.05
		}

		if rl.IsKeyDown(rl.KeyZ) {
			iterSlider /= 1.05
		} else if rl.IsKeyDown(rl.KeyX) {
			iterSlider *= 1.05
		}

		if rl.IsKeyDown(rl.KeyUp) {
			offsetY += 30
		} else if rl.IsKeyDown(rl.KeyDown) {
			offsetY -= 30
		} else if rl.IsKeyDown(rl.KeyLeft) {
			offsetX += 30
		} else if rl.IsKeyDown(rl.KeyRight) {
			offsetX -= 30
		}

		if rl.IsKeyPressed(rl.KeyR) {
			mMinSlider = mMin
			mMaxSlider = mMax
			iterSlider = float64(iterations)
			offsetX = 0
			offsetY = 0
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for y := 0; y < int(renderWidth); y++ {
			if int32(y)+offsetY < 0 || int32(y)+offsetY > winWidth {
				continue
			}
			for x := 0; x < int(renderHeight); x++ {
				if int32(x)+offsetX < 0 || int32(x)+offsetX > winHeight {
					continue
				}

				z := complex(mapTo(float64(x), 0, float64(renderHeight), mMinSlider, mMaxSlider), mapTo(float64(y), 0, float64(renderWidth), mMinSlider, mMaxSlider))
				zi := z
				iter := 0
				for ; iter < int(iterSlider); iter++ {
					zn := z * z
					z = zn + zi
					if real(z)+imag(z) > mMaxSlider {
						break
					}
				}
				var colorSeed uint8 = uint8(mapTo(float64(iter), 0, iterSlider, 0, 255))
				if iter == int(iterSlider) {
					colorSeed = 0
				} else {
					rl.DrawPixel(int32(x)+offsetX, int32(y)+offsetY, calcColor(colorSeed))
				}
				//else if x+int(offsetX) >= 0 || y+int(offsetY) >= 0 {
				// 	rl.DrawPixel(int32(x)+offsetX, int32(y)+offsetY, calcColor(colorSeed))
				// }

			}

		}
		rl.DrawText(strconv.Itoa(int(iterSlider)), 10, 60, 21, rl.White)
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
