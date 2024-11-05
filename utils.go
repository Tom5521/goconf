package conf

import (
	"runtime"
)

var IsUnix = func() bool {
	unixes := [...]string{
		"netbsd",
		"linux",
		"openbsd",
		"darwin",
	}

	for _, o := range unixes {
		if runtime.GOOS == o {
			return true
		}
	}
	return false
}()
