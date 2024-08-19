package search

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var patterns = []string{"node_modules", "__pycache__", "venv", ".env", "vendor/bundle", "target", "bin", "obj", "pkg", "build", "cmake-build-*", "_build", "deps",}

// jumpDirectory takes in the name (case insensitive) of the directory you want to find and a starting directory to begin the search from. This will start a depth-first search from the starting directory to locate the directory passed in in the first argument. the return will be the full path to the directory. Pass . as the second argument if you want to start the search form your current directory.
func JumpDirectory(name string, currDir string) string {
	os.Chdir(currDir)
	cleanPaths, path := getDirs(currDir, name)

	if path != ""{
		return path
	}

	for _, absPath := range cleanPaths{
		path = JumpDirectory(name, absPath)
		if path != "" {
			break
		}
	}

	return path
}


func searchVisDirs(dirs []fs.DirEntry, name string)([]string, string) {
	var visDirs []string
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