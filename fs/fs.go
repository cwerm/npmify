package fs

import (
	"bytes"
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

func CopyFile(src, dst string) error {
	var err error
	var srcfile *os.File
	var dstfile *os.File
	var srcinfo os.FileInfo

	if srcfile, err = os.Open(src); err != nil {
		return err
	}
	defer srcfile.Close()

	if dstfile, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfile.Close()

	if _, err = io.Copy(dstfile, srcfile); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

func WriteNpmifyFile(filePath string, deps *state.Dependencies) {

	sort.Slice(deps.Bower, func(i, j int) bool {
		return deps.Bower[i].Name < deps.Bower[j].Name
	})

	// TODO Move this out to a config flag
	DoExcel(*deps)

	file, err := JSONMarshalIndent(&deps, "", "  ")
	msg.CheckErr(err)

	err = ioutil.WriteFile(filePath, file, 0644)
	msg.CheckErr(err)
}

func WritePackageJsonFile(filePath string, pkg state.PackageJson) {
	file, err := JSONMarshalIndent(&pkg, "", "    ")
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

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func JSONMarshalIndent(t interface{}, prefix, indent string) ([]byte, error) {
	b, err := JSONMarshal(t)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, b, prefix, indent)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}