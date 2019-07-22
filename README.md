# npmify

Are you sick and tired of wrangling with an aging code base? Did you too once bask in the sun when Bower was one of the most popular package managers but never quite moved on? Do you wish there was an easier way to crawl through your endless list of old dependencies to see if you can painlessly convert them over to your `package.json`? Well you're in for a treat.

`npmify` is a command-line utility that takes your weathered `bower.json` file and checks each dependency against Npm's package registry to conveniently give you a list of all outdated and unsupported packages you'd like to convert over.

Listen, we've all been in your shoes before. Save yourself some time and automate the process. Use `npmify` today.

## It's probably best to know...

This project is in it's infancy right now, so there's a fair share of dusts. There are a lot of plans for this project, so if it's not a great fit for your needs right now, please file an issue for anything you'd like to see and maybe stick around for the ride 😃

## TODO

- [ ] Automatically update `package.json`
- [ ] Command line arguments
  - [ ] Independent webserver launch
  - [ ] Config file overrides

## Usage

### Summary

If you're looking to get some quick results, here's what you can do!

1. Clone the repo to your local machine
1. From the project's directory, run `make clean-start`
1. Open up `~/npmify/config.json` in your favorite text editor and fill out the following values:
    * `bower_file_path`
    * `package_json_path`
1. Save your config file and run `make run`
1. If it finished successfully, open up <a href="localhost:1234">localhost:1234</a> to witness true disbelief

### Makefile

There's a handy dandy makefile available for you to use that currently supports some rudimentary commands:
***

#### `make`

Runs the tests (eventually!), builds, and runs the project.
***

#### `make clean-start`

Your best bet for the first time running this. It'll clean, test, build, and run the project for you in 16 characters ~~or~~ no less.
***

#### `make build`

Builds the project and produces an executable.
***

#### `make test`

***Coming soon!***
***

#### `make clean`

Bork your environment? Want to obliterate your binary? Use this.
***

#### `make run`

Runs the existing executable based on the settings in your current config file.
***

#### `make deps`

Want to manually download the go dependencies? Feel free to use this!
***

### Configuration

`npmify` is currently only configured by a json file located at `~/npmify/config.json`. If you're starting from scratch, you can run `npmify` and it will generate this file for you with a handful of defaults.

#### Supported options

```json
{
  "output_dir": Output directory for the results,
  "output_file_name": Output filename,
  "bower_file_path": Absolute filepath to your bower file,
  "package_json_path": Absolute filepath to your package.json (npm or yarn supported),
  "packages": Best not to touch this one for now,
  "version": Version that generated this file so we can help you upgrade down the road
}
```

### Viewing the results

At the end of a successful run, `npmify` will start up a webserver at port 1234 (for your convenience: <a href="localhost:1234">localhost:1234</a>) for you to view the results. Alternatively, an excel file detailing the carnage will be produced in the project's directory, or if you really like reading json files, you can also hop over to the output file as configured by your settings (`~/npmify/npmified.json` by default)

## Contributing

This project is just getting off the ground, so if you have feature requests or run into issues, feel free to file them in the project's issues tab.

Looking to help some more? Terrific! If you want to add a feature, fix a bug, or help get things in order, feel free to fork the repo and submit a PR back to the project once your changes are ready!