package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/layout"
	// "fyne.io/fyne/v2/theme"
	
)

func main() {
	// Creating application object
	var mainApplication fyne.App = app.New()
	
	// Using robotgo library
	var screenWidth, screenHeight int = grabScreenResolution()

	// Running Setup Scripts --> Returns the Main Window and Colour Theme as Objs
	var mainWindow fyne.Window
	var mainTheme colourTheme // Stores background, text and accent colours
	mainWindow, mainTheme = setupMainWindow(mainApplication, screenWidth, screenHeight)

	fmt.Println(mainTheme.BackgroundColour)

	mainWindow.ShowAndRun()
}
