#!/usr/bin/env bash
set -eo pipefail

echo "Installing ge-publish..."

BASE_DIR="${XDG_CONFIG_HOME:-$HOME}"
GE_PUBLISH_DIR="${GE_PUBLISH_DIR-"$BASE_DIR/.ge-publish"}"
GE_PUBLISH_BIN_DIR="$GE_PUBLISH_DIR/bin"

BIN_PATH="$GE_PUBLISH_BIN_DIR/ge-publish"

uname_s=$(uname -s)
PLATFORM=$(echo "$uname_s" | awk '{print tolower($0)}')
case $PLATFORM in
  linux) ;;
  darwin|mac*)
    PLATFORM="darwin"
    ;;
  mingw*|win*)
    PLATFORM="windows"
    ;;
  *)
    err "unsupported platform: $PLATFORM"
    ;;
esac

uname_m=$(uname -m)
ARCHITECTURE=$(echo "$uname_m" | awk '{print tolower($0)}')
if [ "${ARCHITECTURE}" = "x86_64" ]; then
  # Redirect stderr to /dev/null to avoid printing errors if non Rosetta.
  if [ "$(sysctl -n sysctl.proc_translated 2>/dev/null)" = "1" ]; then
    ARCHITECTURE="arm64" # Rosetta.
  else
    ARCHITECTURE="amd64" # Intel.
  fi
elif [ "${ARCHITECTURE}" = "arm64" ] ||[ "${ARCHITECTURE}" = "aarch64" ] ; then
  ARCHITECTURE="arm64" # Arm.
else
  ARCHITECTURE="amd64" # Amd.
fi

echo "platform="$PLATFORM,"arch="$ARCHITECTURE

BIN_URL="https://github.com/grassrootseconomics/ge-publish/releases/latest/download/ge-publish-${PLATFORM}-${ARCHITECTURE}.zip"

# Create the .ge-publish bin directory and ge-publish binary if it doesn't exist.
mkdir -p $GE_PUBLISH_BIN_DIR
tmp=$(mktemp -d 2>/dev/null)
curl -sSf -L "$BIN_URL" -o "$tmp/download.zip"
unzip -j "$tmp/download.zip" -d $tmp
mv "$tmp/ge-publish-${PLATFORM}-${ARCHITECTURE}" $BIN_PATH
chmod +x $BIN_PATH

# Store the correct profile file (i.e. .profile for bash or .zshenv for ZSH).
case $SHELL in
*/zsh)
    PROFILE="${ZDOTDIR-"$HOME"}/.zshenv"
    PREF_SHELL=zsh
    ;;
*/fish)
    PROFILE=$HOME/.config/fish/config.fish
    PREF_SHELL=fish
    ;;
*/bash)
    PROFILE=$HOME/.bashrc
    PREF_SHELL=bash
    ;;
*/ash)
    PROFILE=$HOME/.profile
    PREF_SHELL=ash
    ;;
*)
    echo "ge-publish: could not detect shell, manually add ${GE_PUBLISH_BIN_DIRECTORY} to your PATH."
    exit 1
esac

# Only add ge-publish if it isn't already in PATH.
if [[ ":$PATH:" != *":${GE_PUBLISH_BIN_DIR}:"* ]]; then
    # If the shell is fish, echo fish_add_path instead of export.
    if [[ "$PREF_SHELL" == "fish" ]]; then
        echo >> "$PROFILE" && echo "fish_add_path -a $GE_PUBLISH_BIN_DIR" >> "$PROFILE"
    else
        echo >> "$PROFILE" && echo "export PATH=\"\$PATH:$GE_PUBLISH_BIN_DIR\"" >> "$PROFILE"
    fi
fi

echo
echo "Detected your preferred shell is $PREF_SHELL and added ge-publish to PATH."
echo "Run 'source $PROFILE' or start a new terminal session to use ge-publish."
echo "Then, simply run 'ge-publish --version' to check the version ans start using it."

tolower() {
  echo "$1" | awk '{print tolower($0)}'
}
