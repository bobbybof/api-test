## Production
### Prerequisites
- [Docker](https://docs.docker.com/get-docker/)
### How to run
1. create `.env` by copying from `.env.example`, adjust the value based on your machine
2. run command `docker compose up`
3. app running on `http://localhost:8888`
    
### Development
#### Prerequisite
1. Make sure you already installed PostgreSQL in your system, the [Link](https://www.postgresql.org/download/) to Download and the instruction to installing PostgreSQL
2. Setup `.env` file, you can copy the value from `.env.example` and adjust the value based on your machine
3. for migration, this project use [golang-migrate](https://github.com/golang-migrate/migrate)
4. I used [Air](https://github.com/cosmtrek/air) for hot reloading during development. If it is not yet installed in your machine, you may do the following command to install it
    ```
    go install github.com/cosmtrek/air@latest
    ```
5. make sure your go version is `1.22.3`
#### API Collections
- Bruno
 if you have installed [Bruno](https://github.com/usebruno/bruno) in your machine, you can import the `api-collections/` collection
- Postman
if you have installed [Postman](https://www.postman.com/downloads/) in your machine, you can import the `api-collections/postman.json` collection

#### Runing
0. Install `go.mod` modules
    ```
    go mod download
    ```
1. if you already adjust `.env` you can run the migration
    ```
    make migrateup
    ```
2. Run the project
    ```
    air
    ```
### Testing
- For testing you can run
    ```
    make test
    ```