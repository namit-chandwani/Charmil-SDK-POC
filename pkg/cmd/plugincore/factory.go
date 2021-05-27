package plugincore

import (
	"log"

	ext "github.com/namit-chandwani/HostCLI-POC/pkg/cmd/extension"
	"github.com/spf13/cobra"
)

const (
	hostKey      = "host"
	extensionKey = "extension"
)

type CLI interface {
	AddCommands(cmd *cobra.Command) error
}

type host struct{}
type extension struct{}

func (h host) AddCommands(cmd *cobra.Command) error {
	return NewCommand(cmd)
}

func (e extension) AddCommands(cmd *cobra.Command) error {
	cmd.AddCommand(ext.InstallCmd, ext.InstalledCmd, ext.ListCmd, ext.RemoveCmd)
	return nil
}

func GetCLI(cliType string) CLI {
	switch cliType {
	case hostKey:
		return &host{}
	case extensionKey:
		return &extension{}
	default:
		log.Printf("Undefined CLI type")
		return nil
	}
}
