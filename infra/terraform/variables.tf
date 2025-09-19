variable "project_id" {
  description = "Your GCP project ID"
  type        = string
}

variable "region" {
  description = "GCP region for the cluster"
  type        = string
}

variable "db_password" {
  description = "App DB user password"
  type        = string
  sensitive   = true
}