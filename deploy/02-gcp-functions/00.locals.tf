locals {

  region      = data.google_client_config.this.region
  project     = data.google_client_config.this.project
  project_num = data.google_project.project.number
  project_id  = data.google_project.project.id

  bots_secrets_list = [
    { name = var.telegram_bot1_name, secret = var.telegram_bot1_secret, token = var.telegram_bot1_token },
    { name = var.telegram_bot2_name, secret = var.telegram_bot2_secret, token = var.telegram_bot2_token }
  ]

  bs_tg_secrets = [for index, bsi in local.bots_secrets_list : { name = bsi.name, sid = google_secret_manager_secret.bot_cred_tg_secret[index].secret_id, sv = google_secret_manager_secret_version.bot_cred_tg_secret_version[index].version }]
  bs_tg_tokens  = [for index, bsi in local.bots_secrets_list : { name = bsi.name, sid = google_secret_manager_secret.bot_cred_tg_token[index].secret_id, sv = google_secret_manager_secret_version.bot_cred_tg_token_version[index].version }]

  functions_filenames_timestamp = formatdate("ZZZhhmmDDMMMYYYY", timestamp())

  google_firestore_db_id = split("/", google_firestore_database.database.id)

}
