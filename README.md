## Block 📦

### But why
I'm planning for block to be a lightweight platform-as-a-service framework. 

This idea came to me because of my curiosity for the cloud, scaling, automation and my love for projects that are way to ambitious for a single developer. 
I think services such as DigitalOcean's App Platform, Azure's App Service and Platform.sh etc are a great feat of engineering and I wish to learn at least a little more of how this is done.

Join me on this journey of creating Block!

### Technical stack
The project has only just started, but my plan is to write most of the source code with GoLang, this way I can improve my GoLang skillset. I've been using JavaScript/TypeScript/NodeJs religiously for the past 5 years and I thinks it's time for a change.

### Scripts
I decided the best way to improve my Go skills is to write *everything* in Go, inluding helper scripts. Some scripts would definetly have made more sense to be written for bash.

### Feature list and progress
A list of progress on the project and planned features that need to be made to make Block a reality.

#### Block
Core functionality of a Block such as services that monitor, setup and log what happens on a Block.

- [x] VPS Simulator docker image
- [x] Block initialization
- [X] Image pulling
- [X] Image running
- [-] Deamon monitoring service

**Nice to have features**
- [-] Git repo support
- [-] Builder images for building from source

#### API
The API will manage authentication of users, setup Block instances and all other operations you would expect to see from a Platform as a Service API.

#### Web app
The web app will be a dashboard where users will manage their Blocks.

### Figma Brainstorm file
https://www.figma.com/file/x5uF613CR5OjIn5VgdSkhE/Block?node-id=0%3A1