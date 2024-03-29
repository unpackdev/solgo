package utils

import (
	"runtime/debug"
	"strings"
)

// GetBuildVersionByModule retrieves the build version of a specified module from
// the current application's build information. It parses the build information
// obtained from the runtime debug package to search for a specific module.
func GetBuildVersionByModule(module string) string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	for _, dep := range info.Deps {
		if dep.Path == module {
			return strings.Replace(dep.Version, "v", "", -1)
		}
	}

	return "unknown"
}
