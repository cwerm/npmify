package main

import (
	"fmt"
	"io/ioutil"
	"npmify/util"
)

func main() {

	data, err := ioutil.ReadFile("/Users/cwermert/Projects/pax/paxserver/frontend/bower.json")
	if err != nil {
		fmt.Print(err)
	}

	util.BuildDeps(data)

}

