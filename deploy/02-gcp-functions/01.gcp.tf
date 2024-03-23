provider "google" {
  region = "us-west1"
}

data "google_project" "project" {}
data "google_client_config" "this" {}

resource "google_project_service" "service" {
  count   = length(local.services)
  service = local.services[count.index]

  timeouts {
    create = "30m"
    update = "40m"
  }

  disable_dependent_services = true
  disable_on_destroy         = true
}
