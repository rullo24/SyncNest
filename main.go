package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fmt"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/layout"
	// "fyne.io/fyne/v2/theme"
	// "fyne.io/fyne/v2/widget"
)

func main() {
	// Creating application object
	var mainApplication fyne.App = app.New()
	
	// Using robotgo library
	var screenWidth, screenHeight int = grabScreenResolution()

	// Running Setup Scripts --> Returns the Main Window and Colour Theme as Objs
	var mainWindow fyne.Window
	var mainTheme colourTheme
	mainWindow, mainTheme = setupMainWindow(mainApplication, screenWidth, screenHeight)

	fmt.Println(mainTheme.BackgroundColour)

	










	// // Create a settings bar with your horizontal content
	// settingsBar := container.NewHBox(
	// 	widget.NewLabel("Temp_C:"),
	// 	widget.NewLabel("Temp_D:"),
	// 	widget.NewLabel("Temp_E"),
	// )

	// // Combine the settings bar and main content in a vertical layout
	// content := container.NewVBox(
	// 	container.New(layout.NewVBoxLayout(), settingsBar),
	// )

	// // Set the container size to 1/10th of the window for the settings bar
	// settingsBar.Resize(fyne.NewSize(mainWindow.Canvas().Size().Width/10, mainWindow.Canvas().Size().Height/10))

	// Set the content as the main content of the window
	// mainWindow.SetContent(content)

	mainWindow.ShowAndRun()
}
