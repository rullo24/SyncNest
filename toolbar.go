package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"
)

// Designing toolbar
func createMainToolbar(ref_mainWindow *fyne.Window, iconsHashMap map[string]fyne.Resource) *widget.Toolbar {
	// Main Windows Toolbar to be statically placed above all locations
	var mainWindowToolbar *widget.Toolbar = widget.NewToolbar(
		widget.NewToolbarAction(iconsHashMap["Menu.png"], func() { toolbarMainToolbarMenuSetup(ref_mainWindow) }),
		widget.NewToolbarAction(iconsHashMap["NewFolder.png"], func() {  }), // Creates a new, empty folder in the current directory
		widget.NewToolbarAction(iconsHashMap["AddNewDoc.png"], func() {  }), // Creates a new, empty file in the current directory
		widget.NewToolbarAction(iconsHashMap["SimpleTerminal.png"], func() {  }), // Opens a cmd prompt in the current directory
		widget.NewToolbarAction(iconsHashMap["Compression.png"], func() {  }), // Compression of currently selected files (uses .7zip or windows compression)
		widget.NewToolbarAction(iconsHashMap["Refresh.png"], func() {  }), // Button to refresh the current directory (potentially all sub directories too)
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.QuestionIcon(), func() {  }), // Collects all drives on the system and lists them as buttons
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(iconsHashMap["FTP.png"], func() {  }), // Button used to open FTP connection settings
		widget.NewToolbarAction(iconsHashMap["Settings.png"], func() {  }), // Used for Preference Settings
	)

	return mainWindowToolbar
}

// Function used for defining the functionality of the "three-dashed" menu in the top left of the taskbar
func toolbarMainToolbarMenuSetup(ref_mainWindow *fyne.Window) {
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

// Buttons used for creating an empty folder in the current dir
func toolbarCreateFolderInCurrentDir() {

}

// Creating an empty file in the current dir
func toolbarCreateEmptyFileInCurrentDir() {

}

// Opening Terminal (cmd or otherwise chosen in settings) --> Default cmd
func toolbarOpenTerminalCurrentDir() {

}

// Compressing selected files (defaults to using 7zip via popup). If 7zip N/A, uses base windows compression software.
func toolbarCompressSelected() {

}

// Force a refresh of the dir's files and folders
func toolbarForceFandFRefresh() {

}

// Displaying all available computer drives (as buttons) --> takes to root location of drive
func toolbarGatherComputerDrives() {

}

// Open FTP settings
func toolbarOpenFTPSettings() {

}

// Open SyncNest General Settings
func toolbarOpenMainSettings() {

}