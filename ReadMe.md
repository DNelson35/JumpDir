# JumpDir CLI Tool

`JumpDir` is a command-line tool written in Go that helps you find directories by name. It performs a depth-first search starting from a specified directory and returns the full path to the target directory.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Zsh Integration](#zsh-integration)
- [Help and Flags](#help-and-flags)
- [Examples](#examples)
- [Troubleshooting](#troubleshooting)

## Installation

To install `JumpDir`, follow these steps:

1. **Download the Latest Release**

   Download the `install.sh` script from the [releases page](https://github.com/DNelson35/JumpDir/releases) of this repository.

2. **Make the Installation Script Executable**

   ```sh
   chmod +x install.sh
   ```

3. **Run the Installation Script**

   ```sh
   ./install.sh
   ```

   The script will download the latest `jumpdir` binary, make it executable, and configure your shell to use the `jd` function.

4. **Restart Your Terminal**

   Restart your terminal or run the following command to apply the changes:

   ```sh
   source ~/.zshrc
   ```

## Usage

The `JumpDir` tool requires two arguments:

- `<target_directory>`: The name of the directory you want to find (required).
- `<starting_point>`: The directory where the search should start (optional). If omitted, the search will start from the home directory. If you pass `.` as the `<starting_point>`, the search will start from your current directory.

### Command Syntax

```sh
jd <target_directory> [<starting_point>]
```

### Example

```sh
jd "targetDirName" /path/to/start
```

This command will search for "targetDirName" starting from `/path/to/start`.

If you want to search starting from your current directory, use:

```sh
jd "targetDirName" .
```

This command will search for "targetDirName" starting from your current location in the file system.

## Zsh Integration

The `install.sh` script configures the `jd` function for you. After running the script, the function is automatically added to your `.zshrc` file.

### Function Overview

The `jd` function:

- Searches for the specified directory.
- Changes to the found directory if it exists.

## Help and Flags

To display help information for the `JumpDir` tool, use the `-help` or `-h` flags:

```sh
jd -help
```

### Output

```
Usage: go run main.go <target_directory> [<starting_point>]

Arguments:
  <target_directory>    Target directory (required)
  [<starting_point>]    Starting point (optional)

Flags:
  -h, -help   Show help information
```

## Examples

1. **Find a Directory Starting from the Home Directory**

   ```sh
   jd Documents
   ```

   This searches for a directory named `Documents` starting from your home directory.

2. **Find a Directory Starting from a Specific Directory**

   ```sh
   jd myProject workspace
   ```

   This searches for "myProject" starting from `~/workspace`.

3. **Find a Directory Starting from the Current Directory**

   ```sh
   jd myProject .
   ```

   This searches for "myProject" starting from your current location in the file system.

## Troubleshooting

- **Error: `target_directory` and `starting_point` are required**: Ensure that both arguments are provided if using the `jd` function.
- **Permission Denied**: Verify that `JumpDir` is executable and properly installed.
- **Function Not Found**: Ensure that `install.sh` has been run and the `.zshrc` file has been reloaded.
