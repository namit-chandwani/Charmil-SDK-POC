package main

import (
	"log"

	"github.com/craicoverflow/socket/pkg/pluginloader"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd := &cobra.Command{
		Use:          "git-helper",
		Short:        "Git helper commands",
		SilenceUsage: true,
	}

	err := pluginloader.AddCommands(cmd)
	if err != nil {
		log.Fatal(err)
	}

	err = doc.GenMarkdownTree(cmd, "./docs/commands")
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
