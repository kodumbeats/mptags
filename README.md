# mptags

For some reason, all the fancy tools to edit song tags were too cumbersome and heavy for my liking. `mptags` is a quick-and-dirty tool for bulk assigning song metadata to audio files based on their filename/dir.

## Usage

```
mptags is a tool to bulk assign tags to music files based on their filename/dir

Usage:
  mptags [argument] [--flags]

Arguments:
  path: (optional) album path, defaults to $PWD
    $ mptags /path/to/album

  help: show command information
    $ mptags help

  version: show build version
    $ mptags version

Flags:
  -h, --help: show command information
  -v, --version: show build version
```

## Development

Requires static [taglib](https://taglib.org) libraries to compile - see [go-taglib](https://github.com/wtolson/go-taglib) for more.

**arch (btw)**
    `sudo pacman -S taglib`

### Build

`./bin/build`

# License

MIT
