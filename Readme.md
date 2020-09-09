### Go API with docker example


#### Some docker commands
+ ``docker build -t go-api-docker .`` build docker container
+ ``docker images`` list of all docker containers
+ ``docker run -p 9000:8080 go-api-docker`` run a docker container
+ ``docker stop b28a83cda47e`` stop a docker container
+ ``docker rmi -f d584516265a1`` remove a docker container

#### API URL
+ GET http://localhost:9000/
