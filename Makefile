DOCKER_SIM_TAG=baswilson/block_sim:latest 
DOCKER_SIM_BUILD_FLAGS=--build-arg BLOCK_USER=${BLOCK_USER} \
					--tag ${DOCKER_SIM_TAG} \
					--file build/package/dev/block_sim.dockerfile
DOCKER_SIM_RUN_FLAGS=--privileged \
					--network=block_network \
					-v /var/run/docker.sock:/var/run/docker.sock \
					-p 8080:8080 \
					--env-file build/ci/dev/sim.env

clean:
	go clean
	rm -r ./bin

compile: clean
	mkdir ./bin
	go build -o ./bin ./cmd/...

run: clean compile
	./bin/slave

build_sim:
	@echo "Building Block simulation image"
	@docker build ${DOCKER_SIM_BUILD_FLAGS} .

run_sim:
	@echo "Running Block simulation image"
	@docker network create block_network || true
	@docker run ${DOCKER_SIM_RUN_FLAGS} ${DOCKER_SIM_TAG}

sim: compile build_sim run_sim