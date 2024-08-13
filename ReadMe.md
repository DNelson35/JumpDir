# JumpDir CLI Tool

`JumpDir` is a command-line tool written in Go that helps you find directories by name. It performs a depth-first search starting from a specified directory and returns the full path to the target directory.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Zsh Integration](#zsh-integration)
- [Help and Flags](#help-and-flags)
- [Examples](#examples)

## Installation

1. **Clone the Repository**

    Clone the repository containing the Go code <!-- and the Zsh function script: -->

   ```sh
   git clone https://github.com/DNelson35/JumpDir.git
   ```

2. **Build the Go Tool**

   Navigate to the directory containing the Go code and build the tool:

   ```sh
   cd <path-to-go-code>
   go build -o JumpDir main.go
   ```

   This will create an executable named `JumpDir` in the `./scripts/jump_dir/` directory.

3. **Ensure the Go Binary is Executable**

   Make sure the `JumpDir` binary is executable:

   ```sh
   chmod +x ./scripts/jump_dir/JumpDir
   ```

## Usage

The `JumpDir` tool requires two arguments:

- `<target_directory>`: The name of the directory you want to find (required).
- `<starting_point>`: The directory where the search should start (optional). If omitted, the search will start from the home directory.

### Command Syntax

```sh
./scripts/jump_dir/JumpDir <target_directory> [<starting_point>]
```

### Example

```sh
./scripts/jump_dir/JumpDir "targetDirName" /path/to/start
```

This command will search for "targetDirName" starting from `/path/to/start`.

## Zsh Integration

To use the `JumpDir` tool with a custom `jd` function in your Zsh shell, follow these steps:

1. **Add the Function to Your `.zshrc`**

   Open your `.zshrc` file and add the following function:

   ```sh
   function jd() {
     local target_dir="$1"
     local start_dir="$2"

     if [[ "$target_dir" == "-help" || "$target_dir" == "--help" ]]; then
       ./scripts/jump_dir/JumpDir -help
       return 0
     fi

     if [[ "$start_dir" == "." ]]; then
       local base_path="$(pwd)"
       cd ~
       local sdir="${base_path}"
       local dir=$(./scripts/jump_dir/JumpDir "$target_dir" "$sdir")
     else
       cd ~
       local base_path="$(pwd)"
       local sdir="${base_path}${start_dir:+/$start_dir}"
       local dir=$(./scripts/jump_dir/JumpDir "$target_dir" "$sdir")
     fi
     cd "$dir"
   }
   ```

2. **Reload Your Zsh Configuration**

   Apply the changes by reloading your `.zshrc`:

   ```sh
   source ~/.zshrc
   ```

## Help and Flags

To display help information for the `JumpDir` tool, use the `-help` or `-h` flags:

```sh
./scripts/jump_dir/JumpDir -help
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

   This searches for a directory named Documents starting from your home directory.

2. **Find a Directory Starting from a Specific Directory**

   ```sh
   jd myProject workspace
   ```

   This searches for "myProject" starting from `~/workspace`.

3. **Find a Directory Starting from a Current Directory**

   ```sh
   jd myProject .
   ```

   This searches for "myProject" starting from your current location in the file system.

## Troubleshooting

- **Error: `target_directory` and `starting_point` are required**: Ensure that both arguments are provided if using the `jd` function.
- **Permission Denied**: Verify that `JumpDir` is executable and properly built.
