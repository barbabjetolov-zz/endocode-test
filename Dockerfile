FROM golang:alpine AS builder

# download necessary packages
RUN apk add git make

ENV GO111MODULE=on

WORKDIR /

# copy everything from project directory
COPY . .

# download dependencies
RUN go mod download


# compile binary
RUN make http-service
RUN chmod 111 http-service

# ENTRYPOINT [ "/http-service" ]

#-------------------#
FROM scratch

COPY --from=builder /http-service /http-service

ENTRYPOINT [ "/http-service" ]
