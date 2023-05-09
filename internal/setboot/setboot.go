package setboot

import (
	"os"

	"github.com/emersion/go-autostart"
)

// It register Prox Client to start with the OS(windows)
func SetStartWithOS() error {
	path, err := os.Executable()
	if err != nil {
		return err
	}

	start := &autostart.App{
		Name:        "Prox",
		Exec:        []string{path, "-c", "echo autostart >> ~/autostart.txt"},
		DisplayName: "Prox",
	}

	if !start.IsEnabled() {
		if err := start.Enable(); err != nil {
			return err
		}
	}

	return nil
}
