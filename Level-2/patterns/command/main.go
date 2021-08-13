package main

import (
	"github.com/abr-iv13/WB-Tech/tree/master/Level-2/patterns/command/tv"
)

func main() {
	tv := &tv.Tv{}

	onCommand := &tv.OnCommand{
		Device: tv,
	}

	offCommand := &tv.OffCommand{
		Device: tv,
	}

	onButton := &tv.Button{
		Command: onCommand,
	}
	onButton.press()

	offButton := &tv.Button{
		Command: offCommand,
	}
	offButton.press()
}
