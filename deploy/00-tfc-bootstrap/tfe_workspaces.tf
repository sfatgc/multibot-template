# https://github.com/settings/installations/45251287

resource "tfe_workspace" "gcp-bootstrap" {
  name              = "gcp-bootstrap"
  organization      = data.tfe_organization.main.name
  auto_apply        = true
  queue_all_runs    = true
  working_directory = "deploy/01-gcp-bootstrap"
  trigger_patterns  = ["deploy/01-gcp-bootstrap/*"]
  vcs_repo {
    branch                     = "main"
    identifier                 = "sfatgc/multibot"
    github_app_installation_id = data.tfe_github_app_installation.gha_installation.id #"45251287"
  }
}

locals {
  gcf_envs = {
    "dev"  = {}
    "test" = {}
    "prod" = {}
  }
}

resource "tfe_workspace" "gcp-functions" {
  for_each          = local.gcf_envs
  name              = "${each.key}-gcp-functions"
  organization      = data.tfe_organization.main.name
  auto_apply        = true
  queue_all_runs    = true
  working_directory = "deploy/02-gcp-functions"
  trigger_patterns  = ["deploy/02-gcp-functions/*", "functions/**/*"]
  vcs_repo {
    branch                     = "main"
    identifier                 = "sfatgc/multibot"
    github_app_installation_id = data.tfe_github_app_installation.gha_installation.id #"45251287"
  }
}
