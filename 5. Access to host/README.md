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

Browse in `docker-compose.yml` and see how the `extra_hosts` parameter
brings the possibility to address the host via `host.docker.internal`.