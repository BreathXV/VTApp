package components

// Checks the size of the file/folder
import (
	"fmt"
	"log"
	"os"

	ntfy "github.com/gen2brain/beeep"
)

func CheckFileSize(path string) error {
	file, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	if file.Size() < 32000000 {
		msg := fmt.Sprintf("Processing file now...\n%s", file.Name())
		ntfy.Notify("VTApp", msg, "path/to/icon")
		return nil
		// Reg File
	} else if file.Size() > 32000000 && file.Size() < 650000000 {
		return nil
		// Large File
	} else {
		return nil
		// Return + Fatal - file to large
	}
}
