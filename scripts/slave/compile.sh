#!/bin/bash
export HOME=/root
export GOPATH=/root/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
export GOCACHE=/root/go/cache

git -C /root/block checkout rework
git -C /root/block pull
make -C /root/block compile
echo "Compiled slave"