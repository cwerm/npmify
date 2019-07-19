package fs

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	xlst "github.com/ivahaev/go-xlsx-templater"
	"io"
	"io/ioutil"
	"log"
	"npmify/msg"
	"npmify/state"
	"os"
	"sort"
)

func CopyFile(origPath string, newPath string) {
	orig, err := os.Open(origPath)
	if err != nil {
		log.Fatal(err)
	}
	defer orig.Close()

	newPkg, err := os.OpenFile(newPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer newPkg.Close()

	_, err = io.Copy(orig, newPkg)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteFile(filePath string, deps *state.Dependencies) {

	sort.Slice(deps.Bower, func(i, j int) bool {
		return deps.Bower[i].Name < deps.Bower[j].Name
	})

	DoExcel(*deps)

	file, err := json.MarshalIndent(&deps, "", "  ")
	msg.CheckErr(err)

	err = ioutil.WriteFile(filePath, file, 0644)
	msg.CheckErr(err)
}

func CreateDirectoryIfNotExist(dirName string) bool {
	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		msg.FancyPrint("Creating dir at %s\n", dirName)
		err := os.MkdirAll(dirName, 0755)
		if err != nil {
			log.Panic(err)
		}
		return true
	}

	if src.Mode().IsRegular() {
		fmt.Println(dirName, "already exist as a file!")
		return false
	}

	return false
}

func dependencyMap(d state.Dependencies) map[string]interface{} {
	return structs.Map(&d)
}

// DoExcel takes a struct of Dependencies and parses Dependencies.[]Bower to a spreadsheet.
func DoExcel(deps state.Dependencies) {
	var depMap = dependencyMap(deps)

	ctx := depMap

	doc := xlst.New()
	err := doc.ReadTemplate("./template.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	err = doc.Render(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = doc.Save("./report.xlsx")
	if err != nil {
		log.Fatal(err)
	}

}

func DirectoryExists(dirName string) bool {
	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		return false
	}

	if src.Mode().IsRegular() {
		fmt.Println(dirName, "already exist as a file!")
		return false
	}

	return true
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}