package schema

type SearchFilterImplementation struct {
	Vendor          string `json:"vendor"`
	SoftwareVersion string `json:"software-version"`
}

type SearchFilterImplementations struct {
	Implementation []SearchFilterImplementation `json:"implementation"`
}

type SearchFilterInput struct {
	Implementations SearchFilterImplementations `json:"implementations"`
}

type SearchFilter struct {
	Input SearchFilterInput `json:"input"`
}

type Deviation struct {
	Name     string `json:"name"`
	Revision string `json:"revision"`
}

type Implementation struct {
	Vendor          string      `json:"vendor"`
	Platform        string      `json:"platform"`
	SoftwareVersion string      `json:"software-version"`
	SoftwareFlavor  string      `json:"software-flavor"`
	OsVersion       string      `json:"os-version"`
	FeatureSet      string      `json:"feature-set"`
	OsType          string      `json:"os-type"`
	Feature         string      `json:"feature"`
	Deviation       []Deviation `json:"deviation"`
	ConformanceType string      `json:"conformance-type"`
}

type Implementations struct {
	Implementation []Implementation `json:"implementation"`
}

type ModuleAttributes struct {
	Name     string `json:"name"`
	Revision string `json:"revision"`
	Schema   string `json:"schema"`
}

type ModuleResultModule struct {
	Name          string `json:"name"`
	Revision      string `json:"revision"`
	Organization  string `json:"organization"`
	Namespace     string `json:"namespace"`
	Schema        string `json:"schema"`
	GeneratedFrom string `json:"generated-from"`
	//MaturityLevel string `json:"maturity-level"`
	//DocumentName  string `json:"document-name"`
	//AuthorEmail   string `json:"author-email"`
	//Reference     string `json:"reference"`
	ModuleClassification string `json:"module-classification"`
	CompilationStatus    string `json:"compilation-status"`
	CompilationResult    string `json:"compilation-result"`
	Prefix               string `json:"prefix"`
	YangVersion          string `json:"yang-version"`
	Description          string `json:"description"`
	Contact              string `json:"contact"`
	ModuleType           string `json:"module-type"`
	//BelongsTo            string `json:"belongs-to"`
	TreeType string `json:"tree-type"`
	//YangTree string `json:"yang-tree"`
	//Expires string `json:"expires"`
	Expired ExpiredType `json:"expired"`
	//SubModule    []ModuleAttributes `json:"submodule"`
	Dependencies []ModuleAttributes `json:"dependencies"`
	//Dependents   []ModuleAttributes `json:"dependents"`
	SemanticVersion        string          `json:"semantic-version"`
	DerivedSemanticVersion string          `json:"derived-semantic-version"`
	Implementations        Implementations `json:"implementations"`
}

type ModuleResult struct {
	Module []ModuleResultModule `json:"module"`
}

type Dependencies struct {
	Name string `json:"name"`
}

type AuthorEmailInput struct {
	Dependencies []Dependencies `json:"dependencies"`
}

type AuthorEmailFilter struct {
	Input AuthorEmailInput `json:"input"`
}

type ContributorsResult struct {
	Contributors []string `json:"contributors"`
}

type ModelKeys struct {
	Name         string `json:"name"`
	Revision     string `json:"revision"`
	Organization string `json:"organization"`
}
