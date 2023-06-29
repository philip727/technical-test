# Prerequisites
 - **[Golang](https://golang.org/dl/)**
 - **[SSMS/SQL Server Management Studio](https://learn.microsoft.com/en-us/sql/ssms/download-sql-server-management-studio-ssms?view=sql-server-ver16)**
 - **[Postman](https://www.postman.com/downloads/)**
 - **[Docker](https://www.docker.com/products/cli/)**

# Before Hand
Couldn't get the **Bonus** working. Tried many different ways and couldn't figure out how to get a specific status code from "graphql-go" handlers. Couldn't check for error type and couldn't add to "extensions" for specific error codes.

# Setting up the project locally
###  Cloning the repository
```terminal
git clone https://github.com/philip727/technical-test
```
###  Starting docker
Go into the file where you cloned the repository.
```bash
cd technical-test
```
Start the SMTP Server and Azure MSSQL Server from Docker. We use the -d flag(detached) to run the containers in the background.
```bash
docker-compose up -d
```

### Populating the database
Time to populate the MSSQL Database. Connect to the Database using your preferred tool (server: 127.0.0.1, port: 1433, username: sa, password: YourPassword123)

Two files have been provided:
 1. **create.sql**
 2. **users.sql**
 
Execute them in the same order as above and your MSSQL database will now have a table for employees and will now be populated with 10 new employees. Make sure to run each query in **create.sql** one at a time, sequentially. The passwords are provided in the **users.sql** file.

# Viewing Mailhog
Go to [127.0.0.1:8025](http://127.0.0.1:8025) and this will bring you to the mailhog UI, here you will be able to see all emails that 
have been sent when creating an employee

### Running the application
**MAKE SURE YOU HAVE GO INSTALLED. IF NOT GO HERE: [Golang Download](https://golang.org/dl/)**.
First we need to install the go dependencies, run these commands in order
```bash
go mod init
go mod download
```
Now we start the go app. This will start the go app where we will make all our requests to.
```bash
go run main.go
```

### Building the application
**MAKE SURE YOU HAVE GO INSTALLED. IF NOT, GO HERE: [Golang Download](https://golang.org/dl/)**.
First we need to install the go dependencies, run these commands in order
```bash
go mod init
go mod download
```
Now we build the go app.  This will create an executable for you to run.
```bash
go build
```
Then start the .exe that will be made in the root directory, might be called: **tech-test.exe**.

#  Testing with Postman
**MAKE SURE YOU HAVE POSTMAN INSTALLED. IF NOT, GO HERE:  [Postman Download](https://www.postman.com/downloads/)**.
Import the **tests.postman_collection.json** from the root directory into Postman, click on a request and press the send button and it will send a request.

# Production Readiness
Setting up for production is as simple as changing the **PRODUCTION** key in the **.env** to 1 and it will use all the production keys rather than the test keys.
```dotenv
PRODUCTION=1 # OR 0
```


