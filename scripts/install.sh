#! /bin/bash

sudo apt update

# install Docker
sudo apt-get -y install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
sudo apt-get update
sudo apt-get -y install docker-ce docker-ce-cli containerd.io

# clone repository
cd ~
mkdir StudyTool
git clone https://github.com/PluginSystem-StudyManager/Server
cd StudyTool/Server

# create Docker container
docker build -t server:1.0 .

# run
docker run --name server -d -p 8080:8080 -it server:1.0
