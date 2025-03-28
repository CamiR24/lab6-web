FROM golang:1.20

# Directorio de trabajo
WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main

EXPOSE 8080

CMD ["go", "run", "main"]