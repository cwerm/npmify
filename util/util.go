package util

import (
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	. "github.com/logrusorgru/aurora"
	"github.com/mcuadros/go-version"
	"io/ioutil"
	"log"
	"npmify/fetch"
	"regexp"
	"strconv"
	"strings"
)

type Dependencies struct {
	Bower 			[]Bower `json:"bower"`
	OutdatedCount	int		`json:"outdated_count"`
}

type Bower struct {
	Name 		string `json:"name"`
	Version 	string `json:"version"`
	NpmVersion 	string `json:"npm_version"`
	Type 		string `json:"type"`
	Outdated   	bool   `json:"outdated"`
}

var deps []Bower

func BuildDeps(data []byte, outputPath string) {

	jsonParsed, err := gabs.ParseJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	// Pull in all items under "dependencies"
	getKeys(jsonParsed, "dependencies")

	// Pull in all items under "devDependencies"
	getKeys(jsonParsed, "devDependencies")

	// Pull in all items under "resolutions"
	getKeys(jsonParsed, "resolutions")

	WriteFile(outputPath)
}

func getKeys(jsonData *gabs.Container, bowerKey string) {

	fmt.Printf(Sprintf(Blue("Getting keys for %s\n").Bold(), BrightWhite(bowerKey).Bold()))

	for key, child := range jsonData.S(bowerKey).ChildrenMap() {
		var b = Bower{}

		pkg := fetch.Get(`https://registry.npmjs.org/-/package/` + strings.ToLower(key) + `/dist-tags`)
		var re = regexp.MustCompile(`^(\~|\^)(.*)`)
		pkgJson, _ := gabs.ParseJSON(pkg)

		var bowerVersion = re.ReplaceAllString(child.Data().(string), "${2}")
		var npmVersion = strings.Trim(pkgJson.S("latest").String(), "\"")
		var outdated = version.Compare(bowerVersion, npmVersion, "<")

		b.Name = key
		b.Version = bowerVersion
		b.NpmVersion = npmVersion
		b.Type = bowerKey
		b.Outdated = outdated

		deps = append(deps, b)
	}

}

func newDeps(d []Bower) *Dependencies {
	return &Dependencies{
		OutdatedCount: findOutdated(d),
		Bower: d,
	}
}

func findOutdated(deps []Bower) int {
	var isOutdated []bool
	for _, dep := range deps {
		if dep.Outdated {
			isOutdated = append(isOutdated, dep.Outdated)
		}
	}

	FancyPrint("OUTDATED DEPENDENCIES: %s\n", strconv.Itoa(len(isOutdated)))

	return len(isOutdated)
}

func WriteFile(filePath string) {
	d := newDeps(deps)

	fmt.Println(Sprintf(Blue("Writing data to %s\n"), BrightWhite(filePath).Bold()))
	file, err := json.MarshalIndent(&d, "", "  ")
	CheckErr(err)

	err = ioutil.WriteFile(filePath, file, 0644)
	CheckErr(err)
}

func FancyPrint(format string, str string) {
	if str != "" {
		fmt.Printf(Sprintf(Blue(format).Bold(), BrightWhite(str).Bold()))
	}
}


func CheckErr(err error) {
	if err != nil {
		log.Panicf("ERROR: %s\n", Red(err))
	}
}
