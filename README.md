# Search-Movie-Service

## Rest API using port 9090
## gRPC API using port 8080

#### Run the Testing
```bash
$ make test
```
#### Run the Applications
Here is the steps to run it with `docker-compose`

```bash
#move to directory
$ cd workspace

# Clone into project directory
$ git clone https://github.com/bxcodec/go-clean-arch.git

# move to project
$ cd go-clean-arch

# Build the docker image first
$ make docker

# Run the application
$ make run

# check if the containers are running
$ docker ps

# Execute the call
$ curl localhost:9090/api/v1/movies/tt3606756

# Stop
$ make stop
```
