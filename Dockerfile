FROM golang 

ENV GO111MODULE on

RUN go version

COPY . /src

WORKDIR /src

RUN go mod download

RUN go build -o app

EXPOSE 8080

ENTRYPOINT [ "./app" ]
