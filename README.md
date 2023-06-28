# Setting up the project locally
####  Cloning the repository
```bash
git clone https://github.com/philip727/technical-test
```
####  Starting docker
Go into the file where you cloned the repository.
```bash
cd technical-test
```
Start the SMTP Server and Azure MSSQL Server from Docker. We use the -d flag(detached) to run the containers in the background.
```bash
docker-compose up -d
```

#### Populating the database
Time to populate the MSSQL Database. Connect to the Database using your preferred tool (server: 127.0.0.01, port: 1433, username: sa, password: YourPassword123)

Two files have been provided:
 1. **create.sql**
 2. **users.sql**
 
Run them in the same order as above and your MSSQL database will now be populated with 10 users, the passwords are provided in the **users.sql** file.

#### Running the application
**MAKE SURE YOU HAVE GO INSTALLED. IF NOT GO HERE: [Golang Download](https://golang.org/dl/)**
First we need to install the go dependencies, run these commands in order
```bash
go mod init
go mod download
```
Now we start the go app. This will start the go app where we will make all our requests to.
```bash
go run main.go
```
