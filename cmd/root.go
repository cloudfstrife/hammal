package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hammal",
	Short: "hammal is a tool do transfer file by keyboard",
	Long:  `hammal is a tool do transfer file by keyboard`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

//RegistCmd add sub command
func RegistCmd(c *cobra.Command) {
	rootCmd.AddCommand(c)
}

//Execute root runner
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
