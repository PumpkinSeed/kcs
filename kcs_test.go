package kcs

import "testing"

func TestCheatSheetPrint(t *testing.T) {
	Data.Print("", "")
}

func TestCategory_Sort(t *testing.T) {
	configC := Data.Categories["config"]
	configC.Sort()
}
