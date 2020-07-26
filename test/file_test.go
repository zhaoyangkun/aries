package test

import (
	"aries/util"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestGetFileNameOnly(t *testing.T) {
	log.Println("fileName: ", util.GetFileNameOnly("/a/b/c.md"))
}
