package shortcut

import (
	"errors"
	"path/filepath"
	"strings"

	"imuslab.com/wdos/mod/filesystem/wdosfs"
	"imuslab.com/wdos/mod/utils"
)

/*
	A simple package to better handle shortcuts in WDOS

	Author: Secarian
*/

func ReadShortcut(shortcutContent []byte) (*wdosfs.ShortcutData, error) {
	//Split the content of the shortcut files into lines
	fileContent := strings.ReplaceAll(strings.TrimSpace(string(shortcutContent)), "\r\n", "\n")
	lines := strings.Split(fileContent, "\n")

	if len(lines) < 4 {
		return nil, errors.New("Corrupted Shortcut File")
	}

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
	}

	//Render it as shortcut data
	result := wdosfs.ShortcutData{
		Type: lines[0],
		Name: lines[1],
		Path: lines[2],
		Icon: lines[3],
	}

	return &result, nil
}

// Generate the content of a shortcut base the the four important field of shortcut information
func GenerateShortcutBytes(shortcutTarget string, shortcutType string, shortcutText string, shortcutIcon string) []byte {
	//Check if there are desktop icon. If yes, override icon on module
	if shortcutType == "module" && utils.FileExists(wdosfs.ToSlash(filepath.Join("./web/", filepath.Dir(shortcutIcon), "/desktop_icon.png"))) {
		shortcutIcon = wdosfs.ToSlash(filepath.Join(filepath.Dir(shortcutIcon), "/desktop_icon.png"))
	}

	//Clean the shortcut text
	shortcutText = wdosfs.FilterIllegalCharInFilename(shortcutText, " ")
	return []byte(shortcutType + "\n" + shortcutText + "\n" + shortcutTarget + "\n" + shortcutIcon)
}
