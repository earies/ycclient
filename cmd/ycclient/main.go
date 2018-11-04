package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/earies/ycclient/pkg/schema"
	"github.com/earies/ycclient/pkg/transport/rest"
)

func main() {
	initFlags()

	if *versionFlag {
		p := PackageInfo{
			Version:   Version,
			BuildTime: BuildTime,
		}
		renderVersion(p)
		os.Exit(0)
	}

	var conf Config
	if *config != "" && *base == "" {
		initConfig(*config, &conf)
	} else if *base != "" {
		conf.Base.Url = *base
	} else {
		log.Fatalf("Must specify configuration file or --base flag")
	}

	switch {
	case *getFlag:
		u, err := url.Parse(conf.Base.Url)
		if err != nil {
			log.Fatal(err)
		}
		h := http.Client{}
		c := rest.Client{
			BaseURL:    u,
			HTTPClient: &h,
		}

		switch {
		case *contributors:
			contributors, err := c.Contributors()
			if err != nil {
				log.Fatal(err)
			}
			renderContributors(contributors)
		case *module != "":
			if *revision == "" {
				log.Fatalf("Must specify --revision flag with module")
			}
			if *org == "" {
				log.Fatalf("Must specify --org flag with module")
			}
			if *vendor == "" {
				log.Fatalf("Must specify --vendor flag with module")
			}
			var filter schema.SearchFilter
			var impl schema.SearchFilterImplementation
			impl.Vendor = *vendor
			filter.Input.Implementations.Implementation = append(filter.Input.Implementations.Implementation, impl)

			var keys schema.ModelKeys
			keys.Name = *module
			keys.Revision = *revision
			keys.Organization = *org

			module, err := c.Module(keys, filter)
			if err != nil {
				log.Fatal(err)
			}
			renderModule(module)
		default:
			log.Fatalf("Must specify one of: [contributors] with action '-get'")
		}
	default:
		log.Fatalf("Must specify at least 1 option\n")
	}
}
