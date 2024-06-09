package components

import (
	"log"

	"github.com/breathxv/vtapp/src/errors"
	fs "github.com/fsnotify/fsnotify"
)

var (
	w *fs.Watcher
)

// TODO: Add better docs through function due to complexity
// Dir watcher takes a path (string) and watches
// that file/folder through the Windows Event
// Manager. It triggers a (go) loop which reads
// the events which then triggers the file checks.
// It can return err (error) which occurs when the
// function or, loop failed to operate successfully.
func DirWatcher(path string) (err error) {
	w, err := fs.NewWatcher()
	if err != nil {
		errors.ToastAlert()
		log.Print("Failed to create Watcher:", err)
		return err
	}
	defer w.Close()

	go watchLoop(w)

	err = w.Add(path)
	if err != nil {
		log.Fatalf("%q: %s", path, err)
		return err
	}

	log.Printf("ready; press ^C to exit")
	<-make(chan struct{}) // Block
	return nil
}

func watchLoop(w *fs.Watcher) {
	i := 0
	for {
		select {
		// Read from Errors.
		case err, ok := <-w.Errors:
			if !ok { // Watcher.Close() was called.
				return
			}
			log.Printf("ERROR: %s", err)
		// Read from Events.
		case e, ok := <-w.Events:
			if !ok { // Watcher.Close() was called.
				return
			}
			// Check for create event otherwise write event triggers multiple times
			if e.Has(fs.Create) {
				i++
				log.Printf("%3d", i)
				CheckFileSize(e.Name)
			} else {
				return
			}
		}
	}
}
