package main

import (
	"github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/src/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		logrus.Fatal(err)
	}
}
