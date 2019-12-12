FROM golang:1.13.4-buster AS build
WORKDIR /opt/build
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o dummy-pdf-or-png



FROM debian:stretch
WORKDIR /opt
COPY --from=build /opt/build/dummy-pdf-or-png /opt/build/dummy.pdf /opt/build/dummy.png /opt/build/corrupt-dummy.pdf /opt/
CMD ["./dummy-pdf-or-png"]
