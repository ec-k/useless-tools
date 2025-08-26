package main

import (
	"flag"
	"fmt"
	"os"
	"useless-tools/useless"
)

var (
	processorName = flag.String("processor", "", "Specify the processor to use (e.g., 'randomize')")
	inputString   = flag.String("input", "", "Input string for the selected processor")
)

func main() {
	flag.Parse()

	if *processorName == "" {
		fmt.Fprintln(os.Stderr, "Error: A processor must be specified using --processor flag.")
		fmt.Fprintln(os.Stderr, "Usage: ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var processor useless.Processor

	switch *processorName {
	case "randomize":
		processor = useless.Randomizer{}
	default:
		fmt.Fprintf(os.Stderr, "Error: Unknown processor '%s'.\n", *processorName)
		os.Exit(1)
	}

	output, err := processor.Process(*inputString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing input: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(output)
	os.Exit(0)
}
