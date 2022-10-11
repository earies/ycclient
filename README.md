# ycclient

A tool to fetch details of a specific module from YANG Catalog.

```
Usage of ycclient:
  -base string
    	Base URL for YANG catalog
  -config string
    	Configuration file
  -contributors
    	Retrieve contributing entities to the YANG catalog
  -get
    	Perform REST GET request to YANG catalog
  -module string
    	Retrieve module metadata
  -org string
    	Organization to retrieve
  -revision string
    	Module revision date
  -vendor string
    	Vendor to retrieve
  -verbose
    	Enable verbose output
  -version
    	Build information
```

The base url to YANG Catalog must be either passed with the `--base` argument, or specified in a toml config file like so:
```
[base]
url = "https://yangcatalog.org"
```

### Examples
Get a list of organizations that have contributed modules to YANG Catalog
```
ycclient --base https://yangcatalog.org --get --contributors
```
Fetch the details of the `Cisco-NX-OS-device@2022-08-18` YANG module
```
ycclient --base https://yangcatalog.org --get --module Cisco-NX-OS-device --revision 2022-08-18 --org cisco --vendor cisco
```

### Compatibility
ycclient fetches data from the following YANG Catalog endpoints:
- `api/contributors`
- `api/search/modules/<name>,<revision>,<organization>`

These endpoints are covered by tests on YANG Catalog's side and guaranteed to maintain compatibility.