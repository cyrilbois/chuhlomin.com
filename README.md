# micro

[![main](https://github.com/chuhlomin/micro/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/chuhlomin/micro/actions/workflows/main.yml) [![release](https://github.com/chuhlomin/micro/actions/workflows/release.yml/badge.svg)](https://github.com/chuhlomin/micro/actions/workflows/release.yml)

Microblog, stored as Markdown files in Git repository.

GitHub Actions uses [genblog](https://github.com/chuhlomin/genblog) Go app that generates static site using `templates` directory.

Powers https://chuhlomin.com/blog

## Local setup

Set environment variables values from `.env` file.
You may use [alias e](https://github.com/chuhlomin/e) for that.

```bash
e
```

Use [genblog](https://github.com/chuhlomin/genblog) binary to generate static
site from Markdown files and templates. Also copy files from `static` directory
to `output`

```bash
genblog
cp -R static/ output/

# same as:
make build
```

Use [fswatch](https://github.com/emcrisostomo/fswatch) to update the site on every file change.

```bash
# brew install fswatch
fswatch -or -e "output" -e ".git" . | xargs -n1 sh -c "genblog; cp -R static/ output/"

# same as:
fswatch --one-per-batch --recursive --exclude="output" --exclude=".git" . | xargs -n1 sh -c "genblog; cp -R static/ output/"

# same as:
make watch
```

Serve output folder locally via Nginx using docker-compose:

```bash
docker-compose up -d nginx
# or
make run-docker
```

Open http://127.0.0.1:8080
