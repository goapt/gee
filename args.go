package gee

import (
	"os"
)

var ExtCliArgs map[string]string
var CliArgs []string

func ArgsInit() {
	ExtCliArgs = make(map[string]string)
	CliArgs = parseArgs(os.Args, "config", "debug", "show-sql", "debug-tag", "env")
}

func parseArgs(args []string, excludes ...string) []string {

	rs := make([]string, 0)
	for _, arg := range args[1:] {
		isFind := false
		for _, ext := range excludes {
			prefix := "--" + ext + "="
			lenPrefix := len(prefix)
			if len(arg) > lenPrefix && prefix == arg[0:lenPrefix] {
				isFind = true
				ExtCliArgs[ext] = arg[lenPrefix:]
			}
		}

		if !isFind {
			rs = append(rs, arg)
		}
	}
	return rs
}
