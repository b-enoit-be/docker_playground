### Goal

Demonstration of:
* the possibility to define environment variables, accessible inside the container and, so, the application
* the possibility to give those variables some default values (from shell substitution)
* to adapt them via the command line
* or via environment files

### Running this example

```shell
docker compose up
USR='John Doe' docker compose up
cp .env.sample .env && docker compose up && rm .env
```

Now, look at [_app/hello-world.go_][1], [_docker-compose.yml_][2], [_.env.sample_][3], [_.acme.env_][4]
and see how the variables are used and defined.

[1]: app/hello-world.go
[2]: docker-compose.yml
[3]: .env.sample
[4]: .acme.env
