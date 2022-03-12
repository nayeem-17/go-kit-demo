## Go-kit example

## Account micro-service 
### API list
- GET  /account/{id}    
    for fetching an account by id. 
- POST /account/    
    for creating an account.

Project structure:
![tree](tree.png)

- **controller**: This folder contains the controller (struct and interface) for the account micro-service. 
- **database**: This folder contains the database config and models in [`model`](./account/database/model/) folder, for the account micro-service.
- **convertor**: This folder contains the  request ,response decoder encoder & dto for the account micro-service.
- **transport**: This folder contains the endpoints file and server file for the account micro-service.

## How to run this microservice:
- Clone this repository
- Install [Go](https://go.dev/doc/install)
- Install all dependencies by running this command
```bash
    go mod vendor
```
- Create a `.env` file in the root directory of this project. & add the following content:
    ```
    MONGO_URI=mongodb://root:example@localhost:27017/ # your mongodb url
    ```
- Run this command: `go run main.go`