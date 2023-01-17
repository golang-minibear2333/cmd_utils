/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// viperallCmd represents the viperall command
var viperallCmd = &cobra.Command{
	Use:   "viperall",
	Short: "Show cfg all",
	Long: `Show the contents of the entire configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.AllSettings())
	},
}

func init() {
	rootCmd.AddCommand(viperallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viperallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viperallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
