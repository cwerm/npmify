package state

type Configuration struct {
	OutputDir		string `json:"output_dir"`
	OutputFileName 	string `json:"output_file_name"`
	BowerFilePath	string `json:"bower_file_path"`
	PackageJsonPath string `json:"package_json_path"`
	Packages		[]Package `json:"packages"`
	Version			string	`json:"version"`
}

type Package struct {
	Name string `json:"name"`
	NpmVersion string `json:"npm_version"`
	NpmPackage string `json:"npm_package"`
}

type Dependencies struct {
	Bower 				[]Bower `json:"bower"`
	OutdatedCount		int		`json:"outdated_count"`
	TotalDependencies 	int 	`json:"total_dependencies"`
}

type Bower struct {
	Name 		string `json:"name"`
	Version 	string `json:"version"`
	NpmVersion 	string `json:"npm_version"`
	License		string	`json:"license"`
	Type 		string `json:"type"`
	Outdated   	bool   `json:"outdated"`
	Group		string	`json:"group"`
}

type PackageJson struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Dependencies map[string]interface{} `json:"dependencies"`
	DevDependencies map[string]interface{} `json:"devDependencies"`
	Engines map[string]interface{} `json:"engines"`
	Scripts map[string]interface{} `json:"scripts"`
}

type NpmDependency struct {
	PackageName 	string `json:"package_name"`
	PackageVersion  string `json:"package_version"`
}