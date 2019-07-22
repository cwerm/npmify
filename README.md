# npmify

## TODO
- [ ] Better documentation
- [ ] Remove bloat from binary (via `go build`) - [Issue](https://github.com/cwerm/npmify/issues/3)
- [ ] Automatically update `package.json` - [Issue](https://github.com/cwerm/npmify/issues/4)

## Usage
_More detailed instructions in the works..._

- Clone the repo
- `cd` to the repo directory and run `./npmify`
    - You might have to `chmod +x ./npmify`
- After the first run, open "$HOME/npmify" and edit the config.json
- Run `./npmify` again.<sup>[1]</sup>
- After the numbers are crunched, open your browser to <a href="http://localhost:1234">http://localhost:1234</a>
- Shake your head in disbelief. 

### Footnotes
1. To only run the web server, without parsing your bower files, all the flag `--webOnly true` when you run `npmify`. 