# fluentd helm chart
Custom chart to deploy `fluentd` in the kubernetes cluster. The default setting are compliant with the ones for the other elements in the EFK stack, referenced in the terraform file, that is everything will work out of the box. 

## Configuration
| **Parameter**       | **Description**                         | **Default value**       |
| ----------------    | --------------------------------------- | ----------------------- |
|`deploymentName`|name of the deployment|`fluentd`|
|`namespace`|kubernetes namespace to deploy fluentd to|`monitoring`|
|`terminationGracePeriodSeconds`|grace period for pod termination, in seconds|`30`|
|`daemonSet.env.elasticsearchHost`|elasticsearch host to connect to|`elasticsearch-master.monitoring.svc.cluster.local`|
|`daemonSet.env.elasticsearchPort`|elasticsearch port for logs ingestion|`9200`|
|`daemonSet.env.elasticsearchScheme`|communication scheme with wlasticsearch|`http`|