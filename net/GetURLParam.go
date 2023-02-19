package net

import "strings"

func GetURLParam(path string) (param string) {
	parts := strings.Split(path, "/")
	param = parts[len(parts)-1]
	return param
}