## Frameworks / Packages

* https://github.com/gin-gonic/gin
* https://github.com/go-gorm/gorm
* https://github.com/swaggo/swag

```
  swag init --parseDependency --parseInternal
  ```

* https://github.com/swaggo/gin-swagger
* https://github.com/gin-contrib/cors

### Testing

I really do view testing as crucial and it's been my experience at companies in the past that if writing proper
unit-tests get put off in the early stages of a project, then it just never gets done. However, this being a toy-project
which will never have to scale, I'm opting to skip them. What *might* scale is some of what I've gleaned after a few hours
of scouring the internet about testing in Golang. Maybe I'll find myself coming back to some of the links below.

* https://romanyx90.medium.com/testing-database-interactions-using-go-d9512b6bb449

  The article above lists a variety of approaches to testing a database in GO. 
  + https://github.com/go-testfixtures/testfixtures 
  
  This one in particular looks interesting. We could spin up a test database in docker-compose, seed it
  using the test fixtures, and then tap the api and check the responses. For swapping out the databases 
  the following stack overflow link may be of use.
    - https://stackoverflow.com/questions/14249217/how-do-i-know-im-running-within-go-test

## Other Interesting Go Libraries

* Inspired by: https://stackoverflow.blog/2022/03/09/rewriting-bash-scripts-in-go-using-black-box-testing/
  + https://github.com/spf13/cobra/
* Out-of-the-box swagger docs (hey that rhymes!): https://github.com/emicklei/go-restful
* A library with better migrations: https://github.com/go-gormigrate/gormigrate 
* Another library for go migrations: https://github.com/golang-migrate/migrate
* Graph QL in Go: https://www.howtographql.com/graphql-go/0-introduction/

## GO Resources

* https://go.dev/doc/tutorial/web-service-gin
* https://quii.gitbook.io/learn-go-with-tests/
* https://gobyexample.com/

## Issues

* Swaggo: Generate using Gorm types: https://github.com/swaggo/swag/issues/810

## Docker Commands

* Build the docker container (run from this directory)
  

```
  docker build -t walk-the-docs/todolist:0.0 .
  ```

* Run the docker container and forward 8080 
  

```
  docker run -p 8080:8080 walk-the-docs/todolist:0.0
  ```

## Other Random Stuff I've Learned

* Docker refresher from FireShip: https://www.youtube.com/watch?v=gAkwW2tuIqE&t=53s
* https://en.wikipedia.org/wiki/Marshalling_(computer_science)
