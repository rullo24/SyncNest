package main

import (
	colour "image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/theme"
	// "fyne.io/fyne/v2/layout"
)

type colourTheme struct {
	BackgroundColour colour.NRGBA
	TextColour colour.NRGBA
	AccentColour colour.NRGBA
}

// Function that runs all scripts required at the start of the program
func setupMainWindow(mainApplication fyne.App, screenWidth int, screenHeight int) fyne.Window {
	// Creating a window with base sizing equal to half of the screen's resolution (this will later be changed to the last used sizing)
	var mainWindow fyne.Window = mainApplication.NewWindow("SyncNest")

	// Importing external objects for use in setup
	var iconsHashMap map[string]fyne.Resource = importExternalIcons()

	setPhysicalCharacteristicsMainWindow(&mainWindow, screenWidth, screenHeight)
	var baseTheme colourTheme = settingMainColourTheme()
	mainWindowLayoutSetup(&mainWindow, baseTheme, iconsHashMap)

	return mainWindow
}

// Creating the Main Window
func setPhysicalCharacteristicsMainWindow(ref_mainWindow *fyne.Window, screenWidth int, screenHeight int) {
	(*ref_mainWindow).Resize(fyne.NewSize(float32(screenWidth/2), float32(screenHeight/2)))
	(*ref_mainWindow).CenterOnScreen() // TODO: Settings in application to toggle this or "last known".
}

func mainWindowLayoutSetup(ref_mainWindow *fyne.Window, basicDarkTheme colourTheme, iconsHashMap map[string]fyne.Resource) {
	// Set Main Window Background to the Chosen Theme Colour
	var mainWindowBackgroundColourBlock *canvas.Rectangle = canvas.NewRectangle(basicDarkTheme.BackgroundColour)

	// Main Windows Toolbar to be statically placed above all locations
	var mainWindowToolbar *widget.Toolbar = createMainToolbar(ref_mainWindow, iconsHashMap)
	
	// Creating a visible split on the main window
	var mainWindowSplitContainer *container.Split = container.NewVSplit(
		mainWindowToolbar, // Top Toolbar widgets
		mainWindowBackgroundColourBlock, // Controlling Background Colour
	)

	// Controlling the Proportions between Widgets in the Container
	mainWindowSplitContainer.SetOffset(0.05)

	// Tie the buttons and background to the main window
	(*ref_mainWindow).SetContent(mainWindowSplitContainer)
}