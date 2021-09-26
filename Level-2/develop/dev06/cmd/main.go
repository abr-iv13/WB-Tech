package main

func main() {

	args := parseArgs()
	strs := parseArrOfStrings(args)
	outputStr(strs, args)

}
