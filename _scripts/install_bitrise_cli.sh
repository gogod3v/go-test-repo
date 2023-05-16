#!/usr/bin/env bash
set -ex

export BITRISE_CLI_VERSION='2.2.5'

curl -fL https://github.com/bitrise-io/bitrise/releases/download/$BITRISE_CLI_VERSION/bitrise-Linux-x86_64 > /tmp/bitrise
sudo mv /tmp/bitrise /usr/local/bin/bitrise
chmod +x /usr/local/bin/bitrise
bitrise setup

