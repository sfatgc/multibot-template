data google_client_config "this" {}
data google_project "project" {}

locals {    
  region      = data.google_client_config.this.region
  project     = data.google_client_config.this.project
  project_num = data.google_project.project.number
  project_id  = data.google_project.project.id
}
