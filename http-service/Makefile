PROJECT_NAME=$(shell basename `git rev-parse --show-toplevel`)
GIT_COMMIT=$(shell git rev-list -1 HEAD)

OBJDIR=bin
PKGDIR=pkg

TARGET_BIN=http-service
TARGET_PACKAGE=*


compile: 
	go build -ldflags "	-X github.com/barbabjetolov/endocode-test/http-service/pkg/utilities.ProjectName=${PROJECT_NAME} \
						-X github.com/barbabjetolov/endocode-test/http-service/pkg/utilities.GitCommit=${GIT_COMMIT}"	\
						-o ./${OBJDIR}/${TARGET_BIN}

run:
	./${OBJDIR}/${TARGET_BIN}

install:
	go get

test:
	go test -ldflags "	-X github.com/barbabjetolov/endocode-test/http-service/pkg/utilities.ProjectName=${PROJECT_NAME} \
						-X github.com/barbabjetolov/endocode-test/http-service/pkg/utilities.GitCommit=${GIT_COMMIT}" 	\
						./${PKGDIR}/${TARGET_PACKAGE} 

clean:
	go clean
	rm -rf ./${OBJDIR}

all: test compile