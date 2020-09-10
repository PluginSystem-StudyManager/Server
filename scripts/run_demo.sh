#!/bin/bash

sudo apt update

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

# Server/scripts
cd DIR || {
  echo "ERROR: Directory " + DIR + "  not found"
  exit 1
}

# install Docker
sudo chmod +x install_docker
sudo ./install_docker

# Server
cd ..

# build server
sudo docker-compose build

# upload 5 dummy plugins
python3 scripts/mock/file_upload.py 5 --retry &

# Start the server
sudo docker-compose up
