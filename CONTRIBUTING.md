## Contributing
All contributions are welcome and I will be extremely grateful if you make it. You can help gocatgo by contributing with code, new ideas, telling me better ways of doing X, anything! If you want to contribute with code, there are some things that still needs to be implemented (see #todo in the end).

### Running locally
To run locally, a simple database is required: MySQL (you can run it through docker compose).
This database is required because it's where all the user pastes are stored. It's not saved in the server disk like others programs do.
After creating the database, some environment variables is needed to run the program:
```bash
export DBHOST=your.database.hostname
export DBUSER=your.database.username
export DBPASSWORD=your.database.password
export DBNAME=your.database.name
export DBPORT=3306
```

Then, you can run the code:
```bash
go run .
# or
go build -o api && ./api
```
Or you can just use docker compose and run:
```bash
docker compose up -d
```
Then, you can access it through localhost:8080.

## Pull Requests
1. Fork the project.
2. Make your changes locally.
3. Commit and create a Pull Request. In the text, for now, just explain what you did and that's it.
4. After you've created it, just wait a little, I will review it and merge it if it's everything is alright.

## TODO
1. Store images (with limit size)
2. Change some options to get from environment variables
3. Ability to create a paste with password, and be able to delete it if you have the password
4. Improve return code and messages
5. Is the sha256 a good way of verify the code?
6. Probably in the future a retention period will be necessary
7. Create users
