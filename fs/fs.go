package fs

import (
	"fmt"
	"github.com/fatih/structs"
	xlst "github.com/ivahaev/go-xlsx-templater"
	"log"
	"npmify/state"
	"os"
)

func CreateDirectory(dirName string) bool {
	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {
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
func DoExcel(deps state.Dependencies) {
	var depMap = dependencyMap(deps)
	fmt.Println(depMap["Bower"])
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