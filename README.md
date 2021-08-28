# Modanisa Assignment
## Design choices;
- Controller-Handler architecture is used. 
  In this case it would be Handler-Manager.
  Handler is the communication channel with the view layer.
  Manager is the one that gathers data from database.
  The reason was that the project is too small for any cqrs or related architeture. 
- Also, redux is not used due to size of the project.
- Some refactors and design improvements would be needed in the future.

## To download and start the backend application;
``` 
docker pull berkay94/todo-list_app:0.0.1
```
``` 
docker-compose up -d mysql
```
``` 
docker-compose up -d app
```
``` 
go run cmd/migrations/main.go 
```

### You can run the tests after you're done with the app.
```
go test ./ ./internal/handlers
```

## To download and start the frontend application

```
docker pull berkay94/todo-list_frontend:0.0.1
```
```
docker-compose up -d react-ui
```

#### Now you can test it from localhost:3000
