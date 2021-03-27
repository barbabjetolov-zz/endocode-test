# Kibana ingress chart
Custom chart to deploy an ingress to access the Kibana UI from outside the kubernetes cluster.

## Configuration
| **Parameter**       | **Description**                        | **Default value** |
| ----------------    | -------------------------------------- | ----------------- |
|`ingress.name`       | name of the deployment                 | `kibana-ingress`  |
|`ingress.hostName`   | host name                              | `kibana.int`      |
|`ports.containerPort`| port from which the UI can be accessed | `5601`            |