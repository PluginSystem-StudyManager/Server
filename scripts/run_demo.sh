#!/bin/bash

sudo apt update


# clone repository
# TODO: use current directory as code
cd ~ || { echo "ERROR: Directory '~' not found"; exit 1; }
mkdir StudyTool
cd StudyTool || { echo "ERROR: Directory 'StudyTool' not found"; exit 1; }
git clone https://github.com/PluginSystem-StudyManager/Server
cd Server || { echo "ERROR: Directory 'Server' not found"; exit 1; }
git checkout dev # TODO: remove later on

# install Docker
sudo chmod +x .scripts/install_docker
./scripts/install_docker


# build server
sudo docker-compose build

# upload 5 dummy plugins
python3 scripts/mock/file_upload.py 5 --retry&

# Start the server
docker-compose up