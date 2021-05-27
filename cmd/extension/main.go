package main

import (
	"log"

	"github.com/namit-chandwani/HostCLI-POC/pkg/cmd/plugincore"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const extensionKey = "extension"

func main() {
	cmd := &cobra.Command{
		Use:          "extension",
		Short:        "Extension commands",
		SilenceUsage: true,
	}

	extensionCLI := plugincore.GetCLI(extensionKey)
	err := extensionCLI.AddCommands(cmd)
	if err != nil {
		log.Fatal(err)
	}

	// err := plugincore.AddCommands(cmd)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = doc.GenMarkdownTree(cmd, "./docs/commands/extension")
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
