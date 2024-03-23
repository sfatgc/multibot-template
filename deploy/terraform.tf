terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 4.34.0"
    }
    tfe = {
      source  = "hashicorp/tfe"
      version = "0.53.0"
    }
  }
}
