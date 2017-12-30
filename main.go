package main

import (
	"runtime"//bununla buradaki inint bağlantılı neden ?
	"time"
	"stnc/tetris"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)



func init() {
	runtime.LockOSThread()
}

func main() {
	tetris.InitGame()
	// Create the window
	var err error
	err = glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	tetris.Window, err = glfw.CreateWindow(tetris.W, tetris.H, "go-tetris", nil, nil)
	if err != nil {
		panic(err)
	}
	tetris.Window.MakeContextCurrent()
	tetris.Window.SetKeyCallback(tetris.KeyPress)
	if err := gl.Init(); err != nil {
		panic(err)
	}
	// Timer
	ticker := time.NewTicker(time.Millisecond * tetris.TimerPeriod)
	go func() {
		for range ticker.C {
			//fmt.Println("tick ", t)
			tetris.Update()
		}
	}()
	// Init OpenGL
	gl.Ortho(0, tetris.W, tetris.H, 0, -1, 1)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	//gl.ClearColor(20, 20, 0, 0)
	gl.ClearColor(255, 255, 255, 0)
	gl.LineWidth(1)
	gl.Color3f(1, 0, 0)
	for !	tetris.Window.ShouldClose() {
		tetris.DrawScene()
		glfw.PollEvents()
	}
}
