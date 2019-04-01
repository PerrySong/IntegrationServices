# Integration service

## How to build dockerfile for golang project

1. Dependencies management tool:
https://golang.github.io/dep/

2. Write DOCKERFILE
https://stackoverflow.com/questions/47837149/build-docker-with-go-app-cannot-find-package/47837312 (Second answer is helpful)  
https://blog.golang.org/docker 

## Add dependencies

Example
```
$ dep ensure -add github.com/foo/bar github.com/baz/quux
```

## Run
```
$ docker run -p 8080:8080 dockertest
```
Server will be exposed at localhost:8080

## Cheat sheet
Login to local postgres 
```
$ psql -d postgres -U postgres
```
Show table
```
$ \dt
```

List listening port
```
$ netstat -ap tcp | grep -i "listen"
```

start local postgres with brew

```
$ brew services start postgresql
$ brew services stop postgresql
```

Find running process in 5432
```
sudo lsof -i :5432
```


## TODO
1. Does gorm prevent sql injection? https://github.com/jinzhu/gorm/issues/152
2. Remove hard code username/password
3. * Associate id with github token