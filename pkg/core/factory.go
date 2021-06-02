package core

import (
	"log"

	ext "github.com/namit-chandwani/Charmil-SDK-POC/pkg/cmd/charmil"
	"github.com/spf13/cobra"
)

const (
	hostKey   = "host"
	pluginKey = "plugin"
)

type CLI interface {
	AddCommands(cmd *cobra.Command) error
}

type host struct{}
type plugin struct{}

func (h host) AddCommands(cmd *cobra.Command) error {
	return NewCommand(cmd)
}

func (e plugin) AddCommands(cmd *cobra.Command) error {
	cmd.AddCommand(ext.InstallCmd, ext.InstalledCmd, ext.ListCmd, ext.RemoveCmd, ext.ActivateCmd, ext.DeactivateCmd)
	return nil
}

func GetCLI(cliType string) CLI {
	switch cliType {
	case hostKey:
		return &host{}
	case pluginKey:
		return &plugin{}
	default:
		log.Printf("Undefined CLI type")
		return nil
	}
}
