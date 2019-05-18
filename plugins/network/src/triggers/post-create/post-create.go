package main

import (
	"flag"

	"github.com/dokku/dokku/plugins/common"
)

// set bind-all-interfaces to false by default
func main() {
	flag.Parse()
	appName := flag.Arg(0)

	err := common.PropertyWrite("network", appName, "bind-all-interfaces", "false")
	if err != nil {
		common.LogWarn(err.Error())
	}
}
