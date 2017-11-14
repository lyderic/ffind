package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	hiddenPtr := flag.Bool("H", false, "show hidden files/dirs")
	removePtr := flag.Bool("Hr", false, "remove hidden files/dirs")
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
    if info == nil { return nil }
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

func exists(path string) (found bool) {
	found = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		found = false
	}
	return
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
