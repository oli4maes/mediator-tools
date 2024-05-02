package main

import (
	"github.com/oli4maes/mediator-tools/cmd"
)

func main() {
	cmd.Execute()

	// Uncomment to test file generation
	// GOLANG
	//filegeneration.GenerateFiles("test", filegeneration.HandlerFileType, filegeneration.Golang)
	//filegeneration.GenerateFiles("test", filegeneration.RequestFileType, filegeneration.Golang)
	//filegeneration.GenerateFiles("test", filegeneration.ResponseFileType, filegeneration.Golang)

	// C#
	//filegeneration.GenerateFiles("Test", filegeneration.HandlerFileType, filegeneration.CSharp)
	//filegeneration.GenerateFiles("Test", filegeneration.RequestFileType, filegeneration.CSharp)
	//filegeneration.GenerateFiles("Test", filegeneration.ResponseFileType, filegeneration.CSharp)
}
