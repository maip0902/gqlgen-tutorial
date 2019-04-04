package main

import (
"flag"

"github.com/sweetbrain/gqlgen-tutorial"
)

func main() {
	var (
		addr   = flag.String("addr", ":8080", "addr to bind")
		dbconf = flag.String("dbconf", "dbconfig.yml", "database configuration file.")
		env    = flag.String("env", "development", "application envirionment (production, development etc.)")
		debug  = flag.Bool("debug", false, "debug mode. default is false.")
	)
	flag.Parse()
	b := mgqlgen.New()
	b.Init(*dbconf, *env, *debug)
	b.Run(*addr)
}
