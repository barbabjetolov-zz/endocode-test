PROJECT_NAME=http-service
GIT_COMMIT=$(shell git rev-list -1 HEAD)

PKGDIR=pkg

TARGET_BIN=http-service
TARGET_PACKAGE=*

OUTER_PORT=8080
CONTAINER_PORT=8080

GOOS=linux
GOARCH=amd64

ENV_VARS=GOOS=${GOOS} GOARCH=${GOARCH}


compile: 
	env ${ENV_VARS} go build -ldflags "	-X github.com/barbabjetolov/endocode-test/http-service/pkg/utilities.ProjectName=${PROJECT_NAME} \
										-X github.com/barbabjetolov/endocode-test/http-service/pkg/utilities.GitCommit=${GIT_COMMIT}"	\
										-o ${TARGET_BIN}

run:
	./${TARGET_BIN}

install:
	go get

test:
	go test -ldflags "	-X github.com/barbabjetolov/endocode-test/http-service/pkg/utilities.ProjectName=${PROJECT_NAME} \
						-X github.com/barbabjetolov/endocode-test/http-service/pkg/utilities.GitCommit=${GIT_COMMIT} 	\
						-w -s" \
						./${PKGDIR}/${TARGET_PACKAGE} 

clean:
	go clean
	rm -rf ${TARGET_BIN}

all: test compile

docker-build:
	docker build -t ${TARGET_BIN} .

docker-run:
	docker run -dp ${CONTAINER_PORT}:${OUTER_PORT} -it ${TARGET_BIN}:latest

docker-clean:
	docker image prune

docker: docker-build docker-clean docker-run