# Network Monitor

## Prerequisites

- Go 1.21.1 ([Installation](https://go.dev/doc/install))
- Atlas ([Installation](https://atlasgo.io/getting-started/))
- Templ ([Installation](https://templ.guide/quick-start/installation))

### Switching environments

1. Assign the desired env file to `.env` before running the makefile
   ```shell
   cp local.env .env
   ```
2. Update any credentials locally in `.env`
3. Source and export the env variables
   ``` bash
   set -a
   source .env
   set +a
   ```

## Database migrations

### Update schema

1. Update DDL statements in `db/schema.sql`
2. Use Atlas to create versioned migration
    ```shell
    atlas migarte diff -c file://db/atlas.hcl --env local <migration_name>
    ```

### Run migrations

```shell
atlas migrate apply -c file://db/atlas.hcl --env local
```