package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hisamafahri/cm/helper"
	"github.com/hisamafahri/cm/model"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:     "log",
	Aliases: []string{"l"},
	Short:   "Log all of the commit --in a better way",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if current directory is a git repository
		// if it is not, return an error
		_, err := helper.CheckDir()

		if err != nil {
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), err.Error())
			return
		}

		// Check the current branch name
		// Return error if there is an error
		allCommit, err := helper.GetLog()

		if err != nil {
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), err.Error())
			return
		}

		// Load some text for our viewport
		if err != nil {
			fmt.Println("could not load file:", err)
			os.Exit(1)
		}

		p := tea.NewProgram(
			model.Model{Content: allCommit},
			tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
			tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
		)

		if err := p.Start(); err != nil {
			fmt.Println("could not run program:", err)
			os.Exit(1)
		}

		// ======================================
	},
}
