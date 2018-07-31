// sc2lover is a simple discussion board / forum web application.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	logger       = log.New(os.Stdout, "", log.LstdFlags|log.LUTC)
	useEnvConfig = flag.Bool("e", false, "use environment variables as config")
)

func main() {
	flag.Usage = help
	flag.Parse()

	cmds := map[string]func(){
		"start":        startServer,
		"init":         initConfig,
		"gen-key":      genKey,
		"admins":       printAdmins,
		"add-admin":    addAdmin,
		"remove-admin": removeAdmin,
		"help":         help,
	}

	if cmdFunc, ok := cmds[flag.Arg(0)]; ok {
		cmdFunc()
	} else {
		help()
		os.Exit(2)
	}
}

func help() {
	fmt.Fprintln(os.Stderr, `Usage:
	sc2lover start                      - start the server
	sc2lover init                       - create an initial configuration file
	sc2lover gen-key                    - generate a random 32-byte hex-encoded key
	sc2lover admins                     - show the admin list
	sc2lover add-admin <username>       - add a user to the admin list
	sc2lover remove-admin <username>    - remove a user from the admin list
	sc2lover help                       - show this message
Use -e flag to read configuration from environment variables instead of a file. E.g.:
	sc2lover -e start
	sc2lover -e admins`)
}
