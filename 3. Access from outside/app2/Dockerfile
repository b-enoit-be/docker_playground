FROM alpine

COPY index.http index.http

CMD ash -c "while true; do read; cat index.http | nc -lp 8080; done"