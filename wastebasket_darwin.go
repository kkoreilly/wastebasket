// +build darwin

package wastebasket

import (
	"os"
	"os/exec"
)

//Trash moves a files or folder including its content into the systems trashbin.
func Trash(path string) error {
	_, fileError := os.Stat(path)

	if os.IsNotExist(fileError) {
		return nil
	}

	if fileError != nil {
		return fileError
	}

	return exec.Command("trash", path).Run()
}

//Empty clears the platforms trashbin. It uses the `Finder` app to empty the trashbin.
func Empty() error {
	return exec.Command("osascript", "-e", "tell app \"Finder\" to empty").Run()
}
