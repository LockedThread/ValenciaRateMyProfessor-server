FROM golang

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main .
RUN chmod +x /app/main
RUN echo "you're mom"

CMD ["/app/main"]