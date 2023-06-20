```shell
docker compose up --build
docker exec 1basics-app1-1 ash -c "echo 'foo' > /var/log/app.log"
docker exec 1basics-app1-1 ash -c "echo 'foo' > /var/log/app1/app.log"
```