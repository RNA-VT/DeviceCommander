package main

import (
	"github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/src/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
