package main

import (
	"github.com/rna-vt/devicecommander/src/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		logrus.Fatal(err)
	}
}
