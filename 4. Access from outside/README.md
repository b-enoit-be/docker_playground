### Goal

Demonstration of:
* accessing a container from the host
* jumping in an interactive terminal with a container
* how a container can answer an external call (dummy web server with netcat)

TODO: finish this

### Running this example

```shell
docker compose up --build
docker exec 1basics-app1-1 ash -c "echo 'foo' > /var/log/app.log"
docker exec 1basics-app1-1 ash -c "echo 'foo' > /var/log/app1/app.log"
```