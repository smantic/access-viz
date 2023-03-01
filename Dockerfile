FROM golang:1:20 AS builder 
COPY . .
RUN env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -v -o access-viz -a -installsuffix cgo 

FROM alpine
COPY --from=builder /access-viz ./access-viz
EXPOSE 8080
ENTRYPOINT ["./access-viz"]
