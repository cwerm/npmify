package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/logrusorgru/aurora"
	"io/ioutil"

	//"io/ioutil"
	"npmify/fs"
	"npmify/util"
	"npmify/web"
	"os"
	"os/user"
	"path/filepath"
)

type Configuration struct {
	OutputDir		string `json:"output_dir"`
	OutputFileName 	string `json:"output_file_name"`
	BowerFilePath	string `json:"bower_file_path"`
}

const defaultConfigFile = "config.json"

var usr, _ = user.Current()

func main() {

	cfg := SetupConfig()

	bowerFile, err := ioutil.ReadFile(cfg.BowerFilePath)
	util.CheckErr(err)

	outfile := cfg.OutputDir + "/" + cfg.OutputFileName

	util.BuildDeps(bowerFile, outfile)

	web.Init(outfile)
}

func SetupConfig() Configuration {
	fmt.Printf(aurora.Sprintf(aurora.BgBlue("  NPMify v0.0.1  %s").BrightWhite().Bold(), "\n"))

	configuration := Configuration{}

	cfgFile := flag.String("cfg", filepath.Join(usr.HomeDir, "npmify", defaultConfigFile), "Path to your config")
	configuration = settings(*cfgFile)
	flag.Parse()

	return configuration
}

func settings(filename string) Configuration {
	configuration := Configuration{}

	file, err := os.Open(filename)

	// If config file doesn't exist, create it.
	if os.IsNotExist(err) {
		fmt.Println("Config file does not exist, generating default config.")
		defaultDir := filepath.Join(usr.HomeDir, "npmify")

		fs.CreateDirectory(defaultDir)

		file, err = os.Create(filename)
		util.CheckErr(err)

		var b []byte
		// Populate the configuration struct
		configuration := Configuration{filepath.Join(usr.HomeDir, "npmify"), "npmified.json", "~/pax/paxserver/frontend/bower.json"}

		// struct --> json
		b, err = json.MarshalIndent(configuration, "", "  ")
		util.CheckErr(err)

		// Write the new config file
		_, err = file.Write(b)
		util.CheckErr(err)

		// Close the file after write to the fs
		err = file.Close()
		util.CheckErr(err)

		fmt.Printf("Please edit the config file at %s and run the program again", filename)
		os.Exit(0)
	}

	util.CheckErr(err)

	// Parse JSON config
	err = json.NewDecoder(file).Decode(&configuration)
	util.CheckErr(err)

	return configuration
}

