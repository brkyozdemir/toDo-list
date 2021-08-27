#Modanisa Assignment
##Design choices;
- Controller-Handler architecture is used. 
  In this case it would be Handler-Manager. 
  The reason was that the project is too small for any cqrs or related architeture. 
- 
##To download and start the application;
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

##You can run the tests after you're done with the app.
