package main

import rl "github.com/gen2brain/raylib-go/raylib"

func windowInit() {
	rl.InitWindow(100, 250, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawText("WeatherAppV2 Starting ...", 190, 200, 20, rl.RayWhite)

		rl.EndDrawing()
	}
}
