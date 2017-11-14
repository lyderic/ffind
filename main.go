package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const version = "0.0.1"

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	hiddenPtr := flag.Bool("H", false, "show hidden files/dirs")
	removePtr := flag.Bool("R", false, "remove hidden files/dirs")
	flag.Usage = usage
	flag.Parse()
	hidden := *hiddenPtr
	remove := *removePtr

	basedir := "."
	if flag.NArg() > 0 {
		basedir = os.Args[len(os.Args)-1]
	}
	if !exists(basedir) {
		log.Fatalln(basedir, "not found!")
	}

	if hidden {
		listHidden(basedir)
		return
	}

	if remove {
		listHidden(basedir)
		removeHidden(basedir)
		return
	}

	/* This is the default, if no flag is used */
	listAll(basedir)

}

func listAll(basedir string) {
	filepath.Walk(basedir, func(path string, info os.FileInfo, err error) error {
		display(path, info)
		return err
	})
}

func listHidden(basedir string) {
	filepath.Walk(basedir, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		if info.Name()[0] == '.' && len(info.Name()) > 1 {
			display(path, info)
		}
		return err
	})
}

func removeHidden(basedir string) {
	filepath.Walk(basedir, func(path string, info os.FileInfo, err error) error {
		if info.Name()[0] == '.' && len(info.Name()) > 1 {
			err = os.RemoveAll(path)
		}
		return err
	})
}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func display(path string, info os.FileInfo) {
	if info == nil {
		return
	}
	if info.IsDir() {
		fmt.Println(path + "/")
	} else {
		fmt.Println(path)
	}
}

func usage() {
	fmt.Printf("%s v.%s - (c) Lyderic Landry, London 2017\n",
		filepath.Base(os.Args[0]), version)
	fmt.Printf("Usage: %s [-H|-R] [<dir>]\n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}
