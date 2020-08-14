package test

import (
	"aries/utils"
	"fmt"
	"testing"
)

func TestGetRandom(t *testing.T) {
	fmt.Println("random: ", utils.CreateRandomCode(4))
}
