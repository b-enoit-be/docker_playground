```shell
docker compose up --build
```

```shell
docker compose run app2 "foo" 1>/dev/null
docker compose run app2 "bar" 1>/dev/null
docker compose run app2 "baz" 1>/dev/null
```