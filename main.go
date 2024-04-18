package main

import (
	"github.com/oli4maes/mediator-tools/cmd"
)

func main() {
	cmd.Execute()

	// Uncomment to test file generation
	//filegeneration.GenerateFiles("getcocktails", filegeneration.HandlerFileType)
	//filegeneration.GenerateFiles("getcocktails", filegeneration.RequestFileType)
	//filegeneration.GenerateFiles("getcocktails", filegeneration.ResponseFileType)
}
