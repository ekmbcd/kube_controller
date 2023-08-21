DOCKER_REGISTRY ?= localhost:5000
TARGET_OS ?= linux
DOCKER_TAG ?= netpol-controller
DEPLOYMENT_FILE ?= stuff/controller.yaml
PORT ?= :8001

all: compile build push apply

install:
	@printf ${INFO}"Installing dependencies"${END}
	@go mod download && \
		printf ${SUCCESS}"Dependencies installed successfully"${END} || \
		(printf ${ERROR}"Dependencies installation failed"${END} && exit 1)

compile:
	@GOOS=${TARGET_OS} go build -o ./app . && \
		printf ${SUCCESS}"Compiled successfully"${END} || \
		(printf ${ERROR}"Compilation failed"${END} && exit 1)

build:
	@printf ${INFO}"Building image"${END}
	@docker build -t ${DOCKER_TAG} . && \
		printf ${SUCCESS}"Image built successfully"${END} || \
		(printf ${ERROR}"Image build failed"${END} && exit 1)

push:
	@printf ${INFO}"Pushing image"${END}
	@docker tag ${DOCKER_TAG} ${DOCKER_REGISTRY}/${DOCKER_TAG} && \
		docker push ${DOCKER_REGISTRY}/${DOCKER_TAG} && \
		printf ${SUCCESS}"Pushed successfully"${END} || \
		(printf ${ERROR}"Push failed"${END} && exit 1)

apply:
	@printf ${INFO}"Applying deployment"${END}
	@kubectl apply -f ${DEPLOYMENT_FILE} --force && \
		printf ${SUCCESS}"Applied successfully"${END} || \
		(printf ${ERROR}"Apply failed"${END} && exit 1)

restart:
	@printf ${INFO}"Restarting deployment"${END}
	@kubectl rollout restart deployment ${DOCKER_TAG} && \
		printf ${SUCCESS}"Restarted successfully"${END} || \
		(printf ${ERROR}"Restart failed"${END} && exit 1)

run:
	@printf ${INFO}"Running locally"${END}
	@go run main.go -port ${PORT}

air:
	@printf ${INFO}"Running locally with air"${END}
	@air -- -port=${PORT}


BLK="\033[30m"
RED="\033[31m"
GRN="\033[32m"
YLW="\033[33m"
BLU="\033[34m"
MAG="\033[35m"
CYN="\033[36m"
WHT="\033[37m"
RST="\033[0m"
END="\n"${RST}

INFO=${CYN}
SUCCESS=${GRN}
ERROR=${RED}
WARNING=${YLW}
