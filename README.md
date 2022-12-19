# Go Backend API Scaffold

This repository contains the scaffold code to bootstrap the development of you application APIs in go and get your up and running.

## How to get started

- Create go directory locally.

- cd into the directory like so; `cd go`.

- Run `export PATH=$GOPATH/bin:$PATH` to export the go path

- Clone the repository by running this command:

```script
    git clone git@github.com:MaestroJolly/go-be-api-scaffold.git
```
- Next, change directory into the go-be-api-scaffold folder like so; `cd go-be-api-scaffold`.

- Run `go run .` to install the dependencies.

## How to start the app

- Rename the `.env.local` file to `.env` file.
- Set your preferred port number.
- Create a database configuration (Our preferred database in postgres, You can use your preferred database).
- Add the username, password, database name, port number to the env file.
- Run `go run main.go`.
- App should start running on `http://localhost:${YOUR_PREFERRED_PORT}`.



