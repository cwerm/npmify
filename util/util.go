package util

import (
	"bytes"
	"encoding/json"
	"github.com/Jeffail/gabs/v2"
	"github.com/imdario/mergo"
	"github.com/mcuadros/go-version"
	"io/ioutil"
	"npmify/fetch"
	"npmify/fs"
	"npmify/msg"
	"npmify/state"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var deps []state.Bower
var npmDeps []state.NpmDependency
var newPkgName string

func BuildDeps(data []byte, outputPath string) {

	jsonParsed, err := gabs.ParseJSON(data)
	msg.CheckErr(err)

	// Pull in all items under "dependencies"
	getKeys(jsonParsed, "dependencies")

	// Pull in all items under "devDependencies"
	getKeys(jsonParsed, "devDependencies")

	// Pull in all items under "resolutions"
	//getKeys(jsonParsed, "resolutions")

	d := &state.Dependencies{
		OutdatedCount: findOutdated(deps),
		TotalDependencies: getTotalCount(deps),
		Bower: deps,
	}

	fs.WriteNpmifyFile(outputPath, d)
}

func getKeys(jsonData *gabs.Container, bowerKey string) {

	msg.FancyPrint("Getting package names for bower %s\n", bowerKey)

	for key, child := range jsonData.S(bowerKey).ChildrenMap() {
		var b = state.Bower{}

		pkg := fetch.Get(`https://api.npms.io/v2/package/` + strings.ToLower(key))
		var re = regexp.MustCompile(`^(\~|\^)(.*)`)
		pkgJson, _ := gabs.ParseJSON(pkg)

		var bowerVersion = re.ReplaceAllString(child.Data().(string), "${2}")
		var npmVersion = strings.Trim(pkgJson.Path("collected.metadata.version").String(), "\"")
		var outdated = version.Compare(bowerVersion, npmVersion, "<")

		b.Name = key
		b.Version = bowerVersion
		b.NpmVersion = npmVersion
		b.License = strings.Trim(pkgJson.Path("collected.metadata.license").String(), "\"")
		b.Type = bowerKey
		b.Outdated = outdated

		if !IsVersionNumber(bowerVersion) {
			b.Group = "noBowerVersion"
		}
		if !IsVersionNumber(npmVersion) {
			b.Group = "noNpmVersion"
		}

		deps = append(deps, b)
	}
}

// IsVersionNumber takes a string, and runs Atoi on the first character.
// If the first character is a number (or if the version string is "*" or "latest",
// we return true.
func IsVersionNumber(version string) bool {
	firstChar := version[:1]

	if version == "*" || version == "latest" {
		return true
	}

	// There's a more terse way to do this, I just don't have the mental capacity
	// to figure it out right now.
	if _, err := strconv.Atoi(firstChar); err == nil {
		return true
	} else {
		return false
	}

}

func findOutdated(deps []state.Bower) int {
	var isOutdated []bool
	for _, dep := range deps {
		if dep.Outdated {
			isOutdated = append(isOutdated, dep.Outdated)
		}
	}
	return len(isOutdated)
}

func getTotalCount(deps []state.Bower) int {
	return len(deps)
}

func CopyPackageJson(cfg state.Configuration, newFilePath string) {
	originalPkgJson := cfg.PackageJsonPath
	// Save the package name for later.
	newPkgName = newFilePath
	nfp := filepath.Dir(originalPkgJson) + "/" + newFilePath
	err := fs.CopyFile(originalPkgJson, nfp)
	msg.CheckErr(err)
}

func BuildPkgJson(data []byte, cfg state.Configuration) {
	pkgJson, err := ioutil.ReadFile(cfg.PackageJsonPath)
	msg.CheckErr(err)

	npmParsed, err := gabs.ParseJSON(data)
	msg.CheckErr(err)

	for _, dep := range npmParsed.Path("bower").Children() {
		var d = state.NpmDependency{}
		d.PackageName = strings.Trim(dep.Path("name").String(), "\"")
		d.PackageVersion = strings.Trim(dep.Path("version").String(), "\"")

		npmDeps = append(npmDeps, d)
	}

	pkgData, err := gabs.ParseJSON(pkgJson)
	msg.CheckErr(err)

	// Build the package.json data into a struct
	pkgStruct := state.PackageJson{}
	pkgStruct.Name = strings.Trim(pkgData.S("name").String(), "\"")
	pkgStruct.Version = strings.Trim(pkgData.S("version").String(), "\"")
	pkgStruct.Dependencies = buildDependencyObject(npmDeps, pkgJson)
	pkgStruct.DevDependencies = buildDevDependencyObject(npmDeps, pkgJson)
	pkgStruct.Engines = buildPackageEngines(pkgJson)
	pkgStruct.Scripts = buildPackageScripts(pkgJson)

	fs.WritePackageJsonFile(filepath.Dir(cfg.PackageJsonPath) + "/" + "package.npmified.json", pkgStruct)
}

func buildDependencyObject(npmDependencies []state.NpmDependency, originalDependencies []byte) map[string]interface{} {
	var npm = make(map[string]interface{})
	var bower = make(map[string]interface{})

	og, err := gabs.ParseJSON(originalDependencies)
	msg.CheckErr(err)

	for _, dep := range npmDependencies {
		bower[dep.PackageName] = dep.PackageVersion
	}

	for name, ver := range og.S("dependencies").ChildrenMap() {
		npm[name] = strings.Trim(ver.String(), "\"")
	}

	err = mergo.Merge(&bower, npm)
	msg.CheckErr(err)

	return bower
}

func buildDevDependencyObject(npmDependencies []state.NpmDependency, originalDependencies []byte) map[string]interface{} {
	var npm = make(map[string]interface{})
	//var bower = make(map[string]interface{})

	og, err := gabs.ParseJSON(originalDependencies)
	msg.CheckErr(err)

	//for _, dep := range npmDependencies {
	//	bower[dep.PackageName] = dep.PackageVersion
	//}

	for name, ver := range og.S("devDependencies").ChildrenMap() {
		npm[name] = strings.Trim(ver.String(), "\"")
	}

	//err = mergo.Merge(&bower, npm)
	//msg.CheckErr(err)

	return npm
}

func buildPackageEngines(originalDependencies []byte) map[string]interface{} {
	var engines = make(map[string]interface{})

	og, err := gabs.ParseJSON(originalDependencies)
	msg.CheckErr(err)

	for key, val := range og.S("engines").ChildrenMap() {
		// Sorry...
		valString := strings.Replace(strings.Trim(val.String(), "\""), "u003e", ">", -1)
		engines[key] = strings.Trim(valString, "\\")
	}

	return engines
}

func buildPackageScripts(originalDependencies []byte) map[string]interface{} {
	var scripts = make(map[string]interface{})

	og, err := gabs.ParseJSON(originalDependencies)
	msg.CheckErr(err)

	for name, ver := range og.S("scripts").ChildrenMap() {
		scripts[name] = strings.Trim(ver.String(), "\"")
	}

	return scripts
}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
