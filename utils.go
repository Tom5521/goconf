package conf

import (
	"runtime"
)

func isUnix() bool {
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
}
