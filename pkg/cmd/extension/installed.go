package extension

import (
	"fmt"

	"github.com/spf13/cobra"
)

// InstalledCmd represents the installed command
var InstalledCmd = &cobra.Command{
	Use:   "installed",
	Short: "Show installed extensions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("installed called")
	},
}

func init() {
}
