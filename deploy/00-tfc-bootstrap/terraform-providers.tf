# DOC: https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference
# gcloud auth application-default login
provider "google" {
  region = "us-west1"
}
data "google_client_config" "this" {}

# DOC: https://registry.terraform.io/providers/hashicorp/tfe/latest/docs
provider "tfe" {
  token = var.tfe_token
}
data "tfe_github_app_installation" "gha_installation" {
  installation_id = 45251287
}