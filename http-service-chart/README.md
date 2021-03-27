# http-service helm chart
Custom chart to deploy `http-service` and relative ingress to the kubernetes cluster. 

## Configuration
| **Parameter**       | **Description**                         | **Default value**       |
| ----------------    | --------------------------------------- | ----------------------- |
|`service.type`       |**MANDATORY** - kubernetes service type (e.g. ClusterIP)||
|`deploymentName`     |name of the deployment                  | `http-service`           |
|`replicaCount`       |base number of replicas                 | `1`                      |
|`image.repository`   |repository from which to pull the image | `erizzardi/http-service` |
|`image.pullPolicy`