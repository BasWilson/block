#!/bin/bash

# Install docker
sudo snap install docker

# install make
sudo apt-get install make

# Add rc.local script
sudo echo '#!/bin/bash
git -C /root/block checkout rework
git -C /root/block pull
make -C /root/block compile 
/root/block/bin/slave' > /etc/rc.local

# Make rc.local executable
sudo chmod +x /etc/rc.local

# Clone repo
git clone https://github.com/baswilson/block.git /root/block

# Reboot self
sudo reboot