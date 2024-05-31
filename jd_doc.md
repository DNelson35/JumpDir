# Jump Directory CLI Tool Documentation

## Overview

**Jump Directory (JD)** is a Python script designed to be executed as an executable file, enabling direct execution from the terminal boot. This tool aims to streamline navigation through directories by allowing users to define a starting and ending directory as arguments. Upon execution, JD prints the absolute path of the target directory, facilitating immediate transition to the desired location within the terminal. To utilize JD, simply type `jd [start_dir] [end_dir]` in the terminal, and it will automatically navigate to the specified directory.

### Purpose and Benefits

The primary motivation behind creating JD was to eliminate the cumbersome process of manually navigating through directories using `cd` commands or relying on the potentially unreliable and slow Bash `find` command for locating specific files. JD offers a faster and more efficient method for quickly accessing working directories without the need for complex regular expressions or extensive searching.

### Functionality

At the core of JD's functionality is the `find_matching_directory` function, which iterates through directories starting from a given path, searching for a directory that matches the provided pattern. This function utilizes `os.walk` to traverse directories and subdirectories efficiently.

python import os

def find_matching_directory(start_dir, pattern): for root, dirs, files in os.walk(start_dir): for dir_name in dirs: if dir_name == pattern: return os.path.join(root, dir_name) return None


The script is structured to prevent automatic execution of its main logic upon import, ensuring that the core functionality remains accessible only when the script is run directly. This design choice allows for flexible integration into larger projects without unintended side effects.

### Usage

To execute JD, first download or clone the repository form github. Then you need to open your terminal and make the file exactable like in this example:
```zsh
chmod +x scripts/jump_dir/Jump_Directory.py
```
replace the script/jump_dir/Jump_directory.py with the path to where you have the file.
once it is executable you will need to set up a function in the boot process of your terminal and then run it form source. there are a few ways to set up the function here are two examples:

1. Define a shell function named `jd` in your `.bashrc`, `.zshrc`, or equivalent configuration file:

```bash
function jd() {
  cd ~
  local target_dir="$1"
  local dir=$(./scripts/jump_dir/Jump_Directory.py . $target_dir)
  cd $dir
}
```

2. Source your configuration file or restart your terminal to apply the changes, example:

```zsh
source ~/.zshrc
```


3. use your newly created shell function from the command line:

```zsh
jd [end_directory]
```

In this implementation you are first setting your directory back to the root directory then the function assigns your target directory to a variable named `target_dir' which is then passed to the python script the script in this example has the current directory hard-coded as the first argument and for the second argument uses the target directory passed in when jd is called form the command line. I recommend this implementation if you have smaller file system seeing how it can be called no matter what directory you are in and doesn't require an initial starting directory to find the target directory.

The Next configuration I recommend you use if you have a large file system with multiple directories with sub directories of the same name. if you use the first implementation you will get the first matching directory. but in this case you may be looking for a directory with the same name some levels deeper if that is the case the this second implementation is better suited.

1. Define a shell function named `jd` in your `.bashrc`, `.zshrc`, or equivalent configuration file:

```bash
function jd() {
    cd ~
    local start_dir="$1"
    local target_dir="$2"
    local dir=$(./scripts/Jump_Directory.py $start_dir $target_dir)
    cd $dir
}
```

2. Source your configuration file or restart your terminal to apply the changes, example:

```zsh
source ~/.zshrc
```

3. use your newly created shell function from the command line:

```zsh
jd [starting_directory] [end_directory]
```

In this implementation you will again be immediately navigated back to root. The main difference here is you can pass a starting directory the starting directory should be the initial path you would like the script to start looking. so if I had two directories structured like this:

```
- root
 - code
    - javascript
        - code
            - project_1
                - javascript_project.js
    - python
        - code
            - Project_1
                - python_project.py
 - documents
 - etc...
```
If i where looking for project_1 for javascript my command would look like this:

```zsh
jd ./code/javascript project_1
# and for python
jd ./code/python project_1
```

You can think of the starting command as getting the search started in the right direction even though there are multiple matching directories I only need to point it to the path I want it to start its search and it will find the first match in that direction.



