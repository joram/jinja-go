FROM golang:1.7-alpine

# ADD bin/jinja-go /app/jinja-go
WORKDIR /app

CMD ["./jinja-go"]