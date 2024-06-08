package components

// Checks the size of the file/folder
import (
	"fmt"
	"log"
	"os"

	toast "github.com/gen2brain/beeep"
)

// Check file size takes a path (string) for a
// file/zipped folder. It then checks the size
// of the file and invokes the appropriate
// handler (e.g. reg/large). It can return err
// (error) which occurs when the function could
// not complete successfully.
func CheckFileSize(path string) error {
	file, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	if file.Size() < 32000000 {
		msg := fmt.Sprintf("Processing file now...\n%s", file.Name())
		toast.Notify("VTApp", msg, "path/to/icon")
		return nil
		// Reg File
	} else if file.Size() > 32000000 && file.Size() < 650000000 {
		return nil
		// Large File
	} else {
		err := "File to large to process!"
		toast.Alert("VTApp", "File to large to process!", "path/to/icon")
		return fmt.Errorf(err)
		// Return + Fatal - file to large
	}
}
