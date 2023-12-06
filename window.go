package main

import (
	"fyne.io/fyne/v2"
	colour "image/color"
	"fyne.io/fyne/v2/canvas"
)

type colourTheme struct {
	BackgroundColour colour.NRGBA
	TextColour colour.NRGBA
	AccentColour colour.NRGBA
}

// Function that runs all scripts required at the start of the program
func setupMainWindow(mainApplication fyne.App, screenWidth int, screenHeight int) (fyne.Window, colourTheme) {
	// Creating a window with base sizing equal to half of the screen's resolution (this will later be changed to the last used sizing)
	mainWindow := mainApplication.NewWindow("SyncNest")

	setPhysicalCharacteristicsMainWindow(&mainWindow, screenWidth, screenHeight)
	var baseTheme colourTheme = settingColourMainWindow(&mainWindow)

	return mainWindow, baseTheme
}

// Creating the Main Window
func setPhysicalCharacteristicsMainWindow(ref_mainWindow *fyne.Window, screenWidth int, screenHeight int) {
	(*ref_mainWindow).Resize(fyne.NewSize(float32(screenWidth/2), float32(screenHeight/2)))
	(*ref_mainWindow).CenterOnScreen() // TODO: Settings in application to toggle this or "last known".
}

// Changing the Main Window's Background Colour --> Dark Mode
func settingColourMainWindow(ref_mainWindow *fyne.Window) colourTheme {

	// Listed Theme Colours
	var basicDarkTheme colourTheme = colourTheme{
		BackgroundColour: colour.NRGBA{R: 34, G: 34, B: 34, A: 255},
		TextColour: colour.NRGBA{R: 220, G: 220, B: 220, A: 255}, 
		AccentColour: colour.NRGBA{R: 128, G: 128, B: 128, A: 255},
	}

	// Set Main Window Background to the Chosen Theme Colour
	var mainWindowBackgroundCanvasObject *canvas.Rectangle = canvas.NewRectangle(basicDarkTheme.BackgroundColour)
	(*ref_mainWindow).SetContent(mainWindowBackgroundCanvasObject)

	// Return the chosen theme
	return basicDarkTheme
}