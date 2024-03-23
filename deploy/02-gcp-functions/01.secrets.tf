resource "google_secret_manager_secret" "bot_cred_tg_secret" {

  count = length(local.bots_secrets_list)

  secret_id = "${terraform.workspace}-cred_multibot_tg_secret_${local.bots_secrets_list[count.index].name}"

  replication {
    auto {}
  }
}
resource "google_secret_manager_secret_version" "bot_cred_tg_secret_version" {

  count = length(local.bots_secrets_list)

  secret = google_secret_manager_secret.bot_cred_tg_secret[count.index].id

  secret_data = local.bots_secrets_list[count.index].secret
}




resource "google_secret_manager_secret" "bot_cred_tg_token" {

  count = length(local.bots_secrets_list)

  secret_id = "${terraform.workspace}-cred_multibot_tg_token_${local.bots_secrets_list[count.index].name}"

  replication {
    auto {}
  }
}
resource "google_secret_manager_secret_version" "bot_cred_tg_token_version" {

  count = length(local.bots_secrets_list)

  secret = google_secret_manager_secret.bot_cred_tg_secret[count.index].id

  secret_data = local.bots_secrets_list[count.index].token
}




resource "google_secret_manager_secret" "bot_cred_stripe_secret" {
  secret_id = "${terraform.workspace}-cred_multibot_stripe"

  replication {
    auto {}
  }
}
resource "google_secret_manager_secret_version" "bot_cred_stripe_secret_version" {
  secret = google_secret_manager_secret.bot_cred_stripe_secret.id

  secret_data = var.pp_stripe_token
}
resource "google_secret_manager_secret" "bot_cred_yoo_secret" {
  secret_id = "${terraform.workspace}-cred_multibot_yoo"

  replication {
    auto {}
  }
}
resource "google_secret_manager_secret_version" "bot_cred_yoo_secret_version" {
  secret = google_secret_manager_secret.bot_cred_yoo_secret.id

  secret_data = var.pp_yoo_token
}
