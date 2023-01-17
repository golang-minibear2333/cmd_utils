package cmd

import (
	"fmt"
	"github.com/golang-minibear2333/cmd_utils/git/utils"
	"github.com/spf13/cobra"
	"os"
)

var cloneCmd = &cobra.Command{
	Use:   "clone url [destination]",
	Short: "Clone a repository into a new directory",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := utils.ExecuteCommand("git", "clone", args...)
		if err != nil {
			utils.Error(cmd, args, err)
		}

		fmt.Fprintf(os.Stdout, output)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
