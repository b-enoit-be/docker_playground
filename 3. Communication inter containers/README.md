### Goal

Demonstration of:
* the communication between containers
* the isolation of the network of containers towards the outside world
* the host names inside in a container network
* the command to list those networks of containers
* the possibility to expose a subset of interface to the outside world

### Running this example

```shell
docker compose up --build
```
You will get the docker logs.
Now, in another terminal run:
```shell
docker compose run app2 "foo" 1>/dev/null
docker compose run app2 "bar" 1>/dev/null
docker compose run app2 "baz" 1>/dev/null
```
And see how the message are received from `app2` to `app1`.

Now, look inside [_app2/Dockerfile_][1] to see how we use an [_entrypoint_][2], then see the corresponding file, [_app2/bin/entrypoint_][3] file, compare that with the services are defined in the file [_docker-compose.yml_][4] to understand the aliasing made in a Docker Network.

Run:
```shell
docker network ls
```
to discover more about those networks.

Also see how the file is [_app1/Dockerfile_][5] do not restrict the connection in any means possible but is still not reachable from the host:
```shell
nc -zv localhost 12345
```
is going to yield
```none
nc: connectx to localhost port 12345 (tcp) failed: Connection refused
```

Next, get your container down, and run an alternative version of the compose file: [_docker-compose-accessible-outside.yml_][6]

```shell
docker compose down
docker compose \
  --file docker-compose-accessible-outside.yml \
  up
```

Now, the `app1` is also accessible from outside!
```shell
nc -zv localhost 12345
```
is going to yield
```none
Connection to localhost port 12345 [tcp/italk] succeeded!
```

[1]: app2/Dockerfile
[2]: https://docs.docker.com/engine/reference/builder/#entrypoint
[3]: app2/bin/entrypoint
[4]: docker-compose.yml
[5]: app1/Dockerfile
[6]: docker-compose-accessible-outside.yml