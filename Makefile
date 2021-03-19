PROJECT_NAME=http-service
GIT_COMMIT=$(shell git rev-list -1 HEAD)

PKGDIR=pkg

TARGET_BIN=http-service
TARGET_PACKAGE=*

HOST_PORT=8080
CONTAINER_PORT=8080

TAG=latest


GOOS=$(shell uname | tr '[:upper:]' '[:lower:]')
GOARCH=amd64

ENV_VARS=GOOS=${GOOS} GOARCH=${GOARCH}

# local install
clean:
	go clean
	-rm -rf ${TARGET_BIN}

install: clean
	go mod download

compile: install
	env ${ENV_VARS} go build -ldflags "	-X github.com/barbabjetolov/endocode-test/http-service/pkg/utilities.ProjectName=${PROJECT_NAME} \
										-X github.com/barbabjetolov/endocode-test/http-service/pkg/utilities.GitCommit=${GIT_COMMIT}"	\
										-o ${TARGET_BIN}

run:
	./${TARGET_BIN}

test:
	go test -ldflags "	-X github.com/barbabjetolov/endocode-test/http-service/pkg/utilities.ProjectName=${PROJECT_NAME} \
						-X github.com/barbabjetolov/endocode-test/http-service/pkg/utilities.GitCommit=${GIT_COMMIT} 	\
						-w -s" \
						./${PKGDIR}/${TARGET_PACKAGE} 


all: test compile run


#docker
docker-build:
	docker build -t ${TARGET_BIN} .

docker-run:
	docker run -d --name ${TARGET_BIN} -e LISTENING_PORT=${CONTAINER_PORT} -dp ${HOST_PORT}:${CONTAINER_PORT} -it ${TARGET_BIN}:${TAG}

docker-clean:
	-docker image rm ${TARGET_BIN}:${TAG}

docker: docker-build docker-run