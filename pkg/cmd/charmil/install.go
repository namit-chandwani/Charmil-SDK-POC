package charmil

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// InstallCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install plugin(s)",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		// fmt.Println("install called")
		// return LoadCommands(cmd)
		// args = append([]string{"docs"}, args...)
		// fmt.Println(args)
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(cwd)
		c := exec.Command("ls", args...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		var buf bytes.Buffer
		c.Stderr = io.MultiWriter(os.Stderr, &buf)

		return c.Run()
	},
}

func init() {
}
