package main

import (
	colour "image/color"
)

////////////////////////////////////////////////////////////////////////////////////////////
///////////// NEED TO CREATE A THEME CHOOSER BASED ON SETTINGS PREFERENCE //////////////////
////////////////////////////////////////////////////////////////////////////////////////////


// Changing the Main Window's Background Colour --> Dark Mode
func settingMainColourTheme() colourTheme {

	// Listed Theme Colours
	var basicDarkTheme colourTheme = colourTheme{
		BackgroundColour: colour.NRGBA{R: 34, G: 34, B: 34, A: 255},
		TextColour: colour.NRGBA{R: 220, G: 220, B: 220, A: 255}, 
		AccentColour: colour.NRGBA{R: 128, G: 128, B: 128, A: 255},
	}
	
	// Return the chosen theme
	return basicDarkTheme
}