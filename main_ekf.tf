resource "helm_release" "elasticsearch" {
	name = "elasticsearch"
	repository = "https://helm.elastic.co"
	chart = "elasticsearch"

	namespace = "monitoring"

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

	namespace = "monitoring"
}

resource "helm_release" "metricbeat" {
	name = "metricbeat"
	repository = "https://helm.elastic.co"
	chart = "metricbeat"

	namespace = "monitoring"	
}

resource "helm_release" "fluentd" {
    name = "fluentd"
    chart = "./fluentd-chart"

	namespace = "monitoring"
}

resource "helm_release" "kibana-ingress" {
    name = "kibana-ingress"
    chart = "./kibana-ingress-chart"

        namespace = "monitoring"
}
