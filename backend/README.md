## Frameworks / Packages

* https://github.com/gin-gonic/gin
* https://github.com/go-gorm/gorm
* https://github.com/go-gormigrate/gormigrate 
* https://github.com/swaggo/swag

```
  swag init --parseDependency --parseInternal
  ```

* https://github.com/swaggo/gin-swagger

## GO Resources

* https://go.dev/doc/tutorial/web-service-gin
* https://quii.gitbook.io/learn-go-with-tests/
* https://gobyexample.com/

## Issues

* Saggo: Generate using Gorm types: https://github.com/swaggo/swag/issues/810

## Docker Commands

* Build the docker container (run from this directory)
  

```
  docker build -t walk-the-docs/todolist:0.0 .
  ```

* Run the docker container and forward 8080 
  

```
  docker run -p 8080:8080 walk-the-docs/todolist:0.0
  ```

## Deployment

* https://learn.hashicorp.com/tutorials/terraform/digitalocean-provider?in=terraform/applications

## Other Random Stuff I've Learned

* Docker refresher from FireShip: https://www.youtube.com/watch?v=gAkwW2tuIqE&t=53s
* Another library for go migrations: https://github.com/golang-migrate/migrate
* https://en.wikipedia.org/wiki/Marshalling_(computer_science)
* https://www.howtographql.com/graphql-go/0-introduction/
