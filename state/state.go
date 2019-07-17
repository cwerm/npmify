package state

type Configuration struct {
	OutputDir		string `json:"output_dir"`
	OutputFileName 	string `json:"output_file_name"`
	BowerFilePath	string `json:"bower_file_path"`
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
	Type 		string `json:"type"`
	Outdated   	bool   `json:"outdated"`
}
