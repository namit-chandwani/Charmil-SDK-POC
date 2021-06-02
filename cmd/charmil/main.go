package main

import (
	"log"

	"github.com/namit-chandwani/Charmil-SDK-POC/pkg/core"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const cliKey = "plugin"

func main() {
	cmd := &cobra.Command{
		Use:          "charmil",
		Short:        "Commands for managing plugins",
		SilenceUsage: true,
	}

	charmilCLI := core.GetCLI(cliKey)
	err := charmilCLI.AddCommands(cmd)
	if err != nil {
		log.Fatal(err)
	}

	// err := core.AddCommands(cmd)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = doc.GenMarkdownTree(cmd, "./docs/commands")
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
