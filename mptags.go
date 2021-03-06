package main

import (
	"bufio"
	"fmt"
	"github.com/wtolson/go-taglib"
	"io/fs"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

var GitCommit string // to track builds, provided at buildtime

func main() {
	args := os.Args[1:] // trim command from list

	for _, arg := range args {
		if arg == "--help" || arg == "-h" || arg == "help" {
			printHelp()
		}
		if arg == "--version" || arg == "-v" || arg == "version" {
			printVersion()
		}
	}

	// use provided path, else current working dir
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if len(args) > 0 {
		pwd = args[0]
	}
	if len(args) > 1 {
		log.Fatal("Too many arguments")
		os.Exit(3)
	}

	// name album after parent dir
	_, album := path.Split(pwd)

	fileSystem := os.DirFS(pwd)
	songs, globErr := fs.Glob(fileSystem, "*")
	if globErr != nil {
		log.Fatal(globErr)
	}

	if len(songs) == 0 {
		log.Fatal("no songs found in: " + album)
	}

	confirmed := false

	for _, s := range songs {
		// format of "# artist - title"
		before, after, _ := strings.Cut(s, " ")
		track, atoiErr := strconv.Atoi(before) // track number
		if atoiErr != nil {
			log.Fatal(atoiErr)
		}

		artist, audioFile, _ := strings.Cut(after, "-")
		artist = strings.TrimSpace(artist) // trim whitespaces

		title, _, _ := strings.Cut(audioFile, ".") // trim extension
		title = strings.TrimSpace(title)

		if confirmed == false {
			fmt.Println("Dry run:")
			fmt.Println("")
			fmt.Println("Album: " + album + " (" + strconv.Itoa(len(songs)) + " songs)")
			fmt.Println("Track: " + strconv.Itoa(track))
			fmt.Println("Title: " + title)
			fmt.Println("Artist: " + artist)
			fmt.Println("")
			fmt.Print("Proceed with changes? [y/N] ")

			reader := bufio.NewReader(os.Stdin)
			confirm, inputErr := reader.ReadString('\n')
			if inputErr != nil {
				log.Fatal(inputErr)
			}
			if strings.ToLower(strings.TrimSpace(confirm)) != "y" {
				os.Exit(0)
			}

			confirmed = true
		}

		// Proceed when confirmed
		song, readErr := taglib.Read(path.Join(pwd, s))
		if readErr != nil {
			log.Fatal(readErr)
		}

		song.SetTrack(track)
		song.SetArtist(artist)
		song.SetTitle(title)
		song.SetAlbum(album)

		saveErr := song.Save()
		if saveErr != nil {
			log.Fatal(saveErr)
		}

		song.Close()
	}
}

func printHelp() {
	fmt.Println(`
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
  `)

	os.Exit(0)
}

func printVersion() {
	fmt.Printf("mptags build %s\n", GitCommit)
	os.Exit(0)
}
