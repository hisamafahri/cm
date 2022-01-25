package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Commit() *cobra.Command {
	return &cobra.Command{
		Use:   "commit",
		Short: "Commit is CLI to help you structurized your commit message",
		Long: `Simple but powerful CLI to help your commit cmessage to follow
		conventional commit message`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("COMMITTED!")
		},
	}
}

func Execute() {
	if err := Commit().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
