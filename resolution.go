package main

import (
	"github.com/go-vgo/robotgo"
)

func grabScreenResolution() (int, int) {
	screenWidth, screenHeight := robotgo.GetScreenSize()
	return screenWidth, screenHeight
}