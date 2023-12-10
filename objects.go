package main

import (
	"fyne.io/fyne/v2"
	"path/filepath"
	"strings"
	"os"
)

// Importing all PNG icons as a Fyne resource (codable icon) --> Returns HashMap
func importExternalIcons() map[string]fyne.Resource {
	// Gathering the file location of the logfile
	var mainFolderDir string = findSyncNestFolderDir()
	var iconsLocation string = filepath.Join(mainFolderDir, "Icons")

	// Gather a list of all files in the directory
	var collectedIconFiles []string = listPNGFiles(iconsLocation)

	// Create a HashMap to hold the Fyne Resources (filenames and objects)
	var iconResourcesHashmap map[string]fyne.Resource = make(map[string]fyne.Resource)

	for _, iconFile := range collectedIconFiles {
		var fileLocation string = filepath.Join(iconsLocation, iconFile)
		iconResource := fyne.NewStaticResource(iconFile, readIconFile(fileLocation))
		iconResourcesHashmap[iconFile] = iconResource
	}

	// Return the hashmap that contains all of the icons --> Use the filename to access each
	return iconResourcesHashmap
}

// Returns a list of all PNG icons in any given directory
func listPNGFiles(iconsDirectoryPath string) []string {
	var directoryPNGFiles []string

	filesInDir, fileGrabErr := os.ReadDir(iconsDirectoryPath)
	if fileGrabErr != nil {
		sendingToLogfile("Cannot Read/Find Icons Directory")
	}

	for _, file := range filesInDir {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".png") {
			directoryPNGFiles = append(directoryPNGFiles, file.Name())
		}
	}
	
	if len(directoryPNGFiles) == 0 {
		sendingToLogfile("Unable to locate Icons")
	}
	
	return directoryPNGFiles
}

func readIconFile(filePath string) []byte {
	// Attempting to open the icon that was discovered
	content, err := os.ReadFile(filePath)
	if err != nil {
		sendingToLogfile("Unable to read an Icon File")
	}
	return content
}

// Returns the main directory location (required in case the application isn't run from its folder)
func findSyncNestFolderDir() string {
	// Gathering the location of the SyncNest .exe (the main SyncNest dir)
	exePath, _ := os.Executable() // Can't check for error --> No way to log

	// Getting the SyncNest folder --> Required in case the app is not started from the SyncNest folder
	var mainFolderDir string = filepath.Dir(exePath)

	return mainFolderDir
}