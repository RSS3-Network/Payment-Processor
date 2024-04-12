# Migration with Atlas

`atlas.hcl` is the configuration file for Atlas.
Think it as a terraform but for databases.
It contains the configurations and environments for:
1. database connection
2. where to store migration files
3. where to look for schema files (gorm models)
4. other configurations

## Create a migration
After modifying any gorm models, create a migration file against your modifications.
```
    atlas migrate diff migration_name --env dev
```

## Apply the migration
Apply the migration to any database via `--url` flag.
```
    atlas migrate apply --env dev
```

## Commit the migration
Make a dedicated commit for the migration to the repository.

For more information, please refer to:
1. [Atlas Quick Introduction](https://atlasgo.io/getting-started/)
2. [Automatic migration planning for GORM](https://atlasgo.io/guides/orms/gorm)
