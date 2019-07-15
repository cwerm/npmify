package util

import (
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"io/ioutil"
	"log"
	"npmify/fetch"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Dependencies struct {
	Bower []Bower `json:"bower"`
}

type Bower struct {
	Name string `json:"name"`
	Version string `json:"version"`
	NpmVersion string `json:"npm_version"`
	AnalyzedDate string `json:"analyzed_date"`
	Type string `json:"type"`
}

var deps []Bower

func BuildDeps(data []byte) {

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

	WriteFile("bower_deps.json")
}

func getKeys(jsonData *gabs.Container, bowerKey string) {

	fmt.Printf("Getting keys for %s\n", bowerKey)

	for key, child := range jsonData.S(bowerKey).ChildrenMap() {
		var b = Bower{}

		pkg := fetch.Get(`https://api.npms.io/v2/package/` + strings.ToLower(key))

		pkgJson, _ := gabs.ParseJSON(pkg)

		b.Name = key
		b.Version = child.Data().(string)
		b.AnalyzedDate = strings.Trim(pkgJson.S("analyzedAt").String(), "\"")
		b.NpmVersion = strings.Trim(pkgJson.S("collected", "metadata", "version").String(), "\"")
		b.Type = bowerKey

		deps = append(deps, b)
	}

}

func newDeps(d []Bower) *Dependencies {
	return &Dependencies{
		Bower: d,
	}
}

func WriteFile(fileName string) {

	d := newDeps(deps)

	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	dataDir := filepath.Join(usr.HomeDir, "dependencies")
	err = os.MkdirAll(dataDir, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	var filePath = dataDir + "/" + fileName
	fmt.Printf("Writing %s to %s\n", d, filePath)
	file, _ := json.MarshalIndent(&d, "", "  ")
	err = ioutil.WriteFile(filePath, file, 0644)
	if err != nil {
		log.Fatalln(err)
	}

}
