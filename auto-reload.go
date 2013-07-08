package boxcars

import (
	"github.com/howeyc/fsnotify"
)

func AutoReload () {
	debug("Enabling")

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		debug("Failed to setup fsnotify")
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				debug("%s has been updated. Event: %s", config, ev)
				Load(config)
			case erv := <-watcher.Error:
				debug("Failed to monitor changes. Error: %s", erv)
			}
		}
	}()

	err = watcher.Watch(config)
	if err != nil {
		debug("Failed to monitor changes on %s", config)
	}

	<-done

	watcher.Close()
}
