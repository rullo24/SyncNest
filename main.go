package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	// Creating application object
	var mainApplication fyne.App = app.New()

	// Using robotgo library to gather resolution stats
	var screenWidth, screenHeight int = grabScreenResolution()

	// Running Setup Scripts --> Returns the Main Window and Colour Theme as Objs
	var mainWindow fyne.Window = setupMainWindow(mainApplication, screenWidth, screenHeight)

	mainWindow.ShowAndRun()
}
