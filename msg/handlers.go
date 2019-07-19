package msg

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
	"log"
)

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