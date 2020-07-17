
FROM golang:latest AS builder


LABEL maintainer="Sergio Chairez <schairezv@gmail.com>"


WORKDIR /app 

COPY go.mod .


RUN go mod download
COPY . . 
ENV PORT 8000
RUN go build 

CMD [ "./neo" ]


# Set necessary environmet variables needed for our image
# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64









