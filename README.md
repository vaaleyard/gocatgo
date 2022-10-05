<div align="center">

<h1> GoCatGo </h1>

<a href="https://gcg.sh">GoCatGo</a> is another pastebin tool with a super focus on transparency<br>

<br>

![lines of code](https://sloc.xyz/github/vaaleyard/gocatgo) ![Code Size](https://img.shields.io/github/languages/code-size/vaaleyard/gocatgo) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)

</div> 

## Requirements
It's a simple tool, you only need `curl`.

# Usage
```bash
# With a file
cat file.txt | curl -F "file=@-" gcg.sh
# or
curl -F "file=@file.txt" gcg.sh
```

```bash
# Passing any string
echo "some cool code" | curl -F "file=@-" gcg.sh
```

```bash
# Upload images
curl -F "file=@image.png" gcg.sh
```

## Privacy, Encryption and Transparency
Everything is encrypted. See [wiki](https://github.com/vaaleyard/gocatgo/wiki).

## Local Development

### Running locally

To run locally, a simple database is required (I'm using PostgreSQL, but since it's managed by an ORM, I think MySQL will also work).
This database is required because it's where all the user pastes are stored. It's not saved in the server disk like others programs do.
After creating the database, some environment variables are needed to run the program:

```bash
export GOCATGO_AES_KEY='passphrasewhichneedstobe32bytes!'
export DBHOST=your.database.hostname
export DBUSER=your.database.username
export DBPASSWORD=your.dabtabase.password
export DBNAME=your.dabtabase.name
export DBPORT=5432
```

Then, you can run the code:

```bash
go run .
# or
go build -o api && ./api
```

### Running via Docker

Ensure you have [Docker Compose](https://docs.docker.com/compose/) installed.

Run `docker compose build` to build the images.

Run `docker compose up` to create and start your containers.

## Contribution
See [CONTRIBUTING.md](./CONTRIBUTING.md)

## FAQ
1. How do I know the code running is the same as the repository?  
  **Answer**: I've created an URL so you can check the sha256 of current code running: gcg.sh/sha256

## License
[MIT](./LICENSE)
