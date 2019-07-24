package util

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/mcuadros/go-version"
	"npmify/fetch"
	"npmify/fs"
	"npmify/msg"
	"npmify/state"
	"regexp"
	"strconv"
	"strings"
)

var deps []state.Bower

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

	fs.WriteFile(outputPath, d)
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
