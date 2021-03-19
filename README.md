# endocode-test

This repo contains:

* http-service


# http-service

`http-service` is a simple golang application that serves HTTP requests. 

## Requirements

#### OS requirements
* macOs, Linux

#### with docker
* docker
#### local install
* go 1.13
* make

## Building and running

#### docker
To build a docker image simply run `make docker`: this will build and run a container that contains the application. By default it listens from the 8080 port.

#### local install
To compile the source code run `make compile`, it will take care of the dependencies.  To run it, launch `make run`. As noted before, by default the service listens from the 8080 port, but it can be changed by setting the environment variable `LISTENING_PORT`.

## Usage

The service listens accepts two endpoints:

* **GET /helloworld** - returns "Hello Stranger". It accepts one query parameter, `name`. If set, it returns "Hello $name"
* **GET /versionz** - returns a JSON with the hash of the latest commit and the project name
