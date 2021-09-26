package main

import (
	"fmt"

	"github.com/abr-iv13/WB-Tech/tree/master/Level-2/dev/unpacker/pgk/unpack"
)

func main() {
	r1, _ := unpack.Unpack("a4bc2d5e")
	r2, _ := unpack.Unpack("abcd")
	r3, _ := unpack.Unpack("45")
	r4, _ := unpack.Unpack("")

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)

}
