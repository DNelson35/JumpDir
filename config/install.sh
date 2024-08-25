#!/bin/bash

# Define the URL to the GitHub release
RELEASE_URL="https://github.com/DNelson35/JumpDir/releases/download/v0.1.0-alpha/jumpdir"

CONFIG_URL="https://github.com/DNelson35/JumpDir/releases/download/v0.1.0-alpha/config.json"

# Define the new installation directory
INSTALL_DIR="$HOME/jumpdir-bin"

# Create the installation directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

# Download the binary
echo "Downloading jumpdir..."
curl -L -o "$INSTALL_DIR/jumpdir" "$RELEASE_URL"

echo "Downloading config"
curl -L -o "$INSTALL_DIR/config.json" "$CONFIG_URL"

# Make the binary executable
chmod +x "$INSTALL_DIR/jumpdir"

# Add function to the user's shell configuration
SHELL_CONFIG="$HOME/.zshrc"  # Change this to .bashrc or other if needed

echo "Configuring shell..."

# Ensure that the jd function is not already present in .zshrc
if ! grep -q 'function jd() {' "$SHELL_CONFIG"; then
  # Add the function to .zshrc
  cat <<'EOL' >> "$SHELL_CONFIG"
# Function to use the tool
function jd() {
  export CONFIG_PATH="$HOME/jumpdir-bin/config.json"
  local target_dir="$1"
  local start_dir="$2"


  if [[ "$target_dir" == "-help" || "$target_dir" == "--help" ]]; then
    ./jumpdir-bin/jumpdir -help
    return 0 
  fi
  if [[ "$start_dir" == "." ]]; then
    local base_path="$(pwd)"
    cd ~
    local sdir="${base_path}"
    local dir=$(./jumpdir-bin/jumpdir $target_dir $sdir)
  else
    cd ~
    local base_path="$(pwd)"
    local sdir="${base_path}${start_dir:+/$start_dir}"
    local dir=$(./jumpdir-bin/jumpdir $target_dir $sdir)
  fi
  cd $dir
}
EOL

  echo "Added jd function to $SHELL_CONFIG."
else
  echo "jd function already present in $SHELL_CONFIG."
fi

echo "Setup complete. Please restart your terminal or run 'source $SHELL_CONFIG' to update your PATH and function."
