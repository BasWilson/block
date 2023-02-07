# Block 📦

## What and why
I'm planning for block to make it easy to deploy and manage **load balanced** services. 

Most cloud platforms are quite complicated to get setup for load balancing, and even more complicated to manage. I want to make it easy to deploy and manage load balanced services. Kind of like Digitalocean's app platform, but with load balancing. 

Block is the technical framework to make this happen, it is not a SaaS, PaaS or anything like that. Though in the future I might create a hosted version of Block, but that is not the goal of this project.

Join me on this journey of creating Block!

## Technical stack
My plan is to write the source completely in Go, I reckon it's the best way to learn a new language. 
There are also some shell scripts that are used to help with development and setup of servers. The master instance will have a sqlite database to store information. If I have time I plan on having a web interface called "Studio" for managing the master instance.

## Feature list and progress
A list of progress on the project and planned features that need to be made to make Block a reality.

### block_slave

- [x] VPS Simulator docker image
- [x] Block initialization
- [X] Image pulling
- [X] Image running
- [X] Running multiple images on one Block
- [X] Health monitoring
- [X] Ability to command the Block via HTTP
- [X] Ability to get logs from Block via HTTP
- [X] Ability to get list of running images on Block via HTTP
- [-] SSL Provisioning

**Nice to have features**
- [-] Git repo support
- [-] Builder images for building from source
- [-] Ability to store vps images in bucket


### block_master

- [X] Apply config
- [-] Adapters to manage slaves. (local, digitalocean, aws, gcp, azure etc)
- [-] Nginx HTTP load balancer 
- [-] Pre pulling images on slaves and then turning them off again
- [-] Caching slave images in nearest blob storage (s3, gcs, azure blob storage)
- [-] Get list of slaves
- [-] SSL Provisioning
- [-] Custom Domain support



## Figma Brainstorm file
https://www.figma.com/file/x5uF613CR5OjIn5VgdSkhE/Block?node-id=0%3A1
