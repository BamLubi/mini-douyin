package utils

import (
	"log"

	"github.com/GUAIK-ORG/go-snowflake/snowflake"
)

func IdGen() int64 {
	flake, err := snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		log.Fatalf("snowflake.NewSnowflake failed with %s\n", err)
	}
	id := flake.NextVal()
	return id
}
