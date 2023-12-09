package main

import (
	colour "image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"
	// "fyne.io/fyne/v2/layout"
)

type colourTheme struct {
	BackgroundColour colour.NRGBA
	TextColour colour.NRGBA
	AccentColour colour.NRGBA
}

// Function that runs all scripts required at the start of the program
func setupMainWindow(mainApplication fyne.App, screenWidth int, screenHeight int) (fyne.Window, colourTheme) {
	// Creating a window with base sizing equal to half of the screen's resolution (this will later be changed to the last used sizing)
	var mainWindow fyne.Window = mainApplication.NewWindow("SyncNest")

	setPhysicalCharacteristicsMainWindow(&mainWindow, screenWidth, screenHeight)
	var baseTheme colourTheme = settingMainColourTheme()
	mainWindowLayoutSetup(&mainWindow, baseTheme)

	return mainWindow, baseTheme
}

// Creating the Main Window
func setPhysicalCharacteristicsMainWindow(ref_mainWindow *fyne.Window, screenWidth int, screenHeight int) {
	(*ref_mainWindow).Resize(fyne.NewSize(float32(screenWidth/2), float32(screenHeight/2)))
	(*ref_mainWindow).CenterOnScreen() // TODO: Settings in application to toggle this or "last known".
}

func mainWindowLayoutSetup(ref_mainWindow *fyne.Window, basicDarkTheme colourTheme, ) {
	// Set Main Window Background to the Chosen Theme Colour
	var mainWindowBackgroundColourBlock *canvas.Rectangle = canvas.NewRectangle(basicDarkTheme.BackgroundColour)

	//////////////////////////////////////////////
	////// TESTING START /////////////////////////
	//////////////////////////////////////////////

	// Create widgets or containers for the Toolbar (HBox)
	var toolbarTestButton *widget.Button = widget.NewButton("Button 1", func() {
		widget.ShowPopUp(widget.NewLabel("You clicked Button 1"), (*ref_mainWindow).Canvas())
	})

	// Main Windows Toolbar to be statically placed above all locations
	var mainWindowToolbar *widget.Toolbar = widget.NewToolbar(
		widget.NewToolbarAction(theme.SettingsIcon(), func() { toolbarSettingsMenuSetup(ref_mainWindow) }), // Used for Toolbar "Three-Dashes" Settings
		widget.NewToolbarAction(theme.FolderNewIcon(), toolbarTestButton.OnTapped), // Creates a new, empty folder in the current directory
		widget.NewToolbarAction(theme.FileIcon(), func() {  }), // Creates a new, empty file in the current directory
		widget.NewToolbarAction(theme.QuestionIcon(), func() {  }), // Opens a cmd prompt in the current directory
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.QuestionIcon(), func() {  }), // Collects all drives on the system and lists them as buttons
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.QuestionIcon(), func() {  }), // Compression of currently selected files (uses .7zip or windows compression)
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {  }), // Button to refresh the current directory (potentially all sub directories too)
		widget.NewToolbarAction(theme.QuestionIcon(), func() {  }), // Button used to open FTP connection settings
	)
	
	var mainWindowSplitContainer *container.Split = container.NewVSplit(
		mainWindowToolbar, // Top Toolbar widgets
		mainWindowBackgroundColourBlock, // Controlling Background Colour
	)

	// Controlling the Proportions between Widgets in the Container
	mainWindowSplitContainer.SetOffset(0.05)

	//////////////////////////////////////////////
	////// TESTING END ///////////////////////////
	//////////////////////////////////////////////

	// Tie the buttons and background to the main window
	(*ref_mainWindow).SetContent(mainWindowSplitContainer)
}

// Function used for defining the functionality of the "three-dashed" menu in the top left of the taskbar
func toolbarSettingsMenuSetup(ref_mainWindow *fyne.Window) {
	// Creating a Menu object that is to store all basic functionality
	var mainWindowToolbarDashedMenu *fyne.Menu = fyne.NewMenu("Options",
		fyne.NewMenuItem("Preferences", func() {
			widget.ShowPopUp(widget.NewLabel("Yo"), (*ref_mainWindow).Canvas())
		}),
	)

	var popupMenu *widget.PopUpMenu = widget.NewPopUpMenu(mainWindowToolbarDashedMenu, (*ref_mainWindow).Canvas())
	var mainApplicationDimensions fyne.Size = (*ref_mainWindow).Canvas().Size()

	popupMenu.ShowAtPosition(fyne.NewPos(10, 0.05*mainApplicationDimensions.Height)) // Positioned just below taskbar 
} 

func toolbarNewItemsMenuSetup(ref_mainWindow *fyne.Window) {





}