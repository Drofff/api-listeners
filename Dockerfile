FROM golang
COPY . /go/src/api-listeners
WORKDIR /go/src/api-listeners
ENV ozzy_rate=""
CMD ["go", "run", "./app"]
