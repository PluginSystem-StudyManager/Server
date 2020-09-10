#! /bin/bash

sudo apt update

# install Docker
sudo chmod +x ./install_docker
./install_docker

# clone repository
cd ~
mkdir StudyTool
cd StudyTool
git clone https://github.com/PluginSystem-StudyManager/Server
cd Server

git checkout dev # TODO: remove later on

# build and run server
sudo docker-compose up
