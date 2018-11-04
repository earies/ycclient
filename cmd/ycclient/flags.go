package main

import (
	"flag"
)

var (
	base         = flag.String("base", "", "Base URL for YANG catalog")
	config       = flag.String("config", "", "Configuration file")
	contributors = flag.Bool("contributors", false, "Retrieve contributing entities to the YANG catalog")
	getFlag      = flag.Bool("get", false, "Perform REST GET request to YANG catalog")
	module       = flag.String("module", "", "Retrieve module metadata")
	org          = flag.String("org", "", "Organization to retrieve")
	revision     = flag.String("revision", "", "Module revision date")
	vendor       = flag.String("vendor", "", "Vendor to retrieve")
	verbose      = flag.Bool("verbose", false, "Enable verbose output")
	versionFlag  = flag.Bool("version", false, "Build information")
)

func initFlags() {
	flag.Parse()
}
