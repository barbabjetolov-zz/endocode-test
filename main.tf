provider "helm" {
  kubernetes {
    host                   = "https://192.168.49.2:8443"

    client_certificate     = file("client.crt")
    client_key             = file("client.key")
    cluster_ca_certificate = file("ca.crt")
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