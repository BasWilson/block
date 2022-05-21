# This dockerfile is meant to simulate a VPS for development purposes
FROM docker:dind

COPY ./init/block.service /etc/systemd/system/block.service

COPY . .
CMD bin/block_init