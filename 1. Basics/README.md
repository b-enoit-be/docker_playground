### Goal

Demonstration of:
* the possibility to run any application without the need to a install complex set of tooling on the host
* the fact that a _Dockerfile_ is a cookbook to have a running environment
* the possibility to [_bind mount_][1] files from the host into a container

### Running this example

```shell
docker compose up
```

Now, adapt the file [_app/hello-world.go_][2] and see the message changing when you re-run the container with the command above.

[1]: https://docs.docker.com/storage/bind-mounts/
[2]: app/hello-world.go