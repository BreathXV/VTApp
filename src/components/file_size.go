package components

// Checks the size of the file/folder
import (
	"log"
	"os"
)

func CheckFileSize(path string) error {
	file, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err)
	}

	if file.Size() < 32000000 {
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
