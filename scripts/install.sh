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
sudo apt install -y docker-compose

# clone repository
cd ~
mkdir StudyTool
cd StudyTool
git clone https://github.com/PluginSystem-StudyManager/Server
cd Server

git checkout dev # TODO: remove later on
sudo docker-compose up

# setup nginx
# TODO uncomment and fix if needed
#sudo echo """
#server {
#        listen 9090;
#        listen [::]:9090;
#
#        location / {
#                proxy_pass http://127.0.0.1:8080;
#        }
#}
#""" | sudo tee /etc/nginx/sites-available/study_tool
#
#sudo systemctl restart nginx
