package main

import (
	"fmt"

	"github.com/Pitasi/openswitch/internal/platform/director"
)

func main() {
	greet()
	director.Run(
		&Updater{},
	)
	director.Wait()
	goodbye()
}

func greet() {
	fmt.Printf(`#######################
#  OPENSWITCH SERVER  #
#  %10s         #
#######################
`, Version)
}

func goodbye() {
	fmt.Println("OpenSwitch shutted down")
}

var Version = "dirty"
