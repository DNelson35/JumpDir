package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/DNelson35/JumpDir/search"
)



func printUsage() {
	    fmt.Println("Usage: go run main.go <target_directory> [<starting_point>]")
    fmt.Println()
    fmt.Println("Arguments:")
    fmt.Println("  <target_directory>    Target directory (required)")
    fmt.Println("  [<starting_point>]    Starting point (optional)")
    fmt.Println()
    fmt.Println("Flags:")
    flag.PrintDefaults()
}



func main(){
	var help bool
	flag.BoolVar(&help, "help", false, "Show help information")
	flag.BoolVar(&help, "h", false, "Show help information") 
	flag.Parse()
	
	if help {
		printUsage()
		os.Exit(0)
	}
	
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Error: <target_directory> and <starting_point> are required.")
		printUsage()
		os.Exit(1)
	}

	name := args[0]
	currDir := args[1]

	result := search.JumpDirectory(name, currDir)
	fmt.Println(result)

}

