package test

import (
	"log"
	"testing"
)

func TestRandomDigit(t *testing.T) {
	s := utils.CreateRandom(6)
	log.Println("s: ", s)
}
