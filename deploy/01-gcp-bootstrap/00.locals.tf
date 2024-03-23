locals {

  region      = data.google_client_config.this.region
  project     = data.google_client_config.this.project
  project_num = data.google_project.project.number
  project_id  = data.google_project.project.id

  services = [
    "iam.googleapis.com",
    "cloudfunctions.googleapis.com",
    "cloudresourcemanager.googleapis.com",
    "run.googleapis.com",
    "cloudbuild.googleapis.com",
    "secretmanager.googleapis.com",
    "firestore.googleapis.com",
  ]

}
