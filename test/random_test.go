package test

import (
	"aries/util"
	"log"
	"testing"
)

func TestRandomDigit(t *testing.T) {
	s := util.CreateRandom(6)
	log.Println("s: ", s)
}
