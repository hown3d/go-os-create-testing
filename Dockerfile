FROM golang

WORKDIR /app
COPY go.mod main.go run-twice.sh /app/

ENTRYPOINT ["sh", "/app/run-twice.sh"]
