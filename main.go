package main

import (
	_ "github.com/lib/pq"
	"github.com/rodiond26/gomigrator/cmd"
)

func main() {
	cmd.Execute()
}
