FROM golang:alpine AS builder

# download necessary packages
RUN apk add git make

RUN unset GOPATH

# copy everything from project dir
COPY . .

# download dependencies
# RUN go mod download

# compile binary
RUN make compile

#-------------------#
FROM scratch

COPY --from=builder http-service /http-service

ENTRYPOINT [ "/http-service" ]