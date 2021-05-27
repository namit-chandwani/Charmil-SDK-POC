package extension

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove extensions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {

}
