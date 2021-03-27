# endocode-test

This repo contains: 
* `http-service`, a small application written in Golang that serves few HTTP endpoints
* scripts (jenkinsfile + helm chart + terraform) to automate the deployment of the application in an existing kubernetes cluster
* helm charts to deploy and EKF + metricbeat stack for logging and monitoring

# http-service

`http-service` is a simple golang application that serves HTTP requests. 

## OS Requirements
* macOs, Linux

## Building and running with docker
This method requires a machine with docker. To run the service in a docker container, simply run `make docker`: dockerd will automatically build an image from the provided Dockerfile and run it with the option `-d`. By default, it listens to the port 8080. This value can be changed by setting the `HOST_PORT` variable. For example, if one would want the service to isten to the port 8081, the command would be:

```shell
make docker HOST_PORT=8081
```

## local install
#### Requirements
* make
* golang 1.13 or newer
* git

To compile the source code run `make http-service`, the dependencies will be taken care of. To execute it, simply run `make run`; this will start the service, by default listening on the port 8080. To change it, export the variable `LISTENING_PORT` and set it equal to a port you'd liek the service to listen on. 

## Usage

The service accepts two endpoints:

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
```json
    {
        "git_commit": "2d23bd462aa5523a0bdcd272d4958700e3cc6eac",
        "project_name": "http-service"
    }
```

## Deploy
This service can be deployed in an existing kubernetes cluster using the provided jenkinsfile, it relies on terraform. The `main.tf` contains all the instructions terraform needs to deploy the service inside the kubernetes cluster; it's based on an helm chart, `http-service-chart` that deploys the service plus an nginx ingress, that allows communication with the service from outside the cluster. By default, the host name to contact the service with is `http-service.int`, but to make the connection possible, the line `$CLUSTER_IP http-service.int` has to be appended to the host `/etc/hosts` file (**N.B.** $CLUSTER_IP is a placeholder, it stands in for the actual cluster ip, obtainable with the command `kubectl cluster-info`). 

## Logging
There is another terraform file, `terraform/monitoring/main.tf`, it contains instructions to install an EKF stack + metricbeat in the kubernetes cluster. It's based on the official elastic helm charts for [Kibana](https://github.com/elastic/helm-charts/tree/6.5.2-alpha1/kibana), [Elasticsearch](https://github.com/elastic/helm-charts/tree/6.5.2-alpha1/elasticsearch) and [metricbeat](https://github.com/elastic/helm-charts/tree/master/metricbeat). They require the official repo to be addedd:

```shell
    helm repo add elastic https://helm.elastic.co
```

and on a custom charts for Fluentd, `fluentd-chart`. It deploys also an nginx ingress that allows acces to the Kibana UI, by default at the address `http://kibana.int`. To access it, add the line `$CLUSTER_IP kibana.int` in the host `/etc/hosts` file. 

## Server stub
A configuraton file to generate a server stub is provided, `openapi.yaml`. To generate the stub run the commands:

```shell
    git clone https://github.com/openapitools/openapi-generator
    cd openapi-generator
    mvn clean package
    java -jar modules/openapi-generator-cli/target/openapi-generator-cli.jar generate \
    -i openapi.yaml \
    -g $LANGUAGE \
    -o /var/tmp/php_api_client
```

There are several `$LANGUAGE` choices, check the [openapi-generator github page](https://github.com/OpenAPITools/openapi-generator) for a complete list.