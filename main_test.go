package main

import (
	"log"
	"os"
	"testing"
)

const dir = "dir"

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
	cleanup()
}

func TestExists(t *testing.T) {
	if !exists(dir) {
		t.Error("'dir' not found!")
	}
}

func TestListAll(t *testing.T) {
	listAll(dir)
}

func TestListHidden(t *testing.T) {
	listHidden(dir)
}

func TestRemoveHidden(t *testing.T) {
	removeHidden(dir)
	if exists("dir/.hidden") {
		t.Error(".hidden has not been removed!")
	}
}

func setup() {
	os.MkdirAll("dir/.hidden/subdir", 0755)
	os.Create("dir/.hidden/toto")
	os.Create("dir/.hidden/.totoHidden")
	os.Create("dir/.hiddenFile")
}

func cleanup() {
	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}
}
