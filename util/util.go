package util

import (
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	. "github.com/logrusorgru/aurora"
	"io/ioutil"
	"log"
	"npmify/fetch"
	"strings"
)

type Dependencies struct {
	Bower []Bower `json:"bower"`
}

type Bower struct {
	Name string `json:"name"`
	Version string `json:"version"`
	NpmVersion string `json:"npm_version"`
	//AnalyzedDate string `json:"analyzed_date"`
	Type string `json:"type"`
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

		pkgJson, _ := gabs.ParseJSON(pkg)

		b.Name = key
		b.Version = child.Data().(string)
		b.NpmVersion = strings.Trim(pkgJson.S("latest").String(), "\"")
		b.Type = bowerKey

		deps = append(deps, b)
	}

}

func newDeps(d []Bower) *Dependencies {
	return &Dependencies{
		Bower: d,
	}
}

func getBowerData(filePath string) {

}

func WriteFile(filePath string) {
	d := newDeps(deps)

	fmt.Println(Sprintf(Blue("Writing data to %s\n"), BrightWhite(filePath).Bold()))
	file, err := json.MarshalIndent(&d, "", "  ")
	CheckErr(err)

	err = ioutil.WriteFile(filePath, file, 0644)
	CheckErr(err)
}


func CheckErr(err error) {
	if err != nil {
		log.Panicf("ERROR: %s\n", Red(err))
	}
}
