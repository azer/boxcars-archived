package boxcars

type Sites map[string]Handlers

var (
	sites Sites
)

func SetupSites (config Config) {
	sites = make(Sites)

	for hostname, options := range config {
		debug("Setting up %s", hostname)
		sites[hostname] = handlersOf(options)
	}
}
