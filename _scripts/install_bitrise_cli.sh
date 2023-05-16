#!/usr/bin/env bash
set -ex

export BITRISE_CLI_VERSION='2.2.5'

curl -fL https://github.com/bitrise-io/bitrise/releases/download/$BITRISE_CLI_VERSION/bitrise-"$(uname -s)"-"$(uname -m)" > /tmp/bitrise
sudo mv /tmp/bitrise /usr/local/bin/bitrise
chmod +x /usr/local/bin/bitrise
bitrise setup

