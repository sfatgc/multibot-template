resource "tfe_variable_set" "global" {
  name         = "Global Varset"
  description  = "Global vars not depending on envs"
  global       = true
  organization = data.tfe_organization.main.name
}

resource "tfe_variable" "google_credentials" {
  key             = "GOOGLE_CREDENTIALS"
  sensitive       = true
  value           = var.google_credentials
  category        = "env"
  description     = "Google Cloud Service Account Credentials"
  variable_set_id = tfe_variable_set.global.id
}

resource "tfe_variable" "google_project" {
  key             = "GOOGLE_PROJECT"
  value           = var.google_project
  category        = "env"
  description     = "Google Cloud Project ID"
  variable_set_id = tfe_variable_set.global.id
}


resource "tfe_variable_set" "test" {
  name         = "Test Varset"
  description  = "Test envs vars"
  global       = true
  organization = data.tfe_organization.main.name
}

resource "tfe_variable" "pp_stripe_token_test" {
  key             = "pp_stripe_token"
  sensitive       = true
  value           = var.pp_stripe_token_test
  category        = "terraform"
  description     = "Stripe token (TEST)"
  variable_set_id = tfe_variable_set.test.id
}

resource "tfe_variable" "pp_yoo_token_test" {
  key             = "pp_yoo_token"
  sensitive       = true
  value           = var.pp_yoo_token_test
  category        = "terraform"
  description     = "YooKassa token (TEST)"
  variable_set_id = tfe_variable_set.test.id
}


resource "tfe_variable_set" "bots" {
  name         = "Bots creds"
  description  = "Bots credentials vars"
  global       = true
  organization = data.tfe_organization.main.name
}
resource "tfe_variable" "telegram_bots_list" {
  key             = "telegram_bots_list"
  sensitive       = false
  value           = var.telegram_bots_list
  category        = "terraform"
  description     = ""
  variable_set_id = tfe_variable_set.bots.id
}

resource "tfe_variable" "telegram_bot1_name" {
  key             = "telegram_bot1_name"
  sensitive       = false
  value           = var.telegram_bot1_name
  category        = "terraform"
  description     = ""
  variable_set_id = tfe_variable_set.bots.id
}
resource "tfe_variable" "telegram_bot1_secret" {
  key             = "telegram_bot1_secret"
  sensitive       = true
  value           = var.telegram_bot1_secret
  category        = "terraform"
  description     = ""
  variable_set_id = tfe_variable_set.bots.id
}
resource "tfe_variable" "telegram_bot1_token" {
  key             = "telegram_bot1_token"
  sensitive       = true
  value           = var.telegram_bot1_token
  category        = "terraform"
  description     = ""
  variable_set_id = tfe_variable_set.bots.id
}

resource "tfe_variable" "telegram_bot2_name" {
  key             = "telegram_bot2_name"
  sensitive       = true
  value           = var.telegram_bot2_name
  category        = "terraform"
  description     = ""
  variable_set_id = tfe_variable_set.bots.id
}
resource "tfe_variable" "telegram_bot2_secret" {
  key             = "telegram_bot2_secret"
  sensitive       = true
  value           = var.telegram_bot2_secret
  category        = "terraform"
  description     = ""
  variable_set_id = tfe_variable_set.bots.id
}
resource "tfe_variable" "telegram_bot2_token" {
  key             = "telegram_bot2_token"
  sensitive       = true
  value           = var.telegram_bot2_token
  category        = "terraform"
  description     = ""
  variable_set_id = tfe_variable_set.bots.id
}
