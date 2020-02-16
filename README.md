## dbenv
- Parse `$DATABASE_URL`, output each variables to .env

## License
- Copyright (c) 2020 Tsuzu
- Under the MIT License

## Usage
- Download from [releases](https://github.com/cs3238-tsuzu/dbenv/releases)

```shell
# Make sure $DATABASE_URL is set
echo $DATABASE_URL

# Pass the output of dbenv to eval
eval "$(dbenv -)"

# Parsed envs will be set!
echo $DATABASE_HOST
```

### Supported envs
- From `SCHEME://USERNAME:PASSWORD@HOST:PORT/DB`

- `$DATABASE_USER`
- `$DATABASE_PASSWORD`
- `$DATABASE_SCHEME`
- `$DATABASE_HOST`
- `$DATABASE_PORT`
- `$DATABASE_DB`

- if `$DATABASE_PORT` is not present, guessed from `$DATABASE_SCHEME` (only available for MySQL and Postgres)

