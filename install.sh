#!/bin/bash

# Define the binary name and URL for downloading
BINARY_NAME="jumpdir"
DOWNLOAD_URL="https://github.com/yourusername/yourrepository/releases/latest/download/$BINARY_NAME"
INSTALL_DIR="/usr/local/bin"

# Download and install the binary
echo "Downloading $BINARY_NAME..."
curl -L "$DOWNLOAD_URL" -o "$INSTALL_DIR/$BINARY_NAME"

# Make the binary executable
chmod +x "$INSTALL_DIR/$BINARY_NAME"

# Add function to the user's shell configuration
SHELL_CONFIG="$HOME/.zshrc"  # Change this to .bashrc or other if needed

echo "Configuring shell..."
cat <<EOL >> "$SHELL_CONFIG"
# Function to use the tool
function jd() {
  local target_dir="\$1"
  local start_dir="\$2"

  if [[ "\$target_dir" == "-help" || "\$target_dir" == "--help" ]]; then
    $INSTALL_DIR/$BINARY_NAME -help
    return 0
  fi
  if [[ "\$start_dir" == "." ]]; then
    local base_path="\$(pwd)"
    cd ~
    local sdir="\${base_path}"
    local dir=\$($INSTALL_DIR/$BINARY_NAME \$target_dir \$sdir)
  else
    cd ~
    local base_path="\$(pwd)"
    local sdir="\${base_path}\${start_dir:+/\$start_dir}"
    local dir=\$($INSTALL_DIR/$BINARY_NAME \$target_dir \$sdir)
  fi
  cd \$dir
}
EOL

echo "Setup complete. Please restart your terminal or run 'source $SHELL_CONFIG'."
