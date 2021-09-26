package main

import (
	"github.com/abr-iv13/WB-Tech/tree/master/Level-2/develop/dev06/pkg/cut"
)

func main() {

	args := cut.ParseArgs()
	strs := cut.ParseArrOfStrings(args)
	cut.OutputStr(strs, args)

}
