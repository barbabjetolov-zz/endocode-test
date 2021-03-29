# Jenkinsfile and jenkins configuration

## Requirements

#### Jenkins plugins
* pipeline
* git
* docker-pipeline

#### Caveats
The pipeline won't run as-is. It requires credentials for dockerhub to push the image. The docker-pipeline plugin also needs access to a docker client, thus, if containerized, has to be provided inside the image, and the docker.sock has to be mounted in.

The non-regression tests step could fail, because jenkins has to resolve the host "http-service.int". To do so, add `$CLUSTERIP http-service.int` to the /etc/hosts file jenkis reads from. 