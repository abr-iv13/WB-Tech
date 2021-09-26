package main

import (
	"flag"
	"log"
	"os"

	"github.com/abr-iv13/WB-Tech/tree/master/Level-2/develop/dev05/pkg/grep"
)

func usage() {
	log.Printf("Usage: ./grep [OPTION]... PATTERN [FILE]... \n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func main() {
	var after = flag.Int("A", 0, "печатать +N строк после совпадения")
	var before = flag.Int("B", 0, "печатать +N строк до совпадения")
	var context = flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	var count = flag.Bool("c", false, "количество строк")                         // +
	var ignoreCase = flag.Bool("i", false, "игнорировать регистр")                // +
	var invert = flag.Bool("v", false, "вместо совпадения, исключать")            // +
	var fixed = flag.Bool("F", false, "точное совпадение со строкой, не паттерн") // +
	var lineNum = flag.Bool("n", false, "печатать номер строки")                  // +

	var showHelp = flag.Bool("h", false, "Show help message") // +

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	args := flag.Args()

	if len(args) < 2 {
		showUsageAndExit(1)
	}

	var pattern = flag.Args()[0]
	var files = flag.Args()[1:]

	grep := grep.New(*after, *before, *context, *count, *ignoreCase, *invert, *fixed, *lineNum, pattern, files)

	err := grep.Execute()
	if err != nil {
		log.Fatalf("grep: %s", err)
	}

	err = grep.Output()
	if err != nil {
		log.Fatalf("grep: %s", err)
	}
}
