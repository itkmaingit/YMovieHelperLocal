package utils

import "strings"

func PathJoinForWindows(args ...string) string {
	return strings.Join(args, "\\")
}
