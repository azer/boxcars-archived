package boxcars

import (
	"github.com/howeyc/fsnotify"
)

func AutoReload () {
	debug("Enabling")

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		debug("Failed to setup fsnotify")
		return
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				debug("%s has been updated. Event: %s", filename, ev)
				ReadConfig()
			case erv := <-watcher.Error:
				debug("Failed to monitor changes. Error: %s", erv)
			}
		}
	}()

	err = watcher.Watch(filename)
	if err != nil {
		debug("Failed to monitor changes on %s", filename)
	}

	<-done

	watcher.Close()
}
