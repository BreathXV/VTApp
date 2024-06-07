package components

import (
	"log"

	fs "github.com/fsnotify/fsnotify"
)

// TODO: Import fsnotify

func DirWatcher(path string) (file_path string, err error) {
	w, err := fs.NewWatcher()
	if err != nil {
		ToastAlert()
		log.Print("Failed to create Watcher:", err)
		return "", err
	}
	defer w.Close()

	go watchLoop(w)

	err = w.Add(path)
	if err != nil {
		log.Fatalf("%q: %s", path, err)
	}

	log.Printf("ready; press ^C to exit")
	<-make(chan struct{}) // Block forever
	return "", nil
}

func watchLoop(w *fs.Watcher) {
	i := 0
	for {
		select {
		// Read from Errors.
		case err, ok := <-w.Errors:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				return
			}
			log.Printf("ERROR: %s", err)
		// Read from Events.
		case e, ok := <-w.Events:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				return
			}

			// Just print the event nicely aligned, and keep track how many
			// events we've seen.
			i++
			log.Printf("%3d", i)
			CheckFileSize(e.Name)
		}
	}
}
