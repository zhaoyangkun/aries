package test

import (
	"aries/utils"
	"log"
	"testing"
)

func TestGetFileNameOnly(t *testing.T) {
	log.Println("fileName: ", utils.GetFileNameOnly("/a/b/c.md"))
}
