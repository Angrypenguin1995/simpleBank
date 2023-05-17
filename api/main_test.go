package api

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	config, err := util.loadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB
}
