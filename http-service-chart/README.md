# http-service helm chart
Custom chart to deploy `http-service` and relative ingress to the kubernetes cluster. 

## Configuration
| **Parameter**       | **Description**                         | **Default value**       |
| ----------------    | --------------------------------------- | ----------------------- |
|`service.type`       |**MANDATORY** - kubernetes service type (e.g. ClusterIP)||
|`deploymentName`     |name of the deployment                  | `http-service`           |
|`replicaCount`       |base number of replicas                 | `1`                      |
|`image.repository`   |repository from which to pull the image | `erizzardi/http-service` |
|`image.pullPolicy`   |policy according to which the docker image is pulled from the repository | `IfNotPresent`|
|`ports.containerPort`|the port exposed by the pod               |`8080`            |
|`ports.listeningPort`|the port from which the service listens to|`8080`            |
|`ingress.name`       |name of the nginx ingress                 |`nginx-ingress`   |
|`ingress.hostName`   |name of the nginx virtual server          |`http-service.int`|
|`resources.limits.cpu`   |maximum cpu cycles allowed for a pod |`100m` |
|`resources.limits.memory`|maximum memory allowed for a pod     |`128Mi`|