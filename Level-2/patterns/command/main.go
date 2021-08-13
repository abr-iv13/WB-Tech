package main

import (
	_ "github.com/abr-iv13/WB-Tech/tree/master/Level-2/patterns/command/tv"
)

func main() {
	tv := &tv.tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.Press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.Press()
}
