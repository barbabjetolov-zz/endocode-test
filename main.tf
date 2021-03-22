
provider "helm" {
  kubernetes {
    host                   = "https://192.16849.2:8443"

    client_certificate     = file("/client.crt")
    client_key             = file("/client.key")
    cluster_ca_certificate = file("/ca.crt")
  }
}

resource "helm_release" "complete" {
    name    = "http-request"
    chart   = "/http-request-chart"

    set {
        name    = "service.type"
        value   = "ClusterIP"
    }
}