package plugin

import (
	"fmt"

	"github.com/spf13/cobra"
)

// InstallCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install plugins",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
	},
}

func init() {
}
