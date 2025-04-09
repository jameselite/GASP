package configfile

var SqlcYamlPG string = `

version: "2"
sql:
  - engine: "postgresql"
    queries: "query.sql"
    schema: "schema.sql"
    gen:
      go:
        package: "%s"
        out: "%s"
        sql_package: "pgx/v5"

`

var SqlcYamlMYSQL string = `

version: "2"
sql:
  - engine: "mysql"
    queries: "query.sql"
    schema: "schema.sql"
    gen:
      go:
        package: "%s"
        out: "%s"

`