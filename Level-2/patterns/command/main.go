package main

import (
	"github.com/abr-iv13/WB-Tech/tree/master/Level-2/patterns/command/tv"
)

func main() {
	television := &tv.Tv{}

	onCommand := &tv.OnCommand{
		Device: television,
	}

	offCommand := &tv.OffCommand{
		Device: television,
	}

	onButton := &tv.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton := &tv.Button{
		Command: offCommand,
	}
	offButton.Press()
}
