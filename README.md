# endocode-test

This repo contains a small application, `http-service`, written in Golang that serves few HTTP endpoints, and the scripts (jenkinsfile + helm chart + terraform plan) needed to automate its deployment in an existing kubernetes cluster.

# http-service

`http-service` is a simple golang application that serves HTTP requests. 

## Requirements

#### OS requirements
* macOs, Linux

#### with docker
* docker
#### local install
* go 1.13 or newer
* make

## Building and running

#### docker
To build a docker image simply run `make docker`: this will build and run a container that contains the application. By default it listens from the 8080 port.

#### local install
To compile the source code run `make compile`, it will take care of the dependencies.  To run it, launch `make run`. As noted before, by default the service listens from the 8080 port, but it can be changed by setting the environment variable `LISTENING_PORT`.

## Usage

The service listens accepts two endpoints:

* **GET /helloworld** - returns "Hello Stranger". It accepts one query parameter, `name`. If set, it returns "Hello $name", sliced by camel case
* **GET /versionz** - returns a JSON with the hash of the latest commit and the project name

For example, to call the first endpoint, you'd do: 
```shell
    curl localhost:8080/helloworld
```

and the answer will be:
```
    Hello Stranger
```

To use the query funcionality, use this request
```shell
    curl localhost:8080/helloworld?name=MarcoRossi
```

and the answer will be:
```
    Hello Marco Rossi
```

The last endpoint can be contacted with
```
    curl localhost:8080/versionz
```

and the answer will be:
```
    {"git_commit":"2d23bd462aa5523a0bdcd272d4958700e3cc6eac","project_name":"http-service"}
```

## Deploy
This service can be deployed in an existing kubernetes cluster using the provided jenkinsfile. The cluster IP has to be set manually in the `main.tf` file. This plan also requires certificates to authenticate to the cluster.
The helm chart contains templates to create a service, a deployment and an nginx-ingress that allows communication with the service. 