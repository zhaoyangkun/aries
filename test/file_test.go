package test

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestGetFileNameOnly(t *testing.T) {
	log.Println("fileName: ", utils.GetFileNameOnly("/a/b/c.md"))
}
