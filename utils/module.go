package utils

import (
	"runtime/debug"
	"strings"
)

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
