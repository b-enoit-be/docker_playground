### Goal

Demonstration of:
* the possibility to address a service available on the host machine supporting the docker container

### Running this example

On your host
```shell
nc -kl 12345
```

In another terminal of your host
```shell
docker compose up -d
docker compose exec test nc host.docker.internal 12345
```

Now, type any message and see it appear in the first terminal
— so, effectively seeing the container message appear on your host.

Browse in [_docker-compose.yml_][1] and see how the [`extra_hosts`][2] parameter
brings the possibility to address the host via `host.docker.internal`.

[1]: docker-compose.yml
[2]: https://docs.docker.com/compose/compose-file/compose-file-v3/#extra_hosts
