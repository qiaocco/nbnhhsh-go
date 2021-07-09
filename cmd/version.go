package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of nbnhhsh",
	Long:  `All software has versions. This is nbnhhsh's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nbnhhsh version is v1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
