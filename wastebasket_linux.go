// +build !windows !darwin

package wastebasket

import (
	"errors"
	"os/exec"
	"syscall"
)

func isCommandAvailable(name string) bool {
	cmd := exec.Command("bash", "-c", name)
	cmd.Start()

	err := cmd.Wait()
	if err == nil {
		return true
	}

	exitError, ok := err.(*exec.ExitError)
	if !ok {
		return false
	}

	status, ok := exitError.Sys().(syscall.WaitStatus)
	if ok {
		return status.ExitStatus() != 127
	}

	return true
}

//Trash moves a files or folder including its content into the systems trashbin.
func Trash(path string) error {
	//gio us the tool that replaces gvfs, therefore it is the first choice.
	if isCommandAvailable("gio") {
		return exec.Command("gio", "trash", "--force", path).Run()
	} else if isCommandAvailable("gvfs-trash") {
		return exec.Command("gvfs-trash", "--force", path).Run()
	}

	return errors.New("The commands `gio` and `gvfs-trash` aren't available")
}

//Empty clears the platforms trashbin.
func Empty() error {
	//gio us the tool that replaces gvfs, therefore it is the first choice.
	if isCommandAvailable("gio") {
		return exec.Command("gio", "trash", "--empty").Run()
	} else if isCommandAvailable("gvfs-trash") {
		return exec.Command("gvfs-trash", "--empty").Run()
	}

	return errors.New("The commands `gio` and `gvfs-trash` aren't available")
}
