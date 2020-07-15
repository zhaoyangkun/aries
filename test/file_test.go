package test

import (
	"aries/util"
	"log"
	"testing"
)

func TestGetFileNameOnly(t *testing.T) {
	log.Println("fileName: ", util.GetFileNameOnly("/a/b/c.md"))
}
