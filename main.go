package main

import (
	"fmt"
	"flag"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"slices"
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

// jumpDirectory takes in the name (case insensitive) of the directory you want to find and a starting directory to begin the search from. This will start a depth-first search from the starting directory to locate the directory passed in in the first argument. the return will be the full path to the directory. Pass . as the second argument if you want to start the search form your current directory.
func jumpDirectory(name string, currDir string) string {
	os.Chdir(currDir)
	cleanPaths, path := getDirs(currDir, name)

	if path != ""{
		return path
	}

	for _, absPath := range cleanPaths{
		path = jumpDirectory(name, absPath)
		if path != "" {
			break
		}
	}

	return path
}


func searchVisDirs(dirs []fs.DirEntry, name string)([]string, string) {
	var visDirs []string
	patterns := []string{"node_modules", "__pycache__", "venv", ".env", "vendor/bundle", "target", "bin", "obj", "pkg", "build", "cmake-build-*", "_build", "deps",}
	for _, dir := range dirs{
		if dir.Name()[0] == '.' || !dir.IsDir() || slices.Contains(patterns, dir.Name()) {
			continue
		}else if checkMatch(name, dir){
			path, _ := filepath.Abs(dir.Name())
			return nil, path
		}

		if len(dirs) != 0{
			path, err := filepath.Abs(dir.Name())
			if err != nil {
				fmt.Println(err)
			}
			visDirs = append(visDirs, path)
		}
		
	}
	return visDirs, ""
}

func checkMatch(name string, dir fs.DirEntry)bool{	
	return strings.EqualFold(dir.Name(), name ) 
}

func getDirs(dir string, name string)([]string, string){
	dirs, err := os.ReadDir(dir)
	
	if err != nil {
		fmt.Println(err)
	}

	return searchVisDirs(dirs, name)
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

	result := jumpDirectory(name, currDir)
	fmt.Println(result)

}

