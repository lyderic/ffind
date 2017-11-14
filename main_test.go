package main

import (
	"log"
	"os"
	"testing"
)

const dir = "dir"

func init() {
	log.SetFlags(log.Lshortfile)
	createDirStructure()

}

func TestListHidden(t *testing.T) {
	if !exists(dir) {
		t.Error("dir not found!")
		return
	}
	listHidden(dir)
}

func TestRemoveHidden(t *testing.T) {
	if !exists(dir) {
		t.Error("dir not found!")
		return
	}
	removeHidden(dir)
	if exists("dir/.hidden") {
		t.Error(".hidden has not been removed!")
	}
}

func TestRemoveDirStructure(t *testing.T) {
	err := os.RemoveAll(dir)
	if err != nil {
		t.Error("Failed to removed testing directory 'dir'")
	}
}

func createDirStructure() {
	os.MkdirAll("dir/.hidden", 0755)
}
