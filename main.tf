provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}

resource "helm_release" "http-service" {
    name    = "http-service"
    chart   = "../../http-service-chart"
    force_update = true
    set {
        name    = "service.type"
        value   = "ClusterIP"
    }

    reuse_values = true
}
