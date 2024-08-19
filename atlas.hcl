// read internal/database/dialer/postgresql/README.md for more information

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/database/dialer/postgresql/table",
    "--dialect", "postgres"
  ]
}

// development environment
env "dev" {
  // pointing to gorm models to generate schemas
  src = data.external_schema.gorm.url
  // the database which is holding the actual tables
  url = "postgres://postgres:dev@localhost:5432/postgres?sslmode=disable"
  // a temporary database for Atlas to do migrations
  dev = "postgres://postgres:mig@localhost:5433/postgres?sslmode=disable"
#   dev = "docker://postgres/16"
  // location of migration files
  migration {
    dir = "file://internal/database/dialer/postgresql/migration?format=goose"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

// production environment
env "prod" {

}

lint {
  // https://atlasgo.io/lint/analyzers#data-dependent-changes
  data_depend {
    error = true
  }
  // https://atlasgo.io/lint/analyzers#backward-incompatible-changes
  incompatible {
    error = true
  }
  // https://atlasgo.io/lint/analyzers#destructive-changes
  destructive {
    error = true
  }
  naming {
    error   = true
    match   = "^[a-z]+$"
    message = "must be lowercase"
  }
}
