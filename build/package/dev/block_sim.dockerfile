# This dockerfile is meant to simulate a VPS for development purposes
FROM docker:dind

COPY . .
CMD bin/block_init