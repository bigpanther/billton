data "external_schema" "billton" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/models",
    "--dialect", "postgres",
  ]
}

env "dev" {
  src = data.external_schema.billton.url
  dev = "postgres://postgres:postgres@localhost:5432/billton-atlas?sslmode=disable"
  url = "postgres://postgres:postgres@localhost:5432/billton-dev?sslmode=disable"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "test" {
  src = data.external_schema.billton.url
  dev = "postgres://postgres:postgres@localhost:5432/billton-atlas?sslmode=disable"
  url = "postgres://postgres:postgres@localhost:5432/billton-test?sslmode=disable"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "prod" {
  src = data.external_schema.billton.url
  dev = "postgres://postgres:postgres@localhost:5432/billton-atlas"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
