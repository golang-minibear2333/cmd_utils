package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// ExecuteCommand 转发子命令
func ExecuteCommand(name string, subname string, args ...string) (string, error) {
	args = append([]string{subname}, args...)

	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()

	return string(bytes), err
}

// Error 打印报错到控制台
func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v\n", cmd.Name(), args, err)
	if err != nil {
		return
	}
	os.Exit(1)
}
