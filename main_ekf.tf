provider "helm" {
  kubernetes {
    host                   = "https://192.168.49.2:8443"

    client_certificate     = file("client.crt")
    client_key             = file("client.key")
    cluster_ca_certificate = file("ca.crt")
  }
}

resource "helm_release" "elasticsearch" {
	name = "elasticsearch"
	repository = "https://helm.elastic.co"
	chart = "elasticsearch"

	namespace = "logging"

	set {
		name = "replicas"
		value = "1"
	}

	set {
		name = "minimumMasterNodes"
		value = "1"
	}
}

resource "helm_release" "kibana" {
	name = "kibana"
	repository = "https://helm.elastic.co"
	chart = "kibana"

	namespace = "logging"
}

resource "helm_release" "fluentd" {
    name = "fluentd"
    chart = "./fluentd-chart"
}