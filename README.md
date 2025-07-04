# diffinlist-htmx

## rerwiting diffinlist in htmx

**main goal of this app is to be the goat web scraper of online music to save them as a playlist**

## Todos

- v1
- [ ] model database
  - [ ] get the initial users and sessions table ready
  - [ ] test out the db
  - [ ] finish setting up playlists table
  - [ ] finish setting up songs table

## MakeFile

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB Container

```bash
make docker-down
```

DB Integrations Test:

```bash
make itest
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```
