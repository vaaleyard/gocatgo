# Contributing

All contributions are welcome and I will be extremely grateful if you make it. You can help `gocatgo` by contributing with code, new ideas, telling me better ways of doing X, or anything! If you want to contribute with code, there are some things that still needs to be implemented (see [#todo](#todo) at the end).

## Running locally

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

## Pull Requests

1. Fork the project.
2. Make your changes locally.
3. Commit and create a Pull Request. In the text, for now, just explain what you did and that's it.
4. After you've created it, just wait a little, I will review it and merge it if everything is alright.

## TODO

1. Store images (with limit size)
2. Change some options to get from environment variables
3. Ability to create a paste with password, and be able to delete it if you have the password
4. Improve return code and messages
5. Provide the go binary to see the shasum
6. Make a docker image
