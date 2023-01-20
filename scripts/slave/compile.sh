#!/bin/bash
git -C /root/block checkout rework
git -C /root/block pull
make -C /root/block compile 