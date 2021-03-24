provider "helm" {
  kubernetes {
    host                   = "https://192.168.49.2:8443"
    config_path = "~/.kube/config"
  }
}

resource "helm_release" "http-service" {
    name    = "http-service"
    chart   = "./http-service-chart"
    force_update = true
    set {
        name    = "service.type"
        value   = "ClusterIP"
    }

    reuse_values = true
}
