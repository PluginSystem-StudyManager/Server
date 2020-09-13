#!/bin/bash

INSTALL=true

while true; do
  case "$1" in
  -n | --noinstall)
    INSTALL=false;
    shift
    ;;
  -h | --help)
    printf "Installs everything that is needed and starts the server.\n\n\t-n, --noinstall: To run it without installing and building everything again";
    exit 0
    ;;
  *) break ;;
  esac
done

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

# Server/scripts
cd "$DIR" || {
  echo "ERROR: Directory " + "$DIR" + "  not found"
  exit 1
}

# Server
cd ..

if [ "$INSTALL" = true ]; then

  sudo apt update

  # install Docker
  sudo chmod +x ./scripts/install_docker
  sudo ./scripts/install_docker

  # build server
  sudo docker-compose build
fi

# configure nginx
original="try_files \$uri \$uri\/ =404;"
new="proxy_pass http:\/\/127.0.0.1:8080;"
sudo sed -i "s/$original/$new/g" /etc/nginx/sites-available/default

# upload 5 dummy plugins
python3 ./scripts/mock/file_upload.py 5 --retry &

# Start the server
sudo docker-compose up

# Undo the changes for nginx to restore previous state
sudo sed -i "s/$new/$original/g" /etc/nginx/sites-available/default
