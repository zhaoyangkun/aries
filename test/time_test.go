package test

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestGetNowTime(t *testing.T) {
	s := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println("s: ", s)
}
