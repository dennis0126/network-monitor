variable "url" {
  type        = string
  description = "DB string"
  default     = getenv("DB_STRING")
}

variable "dev" {
  type        = string
  description = "DB string dev"
  default     = getenv("DB_STRING_DEV")
}

env {
  name = atlas.env
  src  = "file://db/schema.sql"
  url  = var.url
  dev  = var.dev
  migration {
    dir = "file://db/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
  log {
    schema {
      apply = "env: ${atlas.env}"
    }
  }
}