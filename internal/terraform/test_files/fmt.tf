provider "google" {}

resource "google_storage_bucket" "image-store" {
  name = "image-store-bucket"
}
