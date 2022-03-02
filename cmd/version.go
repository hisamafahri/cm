package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version number of cm",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cm v1.3.0")
	},
}
