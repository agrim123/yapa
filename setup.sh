#! /bin/bash

set -e

cd ~/.yapa/yapa && glide install && go build . && cd -

echo "alias yapa=\"~/.yapa/yapa/yapa\"" >> ~/.zshrc
