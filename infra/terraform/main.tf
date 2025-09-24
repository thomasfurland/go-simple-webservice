terraform {
  backend "gcs" {}
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

data "google_client_config" "default" {}

provider "kubernetes" {
  host                   = google_container_cluster.simple_web.endpoint
  token                  = data.google_client_config.default.access_token
  cluster_ca_certificate = base64decode(
    google_container_cluster.simple_web.master_auth[0].cluster_ca_certificate
  )
}

resource "google_container_cluster" "simple_web" {
  name            = "simple-web"
  location        = var.region
  enable_autopilot = true
}

resource "kubernetes_namespace" "simple_web" {
  metadata {
    name = "simple-web"
  }
}

resource "google_sql_database_instance" "postgres" {
  name             = "simple-web-db"
  database_version = "POSTGRES_15"
  region           = var.region

  settings {
    tier = "db-f1-micro"
  }

  deletion_protection = false
}

resource "google_sql_database" "default" {
  name     = "appdb"
  instance = google_sql_database_instance.postgres.name
}

resource "google_sql_user" "app" {
  name     = "appuser"
  instance = google_sql_database_instance.postgres.name
  password = var.db_password
}