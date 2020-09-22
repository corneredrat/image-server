FROM        golang:1.15.1
COPY        . /code
WORKDIR     /code
RUN         go get -u github.com/swaggo/swag/cmd/swag
RUN         swag init
RUN         go mod vendor
RUN         go build -mod vendor -o server
CMD         ./server


