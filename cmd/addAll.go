package cmd

import (
	"fmt"

	"github.com/bitfield/script"
	"github.com/logrusorgru/aurora"
)

func addAllChanges() {
	for _, c := range []string{"git add ."} {
		fmt.Println(aurora.Black(" INFO ").BgBrightWhite().Bold(), "Staging changes...")
		p := script.Exec(c)
		if err := p.Error(); err != nil {
			p.SetError(nil)
			output, _ := p.String()
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), output)
		} else {
			fmt.Println(aurora.Black(" SUCCESS ").BgGreen().Bold(), "Successfully staged changes")
		}
	}
}
