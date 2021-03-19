PROJECT_NAME=http-service
GIT_COMMIT=$(shell git rev-list -1 HEAD)

PKGDIR=pkg

TARGET_BIN=http-service
TARGET_PACKAGE=*

OUTER_PORT=8080
CONTAINER_PORT=8080


GOOS=$(shell uname | tr '[:upper:]' '[:lower:]')
GOARCH=amd64

ENV_VARS=GOOS=${GOOS} GOARCH=${GOARCH}

# local install
clean:
	go clean
	rm -rf ${TARGET_BIN}

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


all: test compile


#docker
docker-build:
	docker build -t ${TARGET_BIN} .

docker-run:
	docker run -dp ${CONTAINER_PORT}:${OUTER_PORT} -it ${TARGET_BIN}:latest

docker-clean:
	-docker image rm ${TARGET_BIN}:latest

docker: docker-clean docker-build docker-run