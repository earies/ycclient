package main

import (
	"log"
	"os"
	"text/template"

	"github.com/earies/ycclient/pkg/schema"
)

const contributorTemplate = `Contributors:
----------------------------------
{{range .}}  Name: {{.}}
{{end}}`

func renderContributors(c []string) {
	report, err := template.New("contributors").Funcs(template.FuncMap{}).Parse(contributorTemplate)
	if err != nil {
		log.Fatal(err)
	}
	err = report.Execute(os.Stdout, c)
	if err != nil {
		log.Fatal(err)
	}
}

const versionTemplate = `Version: {{.Version}}
BuildTime: {{.BuildTime}}
`

func renderVersion(p PackageInfo) {
	report, err := template.New("package").Funcs(template.FuncMap{}).Parse(versionTemplate)
	if err != nil {
		log.Fatal(err)
	}
	err = report.Execute(os.Stdout, p)
	if err != nil {
		log.Fatal(err)
	}
}

const moduleTemplate = `{{range .Module}}Name: {{.Name}}
Revision: {{.Revision}}
Organization: {{.Organization}}
===============================
  Namespace: {{.Namespace}}
  Schema: {{.Schema}}
  Generated From: {{.GeneratedFrom}}
  Module Classification: {{.ModuleClassification}}
  Compilation Status: {{.CompilationStatus}}
  Compilation Result: {{.CompilationResult}}
  Prefix: {{.Prefix}}
  YANG Version: {{.YangVersion}}
  Description: {{.Description}}
  Contact: {{.Contact}}
  Module Type: {{.ModuleType}}
  Tree Type: {{.TreeType}}
  Expired: {{.Expired}}
  Dependencies: {{range .Dependencies}}
    Module Name: {{.Name}}
      Revision: {{.Revision}}
      Schema: {{.Schema}}{{end}}
  Semantic Version: {{.SemanticVersion}}
  Derived Semantic Version: {{.DerivedSemanticVersion}}
  Implementations: {{range .Implementations.Implementation}}
    Vendor: {{.Vendor}}
    Platform: {{.Platform}}
    Software Version: {{.SoftwareVersion}}
    Software Flavor: {{.SoftwareFlavor}}
    OS Version: {{.OsVersion}}
    Feature Set: {{.FeatureSet}}
    OS Type: {{.OsType}}
    Feature: {{.Feature}}
    Deviation: {{range .Deviation}}
      Name: {{.Name}}
      Revision: {{.Revision}}{{end}}
    Conformance Type: {{.ConformanceType}}
    ----{{end}}
{{end}}`

// Can you generalize all render functions
func renderModule(m schema.ModuleResult) {
	report, err := template.New("module").Funcs(template.FuncMap{}).Parse(moduleTemplate)
	if err != nil {
		log.Fatal(err)
	}
	err = report.Execute(os.Stdout, m)
	if err != nil {
		log.Fatal(err)
	}
}
