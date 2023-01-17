package cmd

import (
	"fmt"
	"github.com/golang-minibear2333/cmd_utils/git/utils"
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version subcommand show git version info.",

	Run: func(cmd *cobra.Command, args []string) {
		// 把子命令转发到`git`上去了
		output, err := utils.ExecuteCommand("git", "version", args...)
		if err != nil {
			utils.Error(cmd, args, err)
		}

		fmt.Fprint(os.Stdout, output)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
