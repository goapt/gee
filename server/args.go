package server

import (
	"os"
)

var ExtArgs map[string]string
var Args []string

func ArgsInit() {
	ExtArgs = make(map[string]string)
	Args = ParseArgs("config", "debug", "show-sql", "debug-tag")
}

func ParseArgs(excludes ...string) []string {
	rs := make([]string, 0)
	for _, arg := range os.Args[1:] {
		isFind := false
		for _, ext := range excludes {
			prefix := "--" + ext + "="
			lenPrefix := len(prefix)
			if len(arg) > lenPrefix && prefix == arg[0:lenPrefix] {
				isFind = true
				ExtArgs[ext] = arg[lenPrefix:]
			}
		}

		if !isFind {
			rs = append(rs, arg)
		}
	}
	return rs
}
