package ascii

import (
	"testing"
)

func TestLoadBanner(t *testing.T) {
	// Test loading a valid banner file
	loadBanner("shadow.txt")
	if _, ok := bannerMap["shadow.txt"]; !ok {
		t.Error("Failed to load banner file")
	}
	loadBanner("thinkertoy.txt")
	if _, ok := bannerMap["thinkertoy.txt"]; !ok {
		t.Error("Failed to load banner file")
	}
	loadBanner("standard.txt")
	if _, ok := bannerMap["standard.txt"]; !ok {
		t.Error("Failed to load the banner file")
	}

	// Add more test cases for loading other banner files or error scenarios
}

// func TestGetLetterArray(t *testing.T) {
// 	// Test getting the ASCII representation for a valid character
// 	expected := []string{
// 		"  _   _ ",
// 		" | | | |",
// 		" | |_| |",
// 		" |  _  |",
// 		" |_| |_|",
// 	}
// 	result := GetLetterArray('H', "standard.txt")
// 	if !equalSlices(expected, result) {
// 		t.Errorf("Expected %v, got %v", expected, result)
// 	}

// 	// Add more test cases for different characters and banner files
// }

// func TestPrintAscii(t *testing.T) {
// 	// Test printing ASCII art for a valid string
// 	expected := []string{
// 		"  _   _ ",
// 		" | | | |",
// 		" | |_| |",
// 		" |  _  |",
// 		" |_| |_|",
// 	}

// 	result := GetLetterArray('H', "standard.txt")
// 	if !equalSlices(expected, result) {
// 		t.Errorf("Expected %v, got %v", expected, result)
// 	}

// 	// Add more test cases for handling non-ASCII characters, empty input, modified file content, etc.
// }

// // Helper function to compare string slices
// func equalSlices(a, b []string) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}

// 	for i := range a {
// 		if i >= len(b) || a[i] != b[i] {
// 			return false
// 		}
// 	}

// 	return true
// }
