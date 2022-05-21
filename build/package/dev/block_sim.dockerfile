# This dockerfile is meant to simulate a VPS for development purposes
FROM golang:latest as builder

WORKDIR /block_build
COPY . .

# Compile binaries
RUN make compile

FROM docker:dind

WORKDIR /block

# Copy compiled bins 
COPY --from=builder ./block_build /block
CMD bin/block_init