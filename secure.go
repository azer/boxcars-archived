// +build linux, darwin

package boxcars

import (
	"syscall"
)

func Secure(uid, gid int) {
	debug("Enabled.")

	if syscall.Getuid() != 0 {
		debug("Not running as sudo.")
		return
	}

	// set group id first as you need to be root to change group
	debug("Setting GID (group id) to %v", gid)
	err := syscall.Setgid(gid)
	if err != nil {
		debug("Fatal: Failed to set group id: %v", err)
		debug("  Use -gid parameter to specify the correct group id.")
	}

	debug("Setting UID (user id) to %v", gid)
	err = syscall.Setuid(uid)
	if err != nil {
		debug("Fatal: Failed to set user id: %v", err)
		debug("  Use -uid parameter to specify the correct user id.")
	}
}
