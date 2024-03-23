resource "random_id" "resources_id" {
  byte_length = 8
}

resource "google_storage_bucket" "cf-source-bucket" {
  name                        = "sfatgc-multibot-gcf-source-${random_id.resources_id.hex}" # Every bucket name must be globally unique
  location                    = "US"
  uniform_bucket_level_access = true
}

data "archive_file" "cf-http-source-zip" {
  type        = "zip"
  output_path = "/tmp/${terraform.workspace}-sfatgc-multibot-gcf-source-${local.functions_filenames_timestamp}.zip"
  source_dir  = "../../functions/multibot/"
}
resource "google_storage_bucket_object" "cf-http-object" {
  name   = "${terraform.workspace}-cf-source-multibot-${local.functions_filenames_timestamp}.zip"
  bucket = google_storage_bucket.cf-source-bucket.name
  source = data.archive_file.cf-http-source-zip.output_path
}

resource "google_cloudfunctions2_function" "cf_http_multibot" {

  name        = "multibot-${terraform.workspace}"
  location    = "us-west1"
  description = "TG Multibot Cloud Function (${terraform.workspace} env)"

  build_config {
    runtime     = "go122"
    entry_point = "entrypoint" # Set the entry point
    source {
      storage_source {
        bucket = google_storage_bucket.cf-source-bucket.name
        object = google_storage_bucket_object.cf-http-object.name
      }
    }
  }

  service_config {
    service_account_email = google_service_account.multibot_sa.email
    max_instance_count    = 1
    available_memory      = "256M"
    timeout_seconds       = 60
    ingress_settings      = "ALLOW_ALL"

    environment_variables = {
      "GOOGLE_PROJECT_ID"      = split("/", data.google_project.project.id)[1]
      "GOOGLE_FIRESTORE_DB_ID" = element(local.google_firestore_db_id, length(local.google_firestore_db_id) - 1)
      "TELEGRAM_BOTS_LIST"     = var.telegram_bots_list
    }


    dynamic "secret_environment_variables" {

      for_each = local.bs_tg_secrets

      content {
        key        = "TELEGRAM_BOT_SECRET_${upper(secret_environment_variables.value.name)}"
        project_id = data.google_project.project.project_id
        secret     = secret_environment_variables.value.sid
        version    = secret_environment_variables.value.sv
      }
    }

    dynamic "secret_environment_variables" {

      for_each = local.bs_tg_tokens

      content {
        key        = "TELEGRAM_BOT_TOKEN_${upper(secret_environment_variables.value.name)}"
        project_id = data.google_project.project.project_id
        secret     = secret_environment_variables.value.sid
        version    = secret_environment_variables.value.sv
      }
    }


    secret_environment_variables {
      key        = "PP_STRIPE_TOKEN"
      project_id = data.google_project.project.project_id
      secret     = google_secret_manager_secret.bot_cred_stripe_secret.secret_id
      version    = google_secret_manager_secret_version.bot_cred_stripe_secret_version.version
    }

    secret_environment_variables {
      key        = "PP_YOO_TOKEN"
      project_id = data.google_project.project.project_id
      secret     = google_secret_manager_secret.bot_cred_yoo_secret.secret_id
      version    = google_secret_manager_secret_version.bot_cred_yoo_secret_version.version
    }
  }

  depends_on = [
    google_secret_manager_secret_version.bot_cred_tg_token_version,
    google_secret_manager_secret_version.bot_cred_tg_secret_version,
    google_secret_manager_secret_version.bot_cred_stripe_secret_version,
    google_secret_manager_secret_version.bot_cred_yoo_secret_version
  ]

}

resource "google_cloud_run_service_iam_member" "member" {
  location = google_cloudfunctions2_function.cf_http_multibot.location
  service  = google_cloudfunctions2_function.cf_http_multibot.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}

output "function_uri" {
  value = google_cloudfunctions2_function.cf_http_multibot.url
}
