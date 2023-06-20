### Goal

Demonstration of:
* the possibility to run any application without the need to a install complex set of tooling on the host
* the fact that a [_Dockerfile_][1] is a cookbook to have a running environment
* the possibility to [_bind mount_][2] files from the host into a container

### Running this example

```shell
docker compose up
```

Now, adapt the file [_app/hello-world.go_][3] and see the message changing when you re-run the container with the command above.

[1]: https://docs.docker.com/engine/reference/builder/
[2]: https://docs.docker.com/storage/bind-mounts/
[3]: app/hello-world.go