package helper

import (
	"fmt"

	"github.com/bitfield/script"
	"github.com/logrusorgru/aurora"
)

func Commit(command string) {
	for _, c := range []string{command} {
		fmt.Println() // add break line
		fmt.Println(aurora.Black(" INFO ").BgBrightWhite().Bold(), "Committing changes...")
		p := script.Exec(c)
		if err := p.Error(); err != nil {
			p.SetError(nil)
			output, _ := p.String()
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), output)
		} else {
			output, _ := p.String()
			fmt.Println(aurora.Black(" INFO ").BgBrightWhite().Bold(), output)
			fmt.Println(aurora.Black(" SUCCESS ").BgGreen().Bold(), "Successfully commit changes")
		}
	}
}
