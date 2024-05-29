<p align="center">
<img width="500" height="200" src="./GIME.png" alt="gime">
<br /> <br />
<img alt="GitHub Tag" src="https://img.shields.io/github/v/tag/d1agnoze/gime?include_prereleases">
<img alt="GitHub Release" src="https://img.shields.io/github/v/release/d1agnoze/gime?include_prereleases">
<br/>
<img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/d1agnoze/gime">
<img alt="License" src="https://img.shields.io/badge/License-MIT-green.svg"/>
</p>

----
`Gime` is a lightweight program built with [Wails](https://wails.io/) that help you preview images and watch youtube video right from terminal

Supported Platform:
| Platform | Status    |
|----------|-----------|
| Windows  | Supported |
| Linux    | Untested  |

## Usage

```bash
gime [flag]
```

Pipe from stdin:
```bash
echo image/link.png | gime --img
```
Use `--file` flag:
```bash
gime --img --file=image/link.png
```


## Development

Prerequisite: 
- Go 1.20+
- NPM (Node 15+)
- Wails CLI 
- GNU Make (OPTIONAL)
- upx (OPTIONAL - for application compression)

To run in live development mode, run `make dev` in the project directory. Changes the argument to your liking in the makefile

## Building

To build a redistributable, production mode package, use `make build_linux` or `make build_win`.

## Authors

- [@vdac](https://www.github.com/d1agnoze)
