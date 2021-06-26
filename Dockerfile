FROM golang

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main .
RUN chmod +x /app/main

CMD ["/app/main"]