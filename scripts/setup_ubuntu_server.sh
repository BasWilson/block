#!/bin/bash

# Install docker
sudo snap install docker

# install make
sudo apt-get install make

# Install go
sudo snap install go --classic

# Clone repo
git clone https://github.com/baswilson/block.git /root/block

# Add go script
sudo echo '#!/bin/bash
export GOPATH=/root/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' > /etc/profile.d/go.sh

# Make go script executable
sudo chmod +x /etc/profile.d/go.sh

# Add go path to bashrc
sudo echo 'export GOPATH=/root/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> ~/.bashrc

# Add rc.local script
sudo echo '#!/bin/bash
/root/block/scripts/slave/compile.sh
/root/block/scripts/slave/start.sh' > /etc/rc.local

# Make executables
sudo chmod +x /etc/rc.local
sudo chmod +x /root/block/scripts/slave/compile.sh
sudo chmod +x /root/block/scripts/slave/start.sh

# Compile
/root/block/scripts/slave/compile.sh

# Reboot
sudo reboot