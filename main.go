package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"npmify/msg"
	"npmify/state"
	"npmify/util"

	//"io/ioutil"
	"npmify/fs"
	"npmify/web"
	"os"
	"os/user"
	"path/filepath"
)

const defaultConfigFile = "config.json"
var webOnly *bool
var usr, _ = user.Current()

func main() {

	cfg := SetupConfig()

	msg.FancyPrint("/**************************************\n * NPMify v%s\n **************************************/\n", cfg.Version)

	outfile := cfg.OutputDir + "/" + cfg.OutputFileName

	if !*webOnly {
		bowerFile, err := ioutil.ReadFile(cfg.BowerFilePath)
		msg.CheckErr(err)

		util.BuildDeps(bowerFile, outfile)
	}

	web.Init(outfile)
}

func SetupConfig() state.Configuration {
	configuration := state.Configuration{}
	webOnly = flag.Bool("webOnly", false, "Just run the web server")
	cfgFile := flag.String("cfg", filepath.Join(usr.HomeDir, "npmify", defaultConfigFile), "Path to your config")
	configuration = settings(*cfgFile)
	flag.Parse()

	return configuration
}

func settings(filename string) state.Configuration {
	configuration := state.Configuration{}

	// Create the default directory if it doesn't exist.
	defaultDir := filepath.Join(usr.HomeDir, "npmify")
	fmt.Printf("%s exists? %t\n", defaultDir, fs.DirectoryExists(defaultDir))

	fmt.Printf("%s exists? %t\n", filename, fs.FileExists(filename))
	// If config file doesn't exist, create it.
	if !fs.FileExists(filename) {
		fmt.Println("Config file does not exist, generating default config.")

		fs.CreateDirectoryIfNotExist(defaultDir)

		file, err := os.Create(filename)
		msg.CheckErr(err)

		var b []byte
		// Populate the configuration struct with defaults
		configuration.OutputDir = filepath.Join(usr.HomeDir, "npmify")
		configuration.OutputFileName = "npmified.json"
		configuration.BowerFilePath = "/path/to/bower.json"
		configuration.PackageJsonPath = "/path/to/package.json"
		configuration.Packages = []state.Package{}
		configuration.Version = "0.0.1"

		// struct --> json
		b, err = json.MarshalIndent(configuration, "", "  ")
		msg.CheckErr(err)

		// Write the new config file
		_, err = file.Write(b)
		msg.CheckErr(err)

		// Close the file after write to the fs
		err = file.Close()
		msg.CheckErr(err)

		fmt.Printf("Please edit the config file at %s and run the program again", filename)
		os.Exit(0)
	}

	file, err := os.Open(filename)
	msg.CheckErr(err)

	// Parse JSON config
	err = json.NewDecoder(file).Decode(&configuration)
	msg.CheckErr(err)

	return configuration
}

