package tests

import (
	"log"
	"testing"
)
import "go_server/lib/utils"

func TestStringToSha256(t *testing.T) {

	var testStr = "qrqretdfggdhdfghffdghfh"

	if utils.StringToSha256(testStr) != utils.StringToSha256(testStr) {
		log.Println("SHA-256 strings not equals")
		t.Fail()
	}

}
