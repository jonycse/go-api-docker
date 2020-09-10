### Go API with docker example

## Features
+ Run go api inside a docker container
+ watch a directory, check if any file(s) modified


#### Setup
+ docker-compose
    + docker-compose up
+ docker
    + ``docker build -t go-api-docker .``
    + ``docker run -p 9000:8080 -v /home/jony/Developer/jony/go-api-docker/data:/app/data -it go-api-docker``
    
#### API URL
+ GET http://localhost:9000/



#### Some docker commands
+ ``docker build -t go-api-docker .`` build docker container
+ ``docker images`` list of all docker containers
+ ``docker run -p 9000:8080 -it go-api-docker`` run a docker container
+ ``docker stop b28a83cda47e`` stop a docker container
+ ``docker rmi -f d584516265a1`` remove a docker container
+ ``docker run -p 9000:8080 -v /home/jony/Developer/jony/go-api-docker/data:/app/data -it go-api-docker`` mount __data__ folder inside project folder (optional)
+ ``docker network prune``  remove all networks not used by at least one container
