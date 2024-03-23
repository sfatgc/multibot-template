resource "google_firestore_database" "database" {
  name                    = "${terraform.workspace}-multibot-database"
  location_id             = local.region
  type                    = "FIRESTORE_NATIVE"
  delete_protection_state = "DELETE_PROTECTION_ENABLED"
  deletion_policy         = "DELETE"
}
