package boxcars

type Sites map[string]Handlers

var	sites Sites

func SetupSites(config map[string]map[string]string) {
	newsites := make(Sites)

	for hostname, options := range config {
		debug("Setting up %s", hostname)
		newsites[hostname] = handlersOf(options)
	}

	sites = newsites
}
