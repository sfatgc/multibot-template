resource "google_service_account" "multibot_sa" {
  account_id   = "gcf-multibot-sa"
  display_name = "MultiBot CF Service Account"
}

resource "google_project_iam_binding" "bot_secrets_access" {
  project = data.google_project.project.id
  role    = "roles/secretmanager.secretAccessor"

  condition {
    title       = "multibot_secrets_only"
    description = "Allows access only to the desired secrets"
    expression  = "resource.name.startsWith(\"projects/${local.project_num}/secrets/cred_multibot_\")"
  }

  members = [
    "serviceAccount:${google_service_account.multibot_sa.email}"
  ]
}

resource "google_project_iam_binding" "bot_firestore_access" {
  project = data.google_project.project.id
  role    = "roles/datastore.user"

  condition {
    title       = "multibot_db_only"
    description = "Allows access only to the desired DB"
    expression  = "resource.name.startsWith(\"projects/${local.project_num}/databases/multibot\")"
  }

  members = [
    "serviceAccount:${google_service_account.multibot_sa.email}"
  ]
}
