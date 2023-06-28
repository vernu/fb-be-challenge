# Backend Developer Assignment


## instructions for setting up server

Before proceeding with the setup, ensure that the following prerequisites are met:

- You have a MongoDB database running locally or on atlas
- You have Golang installed on your machine

## Adding the environment variables
The following environment variables are required to run the server:
- PORT
- MONGO_URI - the database connection string
- DEFAULT_DB - the database name
- INFURA_KEY - the infura key for interacting with the ethereum blockchain, you can get one from [here](https://infura.io/)

## Running the server

once all the environment variables are added to .env file, you can run the server by running the following commands:

```bash
$ go mod download
$ go run main.go
```

This will start the server on the port specified in the `.env` file. You can now make requests to the API using tools like cURL or Postman.
